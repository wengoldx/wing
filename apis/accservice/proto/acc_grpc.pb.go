// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
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
	AccLogin(ctx context.Context, in *AccPwd, opts ...grpc.CallOption) (*Token, error)
	ViaToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*AccPwd, error)
	ViaAdmin(ctx context.Context, in *Admin, opts ...grpc.CallOption) (*AEmpty, error)
	GetProfile(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*Profile, error)
	GetProfSumms(ctx context.Context, in *UIDS, opts ...grpc.CallOption) (*ProfSumms, error)
	GetContact(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*Contact, error)
	GetCreatetime(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*CreateTime, error)
	// GRPC interface for store service
	StoreRegister(ctx context.Context, in *RegStore, opts ...grpc.CallOption) (*UUID, error)
	StoreProfile(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*ProfStore, error)
	StoreProfiles(ctx context.Context, in *UIDS, opts ...grpc.CallOption) (*ProfStores, error)
	StoreRename(ctx context.Context, in *RnStore, opts ...grpc.CallOption) (*AEmpty, error)
	StoreBindWx(ctx context.Context, in *WxBind, opts ...grpc.CallOption) (*AEmpty, error)
	SetContact(ctx context.Context, in *Contact, opts ...grpc.CallOption) (*AEmpty, error)
	BindAccount(ctx context.Context, in *Secures, opts ...grpc.CallOption) (*Token, error)
	UnbindUnionID(ctx context.Context, in *AccPwd, opts ...grpc.CallOption) (*AEmpty, error)
	UnbindUnionID2(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*AEmpty, error)
	// GRPC interface for role verification
	ViaAllow(ctx context.Context, in *AuthModel, opts ...grpc.CallOption) (*AuthResp, error)
}

type accClient struct {
	cc grpc.ClientConnInterface
}

func NewAccClient(cc grpc.ClientConnInterface) AccClient {
	return &accClient{cc}
}

