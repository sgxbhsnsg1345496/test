// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: user.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	User_GetUserList_FullMethodName   = "/User/GetUserList"
	User_GetUserMobile_FullMethodName = "/User/GetUserMobile"
	User_GetUserId_FullMethodName     = "/User/GetUserId"
	User_CreatUser_FullMethodName     = "/User/CreatUser"
	User_UpdateUser_FullMethodName    = "/User/UpdateUser"
	User_CheckPassword_FullMethodName = "/User/CheckPassword"
)

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	GetUserList(ctx context.Context, in *PageInfo, opts ...grpc.CallOption) (*UserListResponse, error)
	GetUserMobile(ctx context.Context, in *MobileRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
	GetUserId(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
	CreatUser(ctx context.Context, in *CreatUserInfo, opts ...grpc.CallOption) (*UserInfoResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserInfo, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CheckPassword(ctx context.Context, in *PasswordCheckInfo, opts ...grpc.CallOption) (*CheckResponse, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) GetUserList(ctx context.Context, in *PageInfo, opts ...grpc.CallOption) (*UserListResponse, error) {
	out := new(UserListResponse)
	err := c.cc.Invoke(ctx, User_GetUserList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserMobile(ctx context.Context, in *MobileRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, User_GetUserMobile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserId(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, User_GetUserId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CreatUser(ctx context.Context, in *CreatUserInfo, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, User_CreatUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateUser(ctx context.Context, in *UpdateUserInfo, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, User_UpdateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CheckPassword(ctx context.Context, in *PasswordCheckInfo, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, User_CheckPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	GetUserList(context.Context, *PageInfo) (*UserListResponse, error)
	GetUserMobile(context.Context, *MobileRequest) (*UserInfoResponse, error)
	GetUserId(context.Context, *IdRequest) (*UserInfoResponse, error)
	CreatUser(context.Context, *CreatUserInfo) (*UserInfoResponse, error)
	UpdateUser(context.Context, *UpdateUserInfo) (*emptypb.Empty, error)
	CheckPassword(context.Context, *PasswordCheckInfo) (*CheckResponse, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) GetUserList(context.Context, *PageInfo) (*UserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserList not implemented")
}
func (UnimplementedUserServer) GetUserMobile(context.Context, *MobileRequest) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserMobile not implemented")
}
func (UnimplementedUserServer) GetUserId(context.Context, *IdRequest) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserId not implemented")
}
func (UnimplementedUserServer) CreatUser(context.Context, *CreatUserInfo) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatUser not implemented")
}
func (UnimplementedUserServer) UpdateUser(context.Context, *UpdateUserInfo) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserServer) CheckPassword(context.Context, *PasswordCheckInfo) (*CheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPassword not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_GetUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetUserList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserList(ctx, req.(*PageInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserMobile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MobileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserMobile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetUserMobile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserMobile(ctx, req.(*MobileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserId(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CreatUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatUserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreatUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_CreatUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreatUser(ctx, req.(*CreatUserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateUser(ctx, req.(*UpdateUserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CheckPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordCheckInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CheckPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_CheckPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CheckPassword(ctx, req.(*PasswordCheckInfo))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserList",
			Handler:    _User_GetUserList_Handler,
		},
		{
			MethodName: "GetUserMobile",
			Handler:    _User_GetUserMobile_Handler,
		},
		{
			MethodName: "GetUserId",
			Handler:    _User_GetUserId_Handler,
		},
		{
			MethodName: "CreatUser",
			Handler:    _User_CreatUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _User_UpdateUser_Handler,
		},
		{
			MethodName: "CheckPassword",
			Handler:    _User_CheckPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
