// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: api/retrieval/v1/retrieval.proto

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

// RetrievalClient is the client API for Retrieval service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RetrievalClient interface {
	FindAlerts(ctx context.Context, in *FindAlertsRequest, opts ...grpc.CallOption) (*FindAlertsReply, error)
	GetAlert(ctx context.Context, in *GetAlertRequest, opts ...grpc.CallOption) (*GetAlertReply, error)
	GetImage(ctx context.Context, in *GetImageRequest, opts ...grpc.CallOption) (*GetImageReply, error)
	DeleteAlert(ctx context.Context, in *DeleteAlertRequest, opts ...grpc.CallOption) (*DeleteAlertReply, error)
}

type retrievalClient struct {
	cc grpc.ClientConnInterface
}

func NewRetrievalClient(cc grpc.ClientConnInterface) RetrievalClient {
	return &retrievalClient{cc}
}

func (c *retrievalClient) FindAlerts(ctx context.Context, in *FindAlertsRequest, opts ...grpc.CallOption) (*FindAlertsReply, error) {
	out := new(FindAlertsReply)
	err := c.cc.Invoke(ctx, "/api.retrieval.Retrieval/FindAlerts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *retrievalClient) GetAlert(ctx context.Context, in *GetAlertRequest, opts ...grpc.CallOption) (*GetAlertReply, error) {
	out := new(GetAlertReply)
	err := c.cc.Invoke(ctx, "/api.retrieval.Retrieval/GetAlert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *retrievalClient) GetImage(ctx context.Context, in *GetImageRequest, opts ...grpc.CallOption) (*GetImageReply, error) {
	out := new(GetImageReply)
	err := c.cc.Invoke(ctx, "/api.retrieval.Retrieval/GetImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *retrievalClient) DeleteAlert(ctx context.Context, in *DeleteAlertRequest, opts ...grpc.CallOption) (*DeleteAlertReply, error) {
	out := new(DeleteAlertReply)
	err := c.cc.Invoke(ctx, "/api.retrieval.Retrieval/DeleteAlert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RetrievalServer is the server API for Retrieval service.
// All implementations must embed UnimplementedRetrievalServer
// for forward compatibility
type RetrievalServer interface {
	FindAlerts(context.Context, *FindAlertsRequest) (*FindAlertsReply, error)
	GetAlert(context.Context, *GetAlertRequest) (*GetAlertReply, error)
	GetImage(context.Context, *GetImageRequest) (*GetImageReply, error)
	DeleteAlert(context.Context, *DeleteAlertRequest) (*DeleteAlertReply, error)
	mustEmbedUnimplementedRetrievalServer()
}

// UnimplementedRetrievalServer must be embedded to have forward compatible implementations.
type UnimplementedRetrievalServer struct {
}

func (UnimplementedRetrievalServer) FindAlerts(context.Context, *FindAlertsRequest) (*FindAlertsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAlerts not implemented")
}
func (UnimplementedRetrievalServer) GetAlert(context.Context, *GetAlertRequest) (*GetAlertReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAlert not implemented")
}
func (UnimplementedRetrievalServer) GetImage(context.Context, *GetImageRequest) (*GetImageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImage not implemented")
}
func (UnimplementedRetrievalServer) DeleteAlert(context.Context, *DeleteAlertRequest) (*DeleteAlertReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAlert not implemented")
}
func (UnimplementedRetrievalServer) mustEmbedUnimplementedRetrievalServer() {}

// UnsafeRetrievalServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RetrievalServer will
// result in compilation errors.
type UnsafeRetrievalServer interface {
	mustEmbedUnimplementedRetrievalServer()
}

func RegisterRetrievalServer(s grpc.ServiceRegistrar, srv RetrievalServer) {
	s.RegisterService(&Retrieval_ServiceDesc, srv)
}

func _Retrieval_FindAlerts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAlertsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RetrievalServer).FindAlerts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.retrieval.Retrieval/FindAlerts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RetrievalServer).FindAlerts(ctx, req.(*FindAlertsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Retrieval_GetAlert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAlertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RetrievalServer).GetAlert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.retrieval.Retrieval/GetAlert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RetrievalServer).GetAlert(ctx, req.(*GetAlertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Retrieval_GetImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RetrievalServer).GetImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.retrieval.Retrieval/GetImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RetrievalServer).GetImage(ctx, req.(*GetImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Retrieval_DeleteAlert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAlertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RetrievalServer).DeleteAlert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.retrieval.Retrieval/DeleteAlert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RetrievalServer).DeleteAlert(ctx, req.(*DeleteAlertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Retrieval_ServiceDesc is the grpc.ServiceDesc for Retrieval service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Retrieval_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.retrieval.Retrieval",
	HandlerType: (*RetrievalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindAlerts",
			Handler:    _Retrieval_FindAlerts_Handler,
		},
		{
			MethodName: "GetAlert",
			Handler:    _Retrieval_GetAlert_Handler,
		},
		{
			MethodName: "GetImage",
			Handler:    _Retrieval_GetImage_Handler,
		},
		{
			MethodName: "DeleteAlert",
			Handler:    _Retrieval_DeleteAlert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/retrieval/v1/retrieval.proto",
}