// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: mission/v1/mission.proto

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

// MissionClient is the client API for Mission service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MissionClient interface {
	CreateMission(ctx context.Context, in *CreateMissionRequest, opts ...grpc.CallOption) (*CreateMissionReply, error)
	UpdateMission(ctx context.Context, in *UpdateMissionRequest, opts ...grpc.CallOption) (*UpdateMissionReply, error)
	DeleteMission(ctx context.Context, in *DeleteMissionRequest, opts ...grpc.CallOption) (*DeleteMissionReply, error)
	GetMission(ctx context.Context, in *GetMissionRequest, opts ...grpc.CallOption) (*GetMissionReply, error)
	ListMissionByCameraAndPreset(ctx context.Context, in *ListMissionByCameraAndPresetRequest, opts ...grpc.CallOption) (*ListMissionByCameraAndPresetReply, error)
	ListMissionByCamera(ctx context.Context, in *ListMissionByCameraRequest, opts ...grpc.CallOption) (*ListMissionByCameraReply, error)
}

type missionClient struct {
	cc grpc.ClientConnInterface
}

func NewMissionClient(cc grpc.ClientConnInterface) MissionClient {
	return &missionClient{cc}
}

func (c *missionClient) CreateMission(ctx context.Context, in *CreateMissionRequest, opts ...grpc.CallOption) (*CreateMissionReply, error) {
	out := new(CreateMissionReply)
	err := c.cc.Invoke(ctx, "/api.mission.v1.Mission/CreateMission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *missionClient) UpdateMission(ctx context.Context, in *UpdateMissionRequest, opts ...grpc.CallOption) (*UpdateMissionReply, error) {
	out := new(UpdateMissionReply)
	err := c.cc.Invoke(ctx, "/api.mission.v1.Mission/UpdateMission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *missionClient) DeleteMission(ctx context.Context, in *DeleteMissionRequest, opts ...grpc.CallOption) (*DeleteMissionReply, error) {
	out := new(DeleteMissionReply)
	err := c.cc.Invoke(ctx, "/api.mission.v1.Mission/DeleteMission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *missionClient) GetMission(ctx context.Context, in *GetMissionRequest, opts ...grpc.CallOption) (*GetMissionReply, error) {
	out := new(GetMissionReply)
	err := c.cc.Invoke(ctx, "/api.mission.v1.Mission/GetMission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *missionClient) ListMissionByCameraAndPreset(ctx context.Context, in *ListMissionByCameraAndPresetRequest, opts ...grpc.CallOption) (*ListMissionByCameraAndPresetReply, error) {
	out := new(ListMissionByCameraAndPresetReply)
	err := c.cc.Invoke(ctx, "/api.mission.v1.Mission/ListMissionByCameraAndPreset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *missionClient) ListMissionByCamera(ctx context.Context, in *ListMissionByCameraRequest, opts ...grpc.CallOption) (*ListMissionByCameraReply, error) {
	out := new(ListMissionByCameraReply)
	err := c.cc.Invoke(ctx, "/api.mission.v1.Mission/ListMissionByCamera", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MissionServer is the server API for Mission service.
// All implementations must embed UnimplementedMissionServer
// for forward compatibility
type MissionServer interface {
	CreateMission(context.Context, *CreateMissionRequest) (*CreateMissionReply, error)
	UpdateMission(context.Context, *UpdateMissionRequest) (*UpdateMissionReply, error)
	DeleteMission(context.Context, *DeleteMissionRequest) (*DeleteMissionReply, error)
	GetMission(context.Context, *GetMissionRequest) (*GetMissionReply, error)
	ListMissionByCameraAndPreset(context.Context, *ListMissionByCameraAndPresetRequest) (*ListMissionByCameraAndPresetReply, error)
	ListMissionByCamera(context.Context, *ListMissionByCameraRequest) (*ListMissionByCameraReply, error)
	mustEmbedUnimplementedMissionServer()
}

// UnimplementedMissionServer must be embedded to have forward compatible implementations.
type UnimplementedMissionServer struct {
}

func (UnimplementedMissionServer) CreateMission(context.Context, *CreateMissionRequest) (*CreateMissionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMission not implemented")
}
func (UnimplementedMissionServer) UpdateMission(context.Context, *UpdateMissionRequest) (*UpdateMissionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMission not implemented")
}
func (UnimplementedMissionServer) DeleteMission(context.Context, *DeleteMissionRequest) (*DeleteMissionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMission not implemented")
}
func (UnimplementedMissionServer) GetMission(context.Context, *GetMissionRequest) (*GetMissionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMission not implemented")
}
func (UnimplementedMissionServer) ListMissionByCameraAndPreset(context.Context, *ListMissionByCameraAndPresetRequest) (*ListMissionByCameraAndPresetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMissionByCameraAndPreset not implemented")
}
func (UnimplementedMissionServer) ListMissionByCamera(context.Context, *ListMissionByCameraRequest) (*ListMissionByCameraReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMissionByCamera not implemented")
}
func (UnimplementedMissionServer) mustEmbedUnimplementedMissionServer() {}

// UnsafeMissionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MissionServer will
// result in compilation errors.
type UnsafeMissionServer interface {
	mustEmbedUnimplementedMissionServer()
}

func RegisterMissionServer(s grpc.ServiceRegistrar, srv MissionServer) {
	s.RegisterService(&Mission_ServiceDesc, srv)
}

func _Mission_CreateMission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionServer).CreateMission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.mission.v1.Mission/CreateMission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionServer).CreateMission(ctx, req.(*CreateMissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mission_UpdateMission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionServer).UpdateMission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.mission.v1.Mission/UpdateMission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionServer).UpdateMission(ctx, req.(*UpdateMissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mission_DeleteMission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionServer).DeleteMission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.mission.v1.Mission/DeleteMission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionServer).DeleteMission(ctx, req.(*DeleteMissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mission_GetMission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionServer).GetMission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.mission.v1.Mission/GetMission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionServer).GetMission(ctx, req.(*GetMissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mission_ListMissionByCameraAndPreset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMissionByCameraAndPresetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionServer).ListMissionByCameraAndPreset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.mission.v1.Mission/ListMissionByCameraAndPreset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionServer).ListMissionByCameraAndPreset(ctx, req.(*ListMissionByCameraAndPresetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mission_ListMissionByCamera_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMissionByCameraRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionServer).ListMissionByCamera(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.mission.v1.Mission/ListMissionByCamera",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionServer).ListMissionByCamera(ctx, req.(*ListMissionByCameraRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Mission_ServiceDesc is the grpc.ServiceDesc for Mission service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mission_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.mission.v1.Mission",
	HandlerType: (*MissionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMission",
			Handler:    _Mission_CreateMission_Handler,
		},
		{
			MethodName: "UpdateMission",
			Handler:    _Mission_UpdateMission_Handler,
		},
		{
			MethodName: "DeleteMission",
			Handler:    _Mission_DeleteMission_Handler,
		},
		{
			MethodName: "GetMission",
			Handler:    _Mission_GetMission_Handler,
		},
		{
			MethodName: "ListMissionByCameraAndPreset",
			Handler:    _Mission_ListMissionByCameraAndPreset_Handler,
		},
		{
			MethodName: "ListMissionByCamera",
			Handler:    _Mission_ListMissionByCamera_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mission/v1/mission.proto",
}
