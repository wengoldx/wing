// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/ifscg/ifscg.proto

package ifscg

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Ifscgsrv service

type IfscgsrvService interface {
	/// SendEvent use to send message to blank
	SendEvent(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*EventResponse, error)
	/// RecRespFit use to save data to redis
	RecRespFit(ctx context.Context, in *EventResponse, opts ...client.CallOption) (*Response, error)
}

type ifscgsrvService struct {
	c    client.Client
	name string
}

func NewIfscgsrvService(name string, c client.Client) IfscgsrvService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "ifscg"
	}
	return &ifscgsrvService{
		c:    c,
		name: name,
	}
}

func (c *ifscgsrvService) SendEvent(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*EventResponse, error) {
	req := c.c.NewRequest(c.name, "Ifscgsrv.SendEvent", in)
	out := new(EventResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ifscgsrvService) RecRespFit(ctx context.Context, in *EventResponse, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Ifscgsrv.RecRespFit", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Ifscgsrv service

type IfscgsrvHandler interface {
	/// SendEvent use to send message to blank
	SendEvent(context.Context, *EmptyRequest, *EventResponse) error
	/// RecRespFit use to save data to redis
	RecRespFit(context.Context, *EventResponse, *Response) error
}

func RegisterIfscgsrvHandler(s server.Server, hdlr IfscgsrvHandler, opts ...server.HandlerOption) error {
	type ifscgsrv interface {
		SendEvent(ctx context.Context, in *EmptyRequest, out *EventResponse) error
		RecRespFit(ctx context.Context, in *EventResponse, out *Response) error
	}
	type Ifscgsrv struct {
		ifscgsrv
	}
	h := &ifscgsrvHandler{hdlr}
	return s.Handle(s.NewHandler(&Ifscgsrv{h}, opts...))
}

type ifscgsrvHandler struct {
	IfscgsrvHandler
}

func (h *ifscgsrvHandler) SendEvent(ctx context.Context, in *EmptyRequest, out *EventResponse) error {
	return h.IfscgsrvHandler.SendEvent(ctx, in, out)
}

func (h *ifscgsrvHandler) RecRespFit(ctx context.Context, in *EventResponse, out *Response) error {
	return h.IfscgsrvHandler.RecRespFit(ctx, in, out)
}