// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: api/user-domain/v1/user-domain.proto

package v1

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

// UserDomainClient is the client API for UserDomain service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserDomainClient interface {
	CreateUserDomain(ctx context.Context, in *CreateUserDomainRequest, opts ...grpc.CallOption) (*CreateUserDomainReply, error)
	UpdateUserDomain(ctx context.Context, in *UpdateUserDomainRequest, opts ...grpc.CallOption) (*UpdateUserDomainReply, error)
	DeleteUserDomain(ctx context.Context, in *DeleteUserDomainRequest, opts ...grpc.CallOption) (*DeleteUserDomainReply, error)
	GetUserDomain(ctx context.Context, in *GetUserDomainRequest, opts ...grpc.CallOption) (*GetUserDomainReply, error)
	FindUserDomain(ctx context.Context, in *FindUserDomainRequest, opts ...grpc.CallOption) (*FindUserDomainReply, error)
	ListUserDomain(ctx context.Context, in *ListUserDomainRequest, opts ...grpc.CallOption) (*ListUserDomainReply, error)
}

type userDomainClient struct {
	cc grpc.ClientConnInterface
}

func NewUserDomainClient(cc grpc.ClientConnInterface) UserDomainClient {
	return &userDomainClient{cc}
}

func (c *userDomainClient) CreateUserDomain(ctx context.Context, in *CreateUserDomainRequest, opts ...grpc.CallOption) (*CreateUserDomainReply, error) {
	out := new(CreateUserDomainReply)
	err := c.cc.Invoke(ctx, "/api.userdomain.v1.UserDomain/CreateUserDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDomainClient) UpdateUserDomain(ctx context.Context, in *UpdateUserDomainRequest, opts ...grpc.CallOption) (*UpdateUserDomainReply, error) {
	out := new(UpdateUserDomainReply)
	err := c.cc.Invoke(ctx, "/api.userdomain.v1.UserDomain/UpdateUserDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDomainClient) DeleteUserDomain(ctx context.Context, in *DeleteUserDomainRequest, opts ...grpc.CallOption) (*DeleteUserDomainReply, error) {
	out := new(DeleteUserDomainReply)
	err := c.cc.Invoke(ctx, "/api.userdomain.v1.UserDomain/DeleteUserDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDomainClient) GetUserDomain(ctx context.Context, in *GetUserDomainRequest, opts ...grpc.CallOption) (*GetUserDomainReply, error) {
	out := new(GetUserDomainReply)
	err := c.cc.Invoke(ctx, "/api.userdomain.v1.UserDomain/GetUserDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDomainClient) FindUserDomain(ctx context.Context, in *FindUserDomainRequest, opts ...grpc.CallOption) (*FindUserDomainReply, error) {
	out := new(FindUserDomainReply)
	err := c.cc.Invoke(ctx, "/api.userdomain.v1.UserDomain/FindUserDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDomainClient) ListUserDomain(ctx context.Context, in *ListUserDomainRequest, opts ...grpc.CallOption) (*ListUserDomainReply, error) {
	out := new(ListUserDomainReply)
	err := c.cc.Invoke(ctx, "/api.userdomain.v1.UserDomain/ListUserDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserDomainServer is the server API for UserDomain service.
// All implementations must embed UnimplementedUserDomainServer
// for forward compatibility
type UserDomainServer interface {
	CreateUserDomain(context.Context, *CreateUserDomainRequest) (*CreateUserDomainReply, error)
	UpdateUserDomain(context.Context, *UpdateUserDomainRequest) (*UpdateUserDomainReply, error)
	DeleteUserDomain(context.Context, *DeleteUserDomainRequest) (*DeleteUserDomainReply, error)
	GetUserDomain(context.Context, *GetUserDomainRequest) (*GetUserDomainReply, error)
	FindUserDomain(context.Context, *FindUserDomainRequest) (*FindUserDomainReply, error)
	ListUserDomain(context.Context, *ListUserDomainRequest) (*ListUserDomainReply, error)
	mustEmbedUnimplementedUserDomainServer()
}

// UnimplementedUserDomainServer must be embedded to have forward compatible implementations.
type UnimplementedUserDomainServer struct {
}

func (UnimplementedUserDomainServer) CreateUserDomain(context.Context, *CreateUserDomainRequest) (*CreateUserDomainReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserDomain not implemented")
}
func (UnimplementedUserDomainServer) UpdateUserDomain(context.Context, *UpdateUserDomainRequest) (*UpdateUserDomainReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserDomain not implemented")
}
func (UnimplementedUserDomainServer) DeleteUserDomain(context.Context, *DeleteUserDomainRequest) (*DeleteUserDomainReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserDomain not implemented")
}
func (UnimplementedUserDomainServer) GetUserDomain(context.Context, *GetUserDomainRequest) (*GetUserDomainReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserDomain not implemented")
}
func (UnimplementedUserDomainServer) FindUserDomain(context.Context, *FindUserDomainRequest) (*FindUserDomainReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUserDomain not implemented")
}
func (UnimplementedUserDomainServer) ListUserDomain(context.Context, *ListUserDomainRequest) (*ListUserDomainReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserDomain not implemented")
}
func (UnimplementedUserDomainServer) mustEmbedUnimplementedUserDomainServer() {}

// UnsafeUserDomainServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserDomainServer will
// result in compilation errors.
type UnsafeUserDomainServer interface {
	mustEmbedUnimplementedUserDomainServer()
}

func RegisterUserDomainServer(s grpc.ServiceRegistrar, srv UserDomainServer) {
	s.RegisterService(&UserDomain_ServiceDesc, srv)
}

func _UserDomain_CreateUserDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDomainServer).CreateUserDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.userdomain.v1.UserDomain/CreateUserDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDomainServer).CreateUserDomain(ctx, req.(*CreateUserDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserDomain_UpdateUserDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDomainServer).UpdateUserDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.userdomain.v1.UserDomain/UpdateUserDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDomainServer).UpdateUserDomain(ctx, req.(*UpdateUserDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserDomain_DeleteUserDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDomainServer).DeleteUserDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.userdomain.v1.UserDomain/DeleteUserDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDomainServer).DeleteUserDomain(ctx, req.(*DeleteUserDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserDomain_GetUserDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDomainServer).GetUserDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.userdomain.v1.UserDomain/GetUserDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDomainServer).GetUserDomain(ctx, req.(*GetUserDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserDomain_FindUserDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindUserDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDomainServer).FindUserDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.userdomain.v1.UserDomain/FindUserDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDomainServer).FindUserDomain(ctx, req.(*FindUserDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserDomain_ListUserDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDomainServer).ListUserDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.userdomain.v1.UserDomain/ListUserDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDomainServer).ListUserDomain(ctx, req.(*ListUserDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserDomain_ServiceDesc is the grpc.ServiceDesc for UserDomain service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserDomain_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.userdomain.v1.UserDomain",
	HandlerType: (*UserDomainServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUserDomain",
			Handler:    _UserDomain_CreateUserDomain_Handler,
		},
		{
			MethodName: "UpdateUserDomain",
			Handler:    _UserDomain_UpdateUserDomain_Handler,
		},
		{
			MethodName: "DeleteUserDomain",
			Handler:    _UserDomain_DeleteUserDomain_Handler,
		},
		{
			MethodName: "GetUserDomain",
			Handler:    _UserDomain_GetUserDomain_Handler,
		},
		{
			MethodName: "FindUserDomain",
			Handler:    _UserDomain_FindUserDomain_Handler,
		},
		{
			MethodName: "ListUserDomain",
			Handler:    _UserDomain_ListUserDomain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/user-domain/v1/user-domain.proto",
}
