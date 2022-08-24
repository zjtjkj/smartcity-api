// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: api/home-config/v1/home-config.proto

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

// HomeConfigClient is the client API for HomeConfig service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HomeConfigClient interface {
	SaveHomeConfig(ctx context.Context, in *SaveHomeConfigRequest, opts ...grpc.CallOption) (*SaveHomeConfigReply, error)
	GetHomeConfig(ctx context.Context, in *GetHomeConfigRequest, opts ...grpc.CallOption) (*GetHomeConfigReply, error)
	GetOptions(ctx context.Context, in *GetOptionsRequest, opts ...grpc.CallOption) (*GetOptionsReply, error)
	GetOptionById(ctx context.Context, in *GetOptionByIdRequest, opts ...grpc.CallOption) (*GetOptionByIdReply, error)
}

type homeConfigClient struct {
	cc grpc.ClientConnInterface
}

func NewHomeConfigClient(cc grpc.ClientConnInterface) HomeConfigClient {
	return &homeConfigClient{cc}
}

func (c *homeConfigClient) SaveHomeConfig(ctx context.Context, in *SaveHomeConfigRequest, opts ...grpc.CallOption) (*SaveHomeConfigReply, error) {
	out := new(SaveHomeConfigReply)
	err := c.cc.Invoke(ctx, "/api.home.v1.HomeConfig/SaveHomeConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homeConfigClient) GetHomeConfig(ctx context.Context, in *GetHomeConfigRequest, opts ...grpc.CallOption) (*GetHomeConfigReply, error) {
	out := new(GetHomeConfigReply)
	err := c.cc.Invoke(ctx, "/api.home.v1.HomeConfig/GetHomeConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homeConfigClient) GetOptions(ctx context.Context, in *GetOptionsRequest, opts ...grpc.CallOption) (*GetOptionsReply, error) {
	out := new(GetOptionsReply)
	err := c.cc.Invoke(ctx, "/api.home.v1.HomeConfig/GetOptions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homeConfigClient) GetOptionById(ctx context.Context, in *GetOptionByIdRequest, opts ...grpc.CallOption) (*GetOptionByIdReply, error) {
	out := new(GetOptionByIdReply)
	err := c.cc.Invoke(ctx, "/api.home.v1.HomeConfig/GetOptionById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HomeConfigServer is the server API for HomeConfig service.
// All implementations must embed UnimplementedHomeConfigServer
// for forward compatibility
type HomeConfigServer interface {
	SaveHomeConfig(context.Context, *SaveHomeConfigRequest) (*SaveHomeConfigReply, error)
	GetHomeConfig(context.Context, *GetHomeConfigRequest) (*GetHomeConfigReply, error)
	GetOptions(context.Context, *GetOptionsRequest) (*GetOptionsReply, error)
	GetOptionById(context.Context, *GetOptionByIdRequest) (*GetOptionByIdReply, error)
	mustEmbedUnimplementedHomeConfigServer()
}

// UnimplementedHomeConfigServer must be embedded to have forward compatible implementations.
type UnimplementedHomeConfigServer struct {
}

func (UnimplementedHomeConfigServer) SaveHomeConfig(context.Context, *SaveHomeConfigRequest) (*SaveHomeConfigReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveHomeConfig not implemented")
}
func (UnimplementedHomeConfigServer) GetHomeConfig(context.Context, *GetHomeConfigRequest) (*GetHomeConfigReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHomeConfig not implemented")
}
func (UnimplementedHomeConfigServer) GetOptions(context.Context, *GetOptionsRequest) (*GetOptionsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOptions not implemented")
}
func (UnimplementedHomeConfigServer) GetOptionById(context.Context, *GetOptionByIdRequest) (*GetOptionByIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOptionById not implemented")
}
func (UnimplementedHomeConfigServer) mustEmbedUnimplementedHomeConfigServer() {}

// UnsafeHomeConfigServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HomeConfigServer will
// result in compilation errors.
type UnsafeHomeConfigServer interface {
	mustEmbedUnimplementedHomeConfigServer()
}

func RegisterHomeConfigServer(s grpc.ServiceRegistrar, srv HomeConfigServer) {
	s.RegisterService(&HomeConfig_ServiceDesc, srv)
}

func _HomeConfig_SaveHomeConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveHomeConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeConfigServer).SaveHomeConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.home.v1.HomeConfig/SaveHomeConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeConfigServer).SaveHomeConfig(ctx, req.(*SaveHomeConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HomeConfig_GetHomeConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHomeConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeConfigServer).GetHomeConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.home.v1.HomeConfig/GetHomeConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeConfigServer).GetHomeConfig(ctx, req.(*GetHomeConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HomeConfig_GetOptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOptionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeConfigServer).GetOptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.home.v1.HomeConfig/GetOptions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeConfigServer).GetOptions(ctx, req.(*GetOptionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HomeConfig_GetOptionById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOptionByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeConfigServer).GetOptionById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.home.v1.HomeConfig/GetOptionById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeConfigServer).GetOptionById(ctx, req.(*GetOptionByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HomeConfig_ServiceDesc is the grpc.ServiceDesc for HomeConfig service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HomeConfig_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.home.v1.HomeConfig",
	HandlerType: (*HomeConfigServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveHomeConfig",
			Handler:    _HomeConfig_SaveHomeConfig_Handler,
		},
		{
			MethodName: "GetHomeConfig",
			Handler:    _HomeConfig_GetHomeConfig_Handler,
		},
		{
			MethodName: "GetOptions",
			Handler:    _HomeConfig_GetOptions_Handler,
		},
		{
			MethodName: "GetOptionById",
			Handler:    _HomeConfig_GetOptionById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/home-config/v1/home-config.proto",
}
