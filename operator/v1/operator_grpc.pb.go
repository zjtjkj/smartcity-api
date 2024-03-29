// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: operator/v1/operator.proto

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

// OperatorClient is the client API for Operator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OperatorClient interface {
	GetMission(ctx context.Context, in *GetMissionRequest, opts ...grpc.CallOption) (*GetMissionReply, error)
	RefreshMission(ctx context.Context, in *RefreshMissionRequest, opts ...grpc.CallOption) (*RefreshMissionReply, error)
}

type operatorClient struct {
	cc grpc.ClientConnInterface
}

func NewOperatorClient(cc grpc.ClientConnInterface) OperatorClient {
	return &operatorClient{cc}
}

func (c *operatorClient) GetMission(ctx context.Context, in *GetMissionRequest, opts ...grpc.CallOption) (*GetMissionReply, error) {
	out := new(GetMissionReply)
	err := c.cc.Invoke(ctx, "/api.operator.v1.Operator/GetMission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorClient) RefreshMission(ctx context.Context, in *RefreshMissionRequest, opts ...grpc.CallOption) (*RefreshMissionReply, error) {
	out := new(RefreshMissionReply)
	err := c.cc.Invoke(ctx, "/api.operator.v1.Operator/RefreshMission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperatorServer is the server API for Operator service.
// All implementations must embed UnimplementedOperatorServer
// for forward compatibility
type OperatorServer interface {
	GetMission(context.Context, *GetMissionRequest) (*GetMissionReply, error)
	RefreshMission(context.Context, *RefreshMissionRequest) (*RefreshMissionReply, error)
	mustEmbedUnimplementedOperatorServer()
}

// UnimplementedOperatorServer must be embedded to have forward compatible implementations.
type UnimplementedOperatorServer struct {
}

func (UnimplementedOperatorServer) GetMission(context.Context, *GetMissionRequest) (*GetMissionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMission not implemented")
}
func (UnimplementedOperatorServer) RefreshMission(context.Context, *RefreshMissionRequest) (*RefreshMissionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshMission not implemented")
}
func (UnimplementedOperatorServer) mustEmbedUnimplementedOperatorServer() {}

// UnsafeOperatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OperatorServer will
// result in compilation errors.
type UnsafeOperatorServer interface {
	mustEmbedUnimplementedOperatorServer()
}

func RegisterOperatorServer(s grpc.ServiceRegistrar, srv OperatorServer) {
	s.RegisterService(&Operator_ServiceDesc, srv)
}

func _Operator_GetMission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServer).GetMission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.operator.v1.Operator/GetMission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServer).GetMission(ctx, req.(*GetMissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operator_RefreshMission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshMissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServer).RefreshMission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.operator.v1.Operator/RefreshMission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServer).RefreshMission(ctx, req.(*RefreshMissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Operator_ServiceDesc is the grpc.ServiceDesc for Operator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Operator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.operator.v1.Operator",
	HandlerType: (*OperatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMission",
			Handler:    _Operator_GetMission_Handler,
		},
		{
			MethodName: "RefreshMission",
			Handler:    _Operator_RefreshMission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "operator/v1/operator.proto",
}
