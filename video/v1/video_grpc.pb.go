// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: api/video/v1/video.proto

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

// VideoClient is the client API for Video service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoClient interface {
	FindDevices(ctx context.Context, in *FindDevicesRequest, opts ...grpc.CallOption) (*FindDevicesReply, error)
	FindChannels(ctx context.Context, in *FindChannelsRequest, opts ...grpc.CallOption) (*FindChannelsReply, error)
	StartPlay(ctx context.Context, in *StartPlayRequest, opts ...grpc.CallOption) (*StartPlayReply, error)
	StopPlay(ctx context.Context, in *StopPlayRequest, opts ...grpc.CallOption) (*StopPlayReply, error)
	StartProxy(ctx context.Context, in *StartProxyRequest, opts ...grpc.CallOption) (*StartProxyReply, error)
	StopProxy(ctx context.Context, in *StopProxyRequest, opts ...grpc.CallOption) (*StopProxyReply, error)
	DeleteProxy(ctx context.Context, in *DeleteProxyRequest, opts ...grpc.CallOption) (*DeleteProxyReply, error)
}

type videoClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoClient(cc grpc.ClientConnInterface) VideoClient {
	return &videoClient{cc}
}

func (c *videoClient) FindDevices(ctx context.Context, in *FindDevicesRequest, opts ...grpc.CallOption) (*FindDevicesReply, error) {
	out := new(FindDevicesReply)
	err := c.cc.Invoke(ctx, "/api.video.v1.Video/FindDevices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) FindChannels(ctx context.Context, in *FindChannelsRequest, opts ...grpc.CallOption) (*FindChannelsReply, error) {
	out := new(FindChannelsReply)
	err := c.cc.Invoke(ctx, "/api.video.v1.Video/FindChannels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) StartPlay(ctx context.Context, in *StartPlayRequest, opts ...grpc.CallOption) (*StartPlayReply, error) {
	out := new(StartPlayReply)
	err := c.cc.Invoke(ctx, "/api.video.v1.Video/StartPlay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) StopPlay(ctx context.Context, in *StopPlayRequest, opts ...grpc.CallOption) (*StopPlayReply, error) {
	out := new(StopPlayReply)
	err := c.cc.Invoke(ctx, "/api.video.v1.Video/StopPlay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) StartProxy(ctx context.Context, in *StartProxyRequest, opts ...grpc.CallOption) (*StartProxyReply, error) {
	out := new(StartProxyReply)
	err := c.cc.Invoke(ctx, "/api.video.v1.Video/StartProxy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) StopProxy(ctx context.Context, in *StopProxyRequest, opts ...grpc.CallOption) (*StopProxyReply, error) {
	out := new(StopProxyReply)
	err := c.cc.Invoke(ctx, "/api.video.v1.Video/StopProxy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) DeleteProxy(ctx context.Context, in *DeleteProxyRequest, opts ...grpc.CallOption) (*DeleteProxyReply, error) {
	out := new(DeleteProxyReply)
	err := c.cc.Invoke(ctx, "/api.video.v1.Video/DeleteProxy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoServer is the server API for Video service.
// All implementations must embed UnimplementedVideoServer
// for forward compatibility
type VideoServer interface {
	FindDevices(context.Context, *FindDevicesRequest) (*FindDevicesReply, error)
	FindChannels(context.Context, *FindChannelsRequest) (*FindChannelsReply, error)
	StartPlay(context.Context, *StartPlayRequest) (*StartPlayReply, error)
	StopPlay(context.Context, *StopPlayRequest) (*StopPlayReply, error)
	StartProxy(context.Context, *StartProxyRequest) (*StartProxyReply, error)
	StopProxy(context.Context, *StopProxyRequest) (*StopProxyReply, error)
	DeleteProxy(context.Context, *DeleteProxyRequest) (*DeleteProxyReply, error)
	mustEmbedUnimplementedVideoServer()
}

// UnimplementedVideoServer must be embedded to have forward compatible implementations.
type UnimplementedVideoServer struct {
}

func (UnimplementedVideoServer) FindDevices(context.Context, *FindDevicesRequest) (*FindDevicesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindDevices not implemented")
}
func (UnimplementedVideoServer) FindChannels(context.Context, *FindChannelsRequest) (*FindChannelsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindChannels not implemented")
}
func (UnimplementedVideoServer) StartPlay(context.Context, *StartPlayRequest) (*StartPlayReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartPlay not implemented")
}
func (UnimplementedVideoServer) StopPlay(context.Context, *StopPlayRequest) (*StopPlayReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopPlay not implemented")
}
func (UnimplementedVideoServer) StartProxy(context.Context, *StartProxyRequest) (*StartProxyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartProxy not implemented")
}
func (UnimplementedVideoServer) StopProxy(context.Context, *StopProxyRequest) (*StopProxyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopProxy not implemented")
}
func (UnimplementedVideoServer) DeleteProxy(context.Context, *DeleteProxyRequest) (*DeleteProxyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProxy not implemented")
}
func (UnimplementedVideoServer) mustEmbedUnimplementedVideoServer() {}

// UnsafeVideoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoServer will
// result in compilation errors.
type UnsafeVideoServer interface {
	mustEmbedUnimplementedVideoServer()
}

func RegisterVideoServer(s grpc.ServiceRegistrar, srv VideoServer) {
	s.RegisterService(&Video_ServiceDesc, srv)
}

func _Video_FindDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindDevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).FindDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.video.v1.Video/FindDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).FindDevices(ctx, req.(*FindDevicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_FindChannels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindChannelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).FindChannels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.video.v1.Video/FindChannels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).FindChannels(ctx, req.(*FindChannelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_StartPlay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartPlayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).StartPlay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.video.v1.Video/StartPlay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).StartPlay(ctx, req.(*StartPlayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_StopPlay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopPlayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).StopPlay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.video.v1.Video/StopPlay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).StopPlay(ctx, req.(*StopPlayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_StartProxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartProxyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).StartProxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.video.v1.Video/StartProxy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).StartProxy(ctx, req.(*StartProxyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_StopProxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopProxyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).StopProxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.video.v1.Video/StopProxy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).StopProxy(ctx, req.(*StopProxyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_DeleteProxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProxyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).DeleteProxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.video.v1.Video/DeleteProxy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).DeleteProxy(ctx, req.(*DeleteProxyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Video_ServiceDesc is the grpc.ServiceDesc for Video service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Video_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.video.v1.Video",
	HandlerType: (*VideoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindDevices",
			Handler:    _Video_FindDevices_Handler,
		},
		{
			MethodName: "FindChannels",
			Handler:    _Video_FindChannels_Handler,
		},
		{
			MethodName: "StartPlay",
			Handler:    _Video_StartPlay_Handler,
		},
		{
			MethodName: "StopPlay",
			Handler:    _Video_StopPlay_Handler,
		},
		{
			MethodName: "StartProxy",
			Handler:    _Video_StartProxy_Handler,
		},
		{
			MethodName: "StopProxy",
			Handler:    _Video_StopProxy_Handler,
		},
		{
			MethodName: "DeleteProxy",
			Handler:    _Video_DeleteProxy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/video/v1/video.proto",
}
