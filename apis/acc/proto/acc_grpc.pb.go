// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.3
// source: proto/acc.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AccClient is the client API for Acc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccClient interface {
	ViaToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*UUID, error)
	GetProfSumms(ctx context.Context, in *UUIDS, opts ...grpc.CallOption) (*ProfSumms, error)
	GetProfile(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*Profile, error)
	UpdateContact(ctx context.Context, in *UpdateContect, opts ...grpc.CallOption) (*EmptyMessage, error)
	GetContact(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*ContactInfo, error)
	BindAccount(ctx context.Context, in *BindReq, opts ...grpc.CallOption) (*Token, error)
}

type accClient struct {
	cc grpc.ClientConnInterface
}

func NewAccClient(cc grpc.ClientConnInterface) AccClient {
	return &accClient{cc}
}

func (c *accClient) ViaToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*UUID, error) {
	out := new(UUID)
	err := c.cc.Invoke(ctx, "/proto.Acc/ViaToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accClient) GetProfSumms(ctx context.Context, in *UUIDS, opts ...grpc.CallOption) (*ProfSumms, error) {
	out := new(ProfSumms)
	err := c.cc.Invoke(ctx, "/proto.Acc/GetProfSumms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accClient) GetProfile(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*Profile, error) {
	out := new(Profile)
	err := c.cc.Invoke(ctx, "/proto.Acc/GetProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accClient) UpdateContact(ctx context.Context, in *UpdateContect, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := c.cc.Invoke(ctx, "/proto.Acc/UpdateContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accClient) GetContact(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*ContactInfo, error) {
	out := new(ContactInfo)
	err := c.cc.Invoke(ctx, "/proto.Acc/GetContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accClient) BindAccount(ctx context.Context, in *BindReq, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/proto.Acc/BindAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccServer is the server API for Acc service.
// All implementations must embed UnimplementedAccServer
// for forward compatibility
type AccServer interface {
	ViaToken(context.Context, *Token) (*UUID, error)
	GetProfSumms(context.Context, *UUIDS) (*ProfSumms, error)
	GetProfile(context.Context, *UUID) (*Profile, error)
	UpdateContact(context.Context, *UpdateContect) (*EmptyMessage, error)
	GetContact(context.Context, *UUID) (*ContactInfo, error)
	BindAccount(context.Context, *BindReq) (*Token, error)
	mustEmbedUnimplementedAccServer()
}

// UnimplementedAccServer must be embedded to have forward compatible implementations.
type UnimplementedAccServer struct {
}

func (UnimplementedAccServer) ViaToken(context.Context, *Token) (*UUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViaToken not implemented")
}
func (UnimplementedAccServer) GetProfSumms(context.Context, *UUIDS) (*ProfSumms, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfSumms not implemented")
}
func (UnimplementedAccServer) GetProfile(context.Context, *UUID) (*Profile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedAccServer) UpdateContact(context.Context, *UpdateContect) (*EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateContact not implemented")
}
func (UnimplementedAccServer) GetContact(context.Context, *UUID) (*ContactInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContact not implemented")
}
func (UnimplementedAccServer) BindAccount(context.Context, *BindReq) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BindAccount not implemented")
}
func (UnimplementedAccServer) mustEmbedUnimplementedAccServer() {}

// UnsafeAccServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccServer will
// result in compilation errors.
type UnsafeAccServer interface {
	mustEmbedUnimplementedAccServer()
}

func RegisterAccServer(s grpc.ServiceRegistrar, srv AccServer) {
	s.RegisterService(&Acc_ServiceDesc, srv)
}

func _Acc_ViaToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccServer).ViaToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Acc/ViaToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccServer).ViaToken(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _Acc_GetProfSumms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UUIDS)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccServer).GetProfSumms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Acc/GetProfSumms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccServer).GetProfSumms(ctx, req.(*UUIDS))
	}
	return interceptor(ctx, in, info, handler)
}

func _Acc_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Acc/GetProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccServer).GetProfile(ctx, req.(*UUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Acc_UpdateContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateContect)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccServer).UpdateContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Acc/UpdateContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccServer).UpdateContact(ctx, req.(*UpdateContect))
	}
	return interceptor(ctx, in, info, handler)
}

func _Acc_GetContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccServer).GetContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Acc/GetContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccServer).GetContact(ctx, req.(*UUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Acc_BindAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccServer).BindAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Acc/BindAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccServer).BindAccount(ctx, req.(*BindReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Acc_ServiceDesc is the grpc.ServiceDesc for Acc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Acc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Acc",
	HandlerType: (*AccServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ViaToken",
			Handler:    _Acc_ViaToken_Handler,
		},
		{
			MethodName: "GetProfSumms",
			Handler:    _Acc_GetProfSumms_Handler,
		},
		{
			MethodName: "GetProfile",
			Handler:    _Acc_GetProfile_Handler,
		},
		{
			MethodName: "UpdateContact",
			Handler:    _Acc_UpdateContact_Handler,
		},
		{
			MethodName: "GetContact",
			Handler:    _Acc_GetContact_Handler,
		},
		{
			MethodName: "BindAccount",
			Handler:    _Acc_BindAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/acc.proto",
}
