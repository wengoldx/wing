// Copyright (c) 2018-2028 Dunyu All Rights Reserved.
//
// Author      : https://www.wengold.net
// Email       : support@wengold.net
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2019/05/22   yangping       New version
// 00002       2019/06/30   zhaixing       Add function from godfs
// -------------------------------------------------------------------

package comm

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/wengoldx/wing/logger"
	"os"
	"os/signal"
	"syscall"
)

// ===========================
// For setup HTTP Server
// ===========================

// Start and excute http server base on beego, by default, it just
// support restful interface not socket.io connection, but you can
// set allowCredentials as true to on socket.io conect function.
//
// `USAGE` :
//
//	// use for restful interface server
//	func main() {}
//		// comm.HttpServer(false) or
//		comm.HttpServer()
//
//	// use for both restful and socket.io server
//	func main() {
//		comm.HttpServer(true)
//	}
func HttpServer(allowCredentials ...bool) {
	ignoreSysSignalPIPE()
	if len(allowCredentials) > 0 {
		accessAllowOriginBy(beego.BeforeRouter, "*", allowCredentials[0])
		accessAllowOriginBy(beego.BeforeStatic, "*", allowCredentials[0])
	} else {
		accessAllowOriginBy(beego.BeforeRouter, "*", false)
		accessAllowOriginBy(beego.BeforeStatic, "*", false)
	}

	// just output log to file on prod mode
	if beego.BConfig.RunMode != "dev" &&
		logger.GetLevel() != logger.LevelDebug {
		beego.BeeLogger.DelLogger(logs.AdapterConsole)
	}
	beego.Run()
}

// Start and excute both restful and socket.io server
func Rest4SioServer() {
	HttpServer(true)
}

// Allow cross domain access for localhost,
// the port number must config in /conf/app.conf file like :
//
// ---
//
//	; Server port of HTTP
//	httpport=3200
func AccessAllowOriginByLocal(category int, allowCredentials bool) {
	if beego.BConfig.Listen.HTTPPort > 0 {
		localhosturl := fmt.Sprintf("http://127.0.0.1:%v/", beego.BConfig.Listen.HTTPPort)
		accessAllowOriginBy(category, localhosturl, allowCredentials)
	}
}

// Ignore system PIPE signal
func ignoreSysSignalPIPE() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGPIPE)
	go func() {
		for {
			select {
			case sig := <-sc:
				if sig == syscall.SIGPIPE {
					logger.E("!! IGNORE BROKEN PIPE SIGNAL !!")
				}
			}
		}
	}()
}

// Allow cross domain access for the given origins
func accessAllowOriginBy(category int, origins string, allowCredentials bool) {
	beego.InsertFilter("*", category, cors.Allow(&cors.Options{
		AllowAllOrigins:  !allowCredentials,
		AllowCredentials: allowCredentials,
		AllowOrigins:     []string{origins}, // use to set allow Origins
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "Authoration", "Token"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
	}))
}