func (c *accClient) AccLogin(ctx context.Context, in *AccPwd, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/proto.Acc/AccLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accClient) ViaToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*AccPwd, error) {
	out := new(AccPwd)
	err := c.cc.Invoke(ctx, "/proto.Acc/ViaToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accClient) ViaAdmin(ctx context.Context, in *Admin, opts ...grpc.CallOption) (*AEmpty, error) {
	out := new(AEmpty)
	err := c.cc.Invoke(ctx, "/proto.Acc/ViaAdmin", in, out, opts...)
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

func (c *accClient) GetProfSumms(ctx context.Context, in *UIDS, opts ...grpc.CallOption) (*ProfSumms, error) {
	out := new(ProfSumms)
	err := c.cc.Invoke(ctx, "/proto.Acc/GetProfSumms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accClient) GetContact(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*Contact, error) {
	out := new(Contact)
	err := c.cc.Invoke(ctx, "/proto.Acc/GetContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accClient) GetCreatetime(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*CreateTime, error) {
	out := new(CreateTime)
	err := c.cc.Invoke(ctx, "/proto.Acc/GetCreatetime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accClient) StoreRegister(ctx context.Context, in *RegStore, opts ...grpc.CallOption) (*UUID, error) {
	out := new(UUID)
	err := c.cc.Invoke(ctx, "/proto.Acc/StoreRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accClient) StoreProfile(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*ProfStore, error) {
	out := new(ProfStore)
	err := c.cc.Invoke(ctx, "/proto.Acc/StoreProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accClient) StoreProfiles(ctx context.Context, in *UIDS, opts ...grpc.CallOption) (*ProfStores, error) {
	out := new(ProfStores)
	err := c.cc.Invoke(ctx, "/proto.Acc/StoreProfiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accClient) StoreRename(ctx context.Context, in *RnStore, opts ...grpc.CallOption) (*AEmpty, error) {
	out := new(AEmpty)
	err := c.cc.Invoke(ctx, "/proto.Acc/StoreRename", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accClient) StoreBindWx(ctx context.Context, in *WxBind, opts ...grpc.CallOption) (*AEmpty, error) {
	out := new(AEmpty)
	err := c.cc.Invoke(ctx, "/proto.Acc/StoreBindWx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accClient) SetContact(ctx context.Context, in *Contact, opts ...grpc.CallOption) (*AEmpty, error) {
	out := new(AEmpty)
	err := c.cc.Invoke(ctx, "/proto.Acc/SetContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accClient) BindAccount(ctx context.Context, in *Secures, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/proto.Acc/BindAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accClient) UnbindUnionID(ctx context.Context, in *AccPwd, opts ...grpc.CallOption) (*AEmpty, error) {
	out := new(AEmpty)
	err := c.cc.Invoke(ctx, "/proto.Acc/UnbindUnionID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accClient) UnbindUnionID2(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*AEmpty, error) {
	out := new(AEmpty)
	err := c.cc.Invoke(ctx, "/proto.Acc/UnbindUnionID2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accClient) ViaAllow(ctx context.Context, in *AuthModel, opts ...grpc.CallOption) (*AuthResp, error) {
	out := new(AuthResp)
	err := c.cc.Invoke(ctx, "/proto.Acc/ViaAllow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccServer is the server API for Acc service.
// All implementations must embed UnimplementedAccServer
// for forward compatibility
type AccServer interface {
	AccLogin(context.Context, *AccPwd) (*Token, error)
	ViaToken(context.Context, *Token) (*AccPwd, error)
	ViaAdmin(context.Context, *Admin) (*AEmpty, error)
	GetProfile(context.Context, *UUID) (*Profile, error)
	GetProfSumms(context.Context, *UIDS) (*ProfSumms, error)
	GetContact(context.Context, *UUID) (*Contact, error)
	GetCreatetime(context.Context, *UUID) (*CreateTime, error)
	// GRPC interface for store service
	StoreRegister(context.Context, *RegStore) (*UUID, error)
	StoreProfile(context.Context, *UUID) (*ProfStore, error)
	StoreProfiles(context.Context, *UIDS) (*ProfStores, error)
	StoreRename(context.Context, *RnStore) (*AEmpty, error)
	StoreBindWx(context.Context, *WxBind) (*AEmpty, error)
	SetContact(context.Context, *Contact) (*AEmpty, error)
	BindAccount(context.Context, *Secures) (*Token, error)
	UnbindUnionID(context.Context, *AccPwd) (*AEmpty, error)
	UnbindUnionID2(context.Context, *UUID) (*AEmpty, error)
	// GRPC interface for role verification
	ViaAllow(context.Context, *AuthModel) (*AuthResp, error)
	mustEmbedUnimplementedAccServer()
}

// UnimplementedAccServer must be embedded to have forward compatible implementations.
type UnimplementedAccServer struct {
}

func (UnimplementedAccServer) AccLogin(context.Context, *AccPwd) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccLogin not implemented")
}
func (UnimplementedAccServer) ViaToken(context.Context, *Token) (*AccPwd, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViaToken not implemented")
}
func (UnimplementedAccServer) ViaAdmin(context.Context, *Admin) (*AEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViaAdmin not implemented")
}
func (UnimplementedAccServer) GetProfile(context.Context, *UUID) (*Profile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedAccServer) GetProfSumms(context.Context, *UIDS) (*ProfSumms, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfSumms not implemented")
}
func (UnimplementedAccServer) GetContact(context.Context, *UUID) (*Contact, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContact not implemented")
}
func (UnimplementedAccServer) GetCreatetime(context.Context, *UUID) (*CreateTime, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCreatetime not implemented")
}
func (UnimplementedAccServer) StoreRegister(context.Context, *RegStore) (*UUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreRegister not implemented")
}
func (UnimplementedAccServer) StoreProfile(context.Context, *UUID) (*ProfStore, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreProfile not implemented")
}
func (UnimplementedAccServer) StoreProfiles(context.Context, *UIDS) (*ProfStores, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreProfiles not implemented")
}
func (UnimplementedAccServer) StoreRename(context.Context, *RnStore) (*AEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreRename not implemented")
}
func (UnimplementedAccServer) StoreBindWx(context.Context, *WxBind) (*AEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreBindWx not implemented")
}
func (UnimplementedAccServer) SetContact(context.Context, *Contact) (*AEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetContact not implemented")
}
func (UnimplementedAccServer) BindAccount(context.Context, *Secures) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BindAccount not implemented")
}
func (UnimplementedAccServer) UnbindUnionID(context.Context, *AccPwd) (*AEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnbindUnionID not implemented")
}
func (UnimplementedAccServer) UnbindUnionID2(context.Context, *UUID) (*AEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnbindUnionID2 not implemented")
}
func (UnimplementedAccServer) ViaAllow(context.Context, *AuthModel) (*AuthResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViaAllow not implemented")
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

func _Acc_AccLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccPwd)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccServer).AccLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Acc/AccLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccServer).AccLogin(ctx, req.(*AccPwd))
	}
	return interceptor(ctx, in, info, handler)
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

func _Acc_ViaAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Admin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccServer).ViaAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Acc/ViaAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccServer).ViaAdmin(ctx, req.(*Admin))
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

func _Acc_GetProfSumms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UIDS)
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
		return srv.(AccServer).GetProfSumms(ctx, req.(*UIDS))
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

func _Acc_GetCreatetime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccServer).GetCreatetime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Acc/GetCreatetime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccServer).GetCreatetime(ctx, req.(*UUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Acc_StoreRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegStore)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccServer).StoreRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Acc/StoreRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccServer).StoreRegister(ctx, req.(*RegStore))
	}
	return interceptor(ctx, in, info, handler)
}

func _Acc_StoreProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccServer).StoreProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Acc/StoreProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccServer).StoreProfile(ctx, req.(*UUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Acc_StoreProfiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UIDS)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccServer).StoreProfiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Acc/StoreProfiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccServer).StoreProfiles(ctx, req.(*UIDS))
	}
	return interceptor(ctx, in, info, handler)
}

