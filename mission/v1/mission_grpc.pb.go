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
	ConfigMission(ctx context.Context, in *ConfigMissionRequest, opts ...grpc.CallOption) (*ConfigMissionReply, error)
	CreateArea(ctx context.Context, in *CreateAreaRequest, opts ...grpc.CallOption) (*CreateAreaReply, error)
	UpdateArea(ctx context.Context, in *UpdateAreaRequest, opts ...grpc.CallOption) (*UpdateAreaReply, error)
	DeleteArea(ctx context.Context, in *DeleteAreaRequest, opts ...grpc.CallOption) (*DeleteAreaReply, error)
	ListArea(ctx context.Context, in *ListAreaRequest, opts ...grpc.CallOption) (*ListAreaReply, error)
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

func (c *missionClient) ConfigMission(ctx context.Context, in *ConfigMissionRequest, opts ...grpc.CallOption) (*ConfigMissionReply, error) {
	out := new(ConfigMissionReply)
	err := c.cc.Invoke(ctx, "/api.mission.v1.Mission/ConfigMission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *missionClient) CreateArea(ctx context.Context, in *CreateAreaRequest, opts ...grpc.CallOption) (*CreateAreaReply, error) {
	out := new(CreateAreaReply)
	err := c.cc.Invoke(ctx, "/api.mission.v1.Mission/CreateArea", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *missionClient) UpdateArea(ctx context.Context, in *UpdateAreaRequest, opts ...grpc.CallOption) (*UpdateAreaReply, error) {
	out := new(UpdateAreaReply)
	err := c.cc.Invoke(ctx, "/api.mission.v1.Mission/UpdateArea", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *missionClient) DeleteArea(ctx context.Context, in *DeleteAreaRequest, opts ...grpc.CallOption) (*DeleteAreaReply, error) {
	out := new(DeleteAreaReply)
	err := c.cc.Invoke(ctx, "/api.mission.v1.Mission/DeleteArea", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *missionClient) ListArea(ctx context.Context, in *ListAreaRequest, opts ...grpc.CallOption) (*ListAreaReply, error) {
	out := new(ListAreaReply)
	err := c.cc.Invoke(ctx, "/api.mission.v1.Mission/ListArea", in, out, opts...)
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
	ConfigMission(context.Context, *ConfigMissionRequest) (*ConfigMissionReply, error)
	CreateArea(context.Context, *CreateAreaRequest) (*CreateAreaReply, error)
	UpdateArea(context.Context, *UpdateAreaRequest) (*UpdateAreaReply, error)
	DeleteArea(context.Context, *DeleteAreaRequest) (*DeleteAreaReply, error)
	ListArea(context.Context, *ListAreaRequest) (*ListAreaReply, error)
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
func (UnimplementedMissionServer) ConfigMission(context.Context, *ConfigMissionRequest) (*ConfigMissionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigMission not implemented")
}
func (UnimplementedMissionServer) CreateArea(context.Context, *CreateAreaRequest) (*CreateAreaReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArea not implemented")
}
func (UnimplementedMissionServer) UpdateArea(context.Context, *UpdateAreaRequest) (*UpdateAreaReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArea not implemented")
}
func (UnimplementedMissionServer) DeleteArea(context.Context, *DeleteAreaRequest) (*DeleteAreaReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArea not implemented")
}
func (UnimplementedMissionServer) ListArea(context.Context, *ListAreaRequest) (*ListAreaReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArea not implemented")
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

func _Mission_ConfigMission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigMissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionServer).ConfigMission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.mission.v1.Mission/ConfigMission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionServer).ConfigMission(ctx, req.(*ConfigMissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mission_CreateArea_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAreaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionServer).CreateArea(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.mission.v1.Mission/CreateArea",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionServer).CreateArea(ctx, req.(*CreateAreaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mission_UpdateArea_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAreaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionServer).UpdateArea(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.mission.v1.Mission/UpdateArea",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionServer).UpdateArea(ctx, req.(*UpdateAreaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mission_DeleteArea_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAreaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionServer).DeleteArea(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.mission.v1.Mission/DeleteArea",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionServer).DeleteArea(ctx, req.(*DeleteAreaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mission_ListArea_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAreaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionServer).ListArea(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.mission.v1.Mission/ListArea",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionServer).ListArea(ctx, req.(*ListAreaRequest))
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
		{
			MethodName: "ConfigMission",
			Handler:    _Mission_ConfigMission_Handler,
		},
		{
			MethodName: "CreateArea",
			Handler:    _Mission_CreateArea_Handler,
		},
		{
			MethodName: "UpdateArea",
			Handler:    _Mission_UpdateArea_Handler,
		},
		{
			MethodName: "DeleteArea",
			Handler:    _Mission_DeleteArea_Handler,
		},
		{
			MethodName: "ListArea",
			Handler:    _Mission_ListArea_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mission/v1/mission.proto",
}
