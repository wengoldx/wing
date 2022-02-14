// Copyright (c) 2019-2029 DY All Rights Reserved.
//
// Author : yangping
// Email  : youhei_yp@163.com
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2022/02/09   yangping       New version
// -------------------------------------------------------------------

package wsio

import (
	"github.com/astaxie/beego"
	sio "github.com/googollee/go-socket.io"
	"github.com/wengoldx/wing/invar"
	"github.com/wengoldx/wing/logger"
	"github.com/wengoldx/wing/wsio/core"
	"net/http"
	"sync"
	"time"
	"unsafe"
)

const (
	serverPingInterval = 30 * time.Second
	serverPingTimeout  = 60 * time.Second
	maxConnectCount    = 200000
)

// Socket.io connecte information, it will generate a socket server
// and register as http Handler to listen socket signalings.
//
// ----
//
// `USAGE` :
//
//	// types.go : define socket event label
//	const (
//		EvtLabel01 = "evt-01"
//		EvtLabel02 = "evt-02"
//	)
//
//	// controllers.go : define socket event func
//	func SIOEventFunc01(c *wsio.SocketController, uuid, params string) string { return "" }
//	func SIOEventFunc02(c *wsio.SocketController, uuid, params string) string { return "" }
//
//	// routers.go : register socket events
//	import "github.com/wengoldx/wing/wsio"
//
//	// register single socket event
//	wsio.On(types.EvtLabel01, socketEvents)
//
//	// register multiple socket events
//	var socketEvents = map[string]wsio.SocketEvent{
//		types.EvtLabel02: ctrl.SIOEventFunc02,
//	}
//	wsio.Ons(socketEvents)
type WingSIO struct {
	lock sync.Mutex             // Mutex sync lock
	h2u  map[uintptr]string     // http request to uuid
	opts map[string]interface{} // client option data caches

	// socket server
	server *sio.Server

	// socket events controllers
	controllers map[string]*SocketController

	// socket golbal callback handler to execute clients
	// authenticate, connect, disconnect actions
	handler *SocketHandler
}

// Socket connection server
var wsc *WingSIO

func init() {
	wsc = &WingSIO{
		h2u:         make(map[uintptr]string),
		opts:        make(map[string]interface{}),
		controllers: make(map[string]*SocketController),
	}

	// set http handler for socke.io
	handler, err := wsc.createHandler()
	if err != nil {
		panic(err)
	}

	// set socket.io routers
	beego.Handler("/"+beego.BConfig.AppName+"/socket.io", handler)
	logger.I("Initialized socket.io routers...")
}

// Set handler to execute clients authenticate, connect and disconnect.
func Handler(handler *SocketHandler) {
	if handler != nil {
		logger.E("Invalid socket event or callback!")
		wsc.handler = handler
	}
}

// Register single socket event, the input params numst not empty and nil.
func On(evt string, callback SocketEvent) error {
	if evt == "" || callback == nil {
		logger.E("Invalid socket event or callback!")
		return invar.ErrInvalidState
	}

	logger.D("Register socket event:", evt)
	evtCtler := &SocketController{evt: evt, handler: callback}

	wsc.controllers[evt] = evtCtler
	wsc.server.On(evt, func(sc sio.Socket, uuid, params string) string {
		return callback(evtCtler, uuid, params)
	})
	return nil
}

// Register multiple socket events from given mapping, in shoule interrupt
// when one event register failed.
func Ons(events map[string]SocketEvent) error {
	for evt, callback := range events {
		if err := On(evt, callback); err != nil {
			return err
		}
	}
	return nil
}

// createHandler create http handler for socket.io
func (cc *WingSIO) createHandler() (http.Handler, error) {
	server, err := sio.NewServer(nil)
	if err != nil {
		return nil, err
	}
	cc.server = server

	// set socket.io ping interval and timeout
	logger.I("Set socket ping-pong and timeout")
	server.SetPingInterval(serverPingInterval)
	server.SetPingTimeout(serverPingTimeout)

	// set max connection count
	server.SetMaxConnection(maxConnectCount)

	// set auth middleware for socket.io connection
	server.SetAllowRequest(func(req *http.Request) error {
		if err = cc.onAuthentication(req); err != nil {
			logger.E("Authenticate err:", err)
			return err
		}
		return nil
	})

	// set connection event
	server.On("connection", func(sc sio.Socket) {
		cc.onConnect(sc)
	})

	// set disconnection event
	server.On("disconnection", func(sc sio.Socket) {
		cc.onDisconnected(sc)
	})

	logger.I("Created socket.io handler")
	return server, nil
}

// onAuthentication event of authentication
func (cc *WingSIO) onAuthentication(req *http.Request) error {
	if err := req.ParseForm(); err != nil {
		logger.E("Prase url err:", err)
		return err
	}

	// get client auth token
	token := req.Form.Get("token")
	if token == "" {
		logger.E("Not found client auth token!")
		return invar.ErrInvalidClient
	}

	// auth client token by handler if set Authenticate function
	// handler, or just use token as uuid when not set.
	uuid := token
	var option interface{} = nil
	if cc.handler != nil && cc.handler.Authenticate != nil {
		uid, opt, err := cc.handler.Authenticate(token)
		if err != nil || uid == "" {
			return invar.ErrInvalidClient
		}
		uuid, option = uid, opt
	}

	// bind http.Request -> uuid
	h := uintptr(unsafe.Pointer(req))
	logger.I("Client:", uuid, "http ptr:", h)
	cc.bindHTTP2UUIDLocked(h, uuid, option)
	return nil
}

// onConnect event of connect
func (cc *WingSIO) onConnect(sc sio.Socket) {
	// found client uuid and unbind -> http.Request
	h := uintptr(unsafe.Pointer(sc.Request()))
	logger.I("Socket http request pointer:", h)

	uuid, option := cc.unbindUUIDFromHTTPLocked(h)
	if uuid == "" {
		logger.E("Invalid socket request: not found uuid")
		sc.Disconnect()
		return
	}

	clientPool := core.Clients()
	if err := clientPool.Register(uuid, sc, option); err != nil {
		logger.E("Faild register socket client:", uuid)
		sc.Disconnect()
		return
	}

	// handle connect callback for socket with uuid
	if cc.handler != nil && cc.handler.Connect != nil {
		if err := cc.handler.Connect(uuid, option); err != nil {
			logger.E("Client:", uuid, "connect socket err:", err)
			sc.Disconnect()
		}
	}
	logger.I("Connected socket client:", uuid)
}

// onDisconnected event of disconnect
func (cc *WingSIO) onDisconnected(sc sio.Socket) {
	clientPool := core.Clients()
	uuid, opt := clientPool.Deregister(sc)

	if cc.handler != nil && cc.handler.Disconnect != nil {
		cc.handler.Disconnect(uuid, opt)
	}
}

// bindHTTP2UUIDLocked bind http request pointer -> uuid on locked status
func (cc *WingSIO) bindHTTP2UUIDLocked(h uintptr, uuid string, opt interface{}) {
	cc.lock.Lock()
	defer cc.lock.Unlock()

	cc.h2u[h] = uuid
	cc.opts[uuid] = opt
}

// unbindUUIDFromHTTPLocked unbind uuid -> http request pointer on locked status
func (cc *WingSIO) unbindUUIDFromHTTPLocked(h uintptr) (string, interface{}) {
	cc.lock.Lock()
	defer cc.lock.Unlock()

	uuid := cc.h2u[h]
	delete(cc.h2u, h)

	opt := cc.opts[uuid]
	delete(cc.opts, uuid)
	return uuid, opt
}