func _Acc_StoreRename_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RnStore)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccServer).StoreRename(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Acc/StoreRename",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccServer).StoreRename(ctx, req.(*RnStore))
	}
	return interceptor(ctx, in, info, handler)
}

func _Acc_StoreBindWx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WxBind)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccServer).StoreBindWx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Acc/StoreBindWx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccServer).StoreBindWx(ctx, req.(*WxBind))
	}
	return interceptor(ctx, in, info, handler)
}

func _Acc_SetContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Contact)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccServer).SetContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Acc/SetContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccServer).SetContact(ctx, req.(*Contact))
	}
	return interceptor(ctx, in, info, handler)
}

func _Acc_BindAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Secures)
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
		return srv.(AccServer).BindAccount(ctx, req.(*Secures))
	}
	return interceptor(ctx, in, info, handler)
}

func _Acc_UnbindUnionID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccPwd)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccServer).UnbindUnionID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Acc/UnbindUnionID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccServer).UnbindUnionID(ctx, req.(*AccPwd))
	}
	return interceptor(ctx, in, info, handler)
}

func _Acc_UnbindUnionID2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccServer).UnbindUnionID2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Acc/UnbindUnionID2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccServer).UnbindUnionID2(ctx, req.(*UUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Acc_ViaAllow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthModel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccServer).ViaAllow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Acc/ViaAllow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccServer).ViaAllow(ctx, req.(*AuthModel))
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
			MethodName: "AccLogin",
			Handler:    _Acc_AccLogin_Handler,
		},
		{
			MethodName: "ViaToken",
			Handler:    _Acc_ViaToken_Handler,
		},
		{
			MethodName: "ViaAdmin",
			Handler:    _Acc_ViaAdmin_Handler,
		},
		{
			MethodName: "GetProfile",
			Handler:    _Acc_GetProfile_Handler,
		},
		{
			MethodName: "GetProfSumms",
			Handler:    _Acc_GetProfSumms_Handler,
		},
		{
			MethodName: "GetContact",
			Handler:    _Acc_GetContact_Handler,
		},
		{
			MethodName: "GetCreatetime",
			Handler:    _Acc_GetCreatetime_Handler,
		},
		{
			MethodName: "StoreRegister",
			Handler:    _Acc_StoreRegister_Handler,
		},
		{
			MethodName: "StoreProfile",
			Handler:    _Acc_StoreProfile_Handler,
		},
		{
			MethodName: "StoreProfiles",
			Handler:    _Acc_StoreProfiles_Handler,
		},
		{
			MethodName: "StoreRename",
			Handler:    _Acc_StoreRename_Handler,
		},
		{
			MethodName: "StoreBindWx",
			Handler:    _Acc_StoreBindWx_Handler,
		},
		{
			MethodName: "SetContact",
			Handler:    _Acc_SetContact_Handler,
		},
		{
			MethodName: "BindAccount",
			Handler:    _Acc_BindAccount_Handler,
		},
		{
			MethodName: "UnbindUnionID",
			Handler:    _Acc_UnbindUnionID_Handler,
		},
		{
			MethodName: "UnbindUnionID2",
			Handler:    _Acc_UnbindUnionID2_Handler,
		},
		{
			MethodName: "ViaAllow",
			Handler:    _Acc_ViaAllow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/acc.proto",
}
