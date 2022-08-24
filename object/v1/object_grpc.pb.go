// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: object/v1/object.proto

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

// ObjectClient is the client API for Object service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ObjectClient interface {
	CreateObject(ctx context.Context, in *CreateObjectRequest, opts ...grpc.CallOption) (*CreateObjectReply, error)
	UpdateObject(ctx context.Context, in *UpdateObjectRequest, opts ...grpc.CallOption) (*UpdateObjectReply, error)
	DeleteObject(ctx context.Context, in *DeleteObjectRequest, opts ...grpc.CallOption) (*DeleteObjectReply, error)
	GetObject(ctx context.Context, in *GetObjectRequest, opts ...grpc.CallOption) (*GetObjectReply, error)
	ListObject(ctx context.Context, in *ListObjectRequest, opts ...grpc.CallOption) (*ListObjectReply, error)
}

type objectClient struct {
	cc grpc.ClientConnInterface
}

func NewObjectClient(cc grpc.ClientConnInterface) ObjectClient {
	return &objectClient{cc}
}

func (c *objectClient) CreateObject(ctx context.Context, in *CreateObjectRequest, opts ...grpc.CallOption) (*CreateObjectReply, error) {
	out := new(CreateObjectReply)
	err := c.cc.Invoke(ctx, "/api.object.v1.Object/CreateObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectClient) UpdateObject(ctx context.Context, in *UpdateObjectRequest, opts ...grpc.CallOption) (*UpdateObjectReply, error) {
	out := new(UpdateObjectReply)
	err := c.cc.Invoke(ctx, "/api.object.v1.Object/UpdateObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectClient) DeleteObject(ctx context.Context, in *DeleteObjectRequest, opts ...grpc.CallOption) (*DeleteObjectReply, error) {
	out := new(DeleteObjectReply)
	err := c.cc.Invoke(ctx, "/api.object.v1.Object/DeleteObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectClient) GetObject(ctx context.Context, in *GetObjectRequest, opts ...grpc.CallOption) (*GetObjectReply, error) {
	out := new(GetObjectReply)
	err := c.cc.Invoke(ctx, "/api.object.v1.Object/GetObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectClient) ListObject(ctx context.Context, in *ListObjectRequest, opts ...grpc.CallOption) (*ListObjectReply, error) {
	out := new(ListObjectReply)
	err := c.cc.Invoke(ctx, "/api.object.v1.Object/ListObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ObjectServer is the server API for Object service.
// All implementations must embed UnimplementedObjectServer
// for forward compatibility
type ObjectServer interface {
	CreateObject(context.Context, *CreateObjectRequest) (*CreateObjectReply, error)
	UpdateObject(context.Context, *UpdateObjectRequest) (*UpdateObjectReply, error)
	DeleteObject(context.Context, *DeleteObjectRequest) (*DeleteObjectReply, error)
	GetObject(context.Context, *GetObjectRequest) (*GetObjectReply, error)
	ListObject(context.Context, *ListObjectRequest) (*ListObjectReply, error)
	mustEmbedUnimplementedObjectServer()
}

// UnimplementedObjectServer must be embedded to have forward compatible implementations.
type UnimplementedObjectServer struct {
}

func (UnimplementedObjectServer) CreateObject(context.Context, *CreateObjectRequest) (*CreateObjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateObject not implemented")
}
func (UnimplementedObjectServer) UpdateObject(context.Context, *UpdateObjectRequest) (*UpdateObjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateObject not implemented")
}
func (UnimplementedObjectServer) DeleteObject(context.Context, *DeleteObjectRequest) (*DeleteObjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteObject not implemented")
}
func (UnimplementedObjectServer) GetObject(context.Context, *GetObjectRequest) (*GetObjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObject not implemented")
}
func (UnimplementedObjectServer) ListObject(context.Context, *ListObjectRequest) (*ListObjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListObject not implemented")
}
func (UnimplementedObjectServer) mustEmbedUnimplementedObjectServer() {}

// UnsafeObjectServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ObjectServer will
// result in compilation errors.
type UnsafeObjectServer interface {
	mustEmbedUnimplementedObjectServer()
}

func RegisterObjectServer(s grpc.ServiceRegistrar, srv ObjectServer) {
	s.RegisterService(&Object_ServiceDesc, srv)
}

func _Object_CreateObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectServer).CreateObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.object.v1.Object/CreateObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectServer).CreateObject(ctx, req.(*CreateObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Object_UpdateObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectServer).UpdateObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.object.v1.Object/UpdateObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectServer).UpdateObject(ctx, req.(*UpdateObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Object_DeleteObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectServer).DeleteObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.object.v1.Object/DeleteObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectServer).DeleteObject(ctx, req.(*DeleteObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Object_GetObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectServer).GetObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.object.v1.Object/GetObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectServer).GetObject(ctx, req.(*GetObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Object_ListObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectServer).ListObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.object.v1.Object/ListObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectServer).ListObject(ctx, req.(*ListObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Object_ServiceDesc is the grpc.ServiceDesc for Object service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Object_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.object.v1.Object",
	HandlerType: (*ObjectServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateObject",
			Handler:    _Object_CreateObject_Handler,
		},
		{
			MethodName: "UpdateObject",
			Handler:    _Object_UpdateObject_Handler,
		},
		{
			MethodName: "DeleteObject",
			Handler:    _Object_DeleteObject_Handler,
		},
		{
			MethodName: "GetObject",
			Handler:    _Object_GetObject_Handler,
		},
		{
			MethodName: "ListObject",
			Handler:    _Object_ListObject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "object/v1/object.proto",
}
