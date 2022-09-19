// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: settings/v1/settings.proto

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

// SettingsClient is the client API for Settings service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SettingsClient interface {
	CreateMapConfig(ctx context.Context, in *CreateMapConfigRequest, opts ...grpc.CallOption) (*CreateMapConfigReply, error)
	GetMapConfig(ctx context.Context, in *GetMapConfigRequest, opts ...grpc.CallOption) (*GetMapConfigReply, error)
	CreateSystemInfo(ctx context.Context, in *CreateSystemInfoRequest, opts ...grpc.CallOption) (*CreateSystemInfoReply, error)
	GetSystemInfo(ctx context.Context, in *GetSystemInfoRequest, opts ...grpc.CallOption) (*GetSystemInfoReply, error)
	DeleteSystemInfo(ctx context.Context, in *DeleteSystemInfoRequest, opts ...grpc.CallOption) (*DeleteSystemInfoReply, error)
	CreateModule(ctx context.Context, in *CreateModuleRequest, opts ...grpc.CallOption) (*CreateModuleReply, error)
	DeleteModule(ctx context.Context, in *DeleteModuleRequest, opts ...grpc.CallOption) (*DeleteModuleReply, error)
	UpdateModule(ctx context.Context, in *UpdateModuleRequest, opts ...grpc.CallOption) (*UpdateModuleReply, error)
	GetModule(ctx context.Context, in *GetModuleRequest, opts ...grpc.CallOption) (*GetModuleReply, error)
	ListModules(ctx context.Context, in *ListModulesRequest, opts ...grpc.CallOption) (*ListModulesReply, error)
	CreateCameraAttr(ctx context.Context, in *CreateCameraAttrRequest, opts ...grpc.CallOption) (*CreateCameraAttrReply, error)
	DeleteCameraAttr(ctx context.Context, in *DeleteCameraAttrRequest, opts ...grpc.CallOption) (*DeleteCameraAttrReply, error)
	GetCameraAttr(ctx context.Context, in *GetCameraAttrRequest, opts ...grpc.CallOption) (*GetCameraAttrReply, error)
	ListCameraAttr(ctx context.Context, in *ListCameraAttrRequest, opts ...grpc.CallOption) (*ListCameraAttrReply, error)
	CreateGeneralParameters(ctx context.Context, in *CreateGeneralParametersRequest, opts ...grpc.CallOption) (*CreateGeneralParametersReply, error)
	UpdateGeneralParameters(ctx context.Context, in *UpdateGeneralParametersRequest, opts ...grpc.CallOption) (*UpdateGeneralParametersReply, error)
	DeleteGeneralParameters(ctx context.Context, in *DeleteGeneralParametersRequest, opts ...grpc.CallOption) (*DeleteGeneralParametersReply, error)
	ListGeneralParameters(ctx context.Context, in *ListGeneralParametersRequest, opts ...grpc.CallOption) (*ListGeneralParametersReply, error)
}

type settingsClient struct {
	cc grpc.ClientConnInterface
}

func NewSettingsClient(cc grpc.ClientConnInterface) SettingsClient {
	return &settingsClient{cc}
}

func (c *settingsClient) CreateMapConfig(ctx context.Context, in *CreateMapConfigRequest, opts ...grpc.CallOption) (*CreateMapConfigReply, error) {
	out := new(CreateMapConfigReply)
	err := c.cc.Invoke(ctx, "/api.settings.v1.Settings/CreateMapConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) GetMapConfig(ctx context.Context, in *GetMapConfigRequest, opts ...grpc.CallOption) (*GetMapConfigReply, error) {
	out := new(GetMapConfigReply)
	err := c.cc.Invoke(ctx, "/api.settings.v1.Settings/GetMapConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) CreateSystemInfo(ctx context.Context, in *CreateSystemInfoRequest, opts ...grpc.CallOption) (*CreateSystemInfoReply, error) {
	out := new(CreateSystemInfoReply)
	err := c.cc.Invoke(ctx, "/api.settings.v1.Settings/CreateSystemInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) GetSystemInfo(ctx context.Context, in *GetSystemInfoRequest, opts ...grpc.CallOption) (*GetSystemInfoReply, error) {
	out := new(GetSystemInfoReply)
	err := c.cc.Invoke(ctx, "/api.settings.v1.Settings/GetSystemInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) DeleteSystemInfo(ctx context.Context, in *DeleteSystemInfoRequest, opts ...grpc.CallOption) (*DeleteSystemInfoReply, error) {
	out := new(DeleteSystemInfoReply)
	err := c.cc.Invoke(ctx, "/api.settings.v1.Settings/DeleteSystemInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) CreateModule(ctx context.Context, in *CreateModuleRequest, opts ...grpc.CallOption) (*CreateModuleReply, error) {
	out := new(CreateModuleReply)
	err := c.cc.Invoke(ctx, "/api.settings.v1.Settings/CreateModule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) DeleteModule(ctx context.Context, in *DeleteModuleRequest, opts ...grpc.CallOption) (*DeleteModuleReply, error) {
	out := new(DeleteModuleReply)
	err := c.cc.Invoke(ctx, "/api.settings.v1.Settings/DeleteModule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) UpdateModule(ctx context.Context, in *UpdateModuleRequest, opts ...grpc.CallOption) (*UpdateModuleReply, error) {
	out := new(UpdateModuleReply)
	err := c.cc.Invoke(ctx, "/api.settings.v1.Settings/UpdateModule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) GetModule(ctx context.Context, in *GetModuleRequest, opts ...grpc.CallOption) (*GetModuleReply, error) {
	out := new(GetModuleReply)
	err := c.cc.Invoke(ctx, "/api.settings.v1.Settings/GetModule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) ListModules(ctx context.Context, in *ListModulesRequest, opts ...grpc.CallOption) (*ListModulesReply, error) {
	out := new(ListModulesReply)
	err := c.cc.Invoke(ctx, "/api.settings.v1.Settings/ListModules", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) CreateCameraAttr(ctx context.Context, in *CreateCameraAttrRequest, opts ...grpc.CallOption) (*CreateCameraAttrReply, error) {
	out := new(CreateCameraAttrReply)
	err := c.cc.Invoke(ctx, "/api.settings.v1.Settings/CreateCameraAttr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) DeleteCameraAttr(ctx context.Context, in *DeleteCameraAttrRequest, opts ...grpc.CallOption) (*DeleteCameraAttrReply, error) {
	out := new(DeleteCameraAttrReply)
	err := c.cc.Invoke(ctx, "/api.settings.v1.Settings/DeleteCameraAttr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) GetCameraAttr(ctx context.Context, in *GetCameraAttrRequest, opts ...grpc.CallOption) (*GetCameraAttrReply, error) {
	out := new(GetCameraAttrReply)
	err := c.cc.Invoke(ctx, "/api.settings.v1.Settings/GetCameraAttr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) ListCameraAttr(ctx context.Context, in *ListCameraAttrRequest, opts ...grpc.CallOption) (*ListCameraAttrReply, error) {
	out := new(ListCameraAttrReply)
	err := c.cc.Invoke(ctx, "/api.settings.v1.Settings/ListCameraAttr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) CreateGeneralParameters(ctx context.Context, in *CreateGeneralParametersRequest, opts ...grpc.CallOption) (*CreateGeneralParametersReply, error) {
	out := new(CreateGeneralParametersReply)
	err := c.cc.Invoke(ctx, "/api.settings.v1.Settings/CreateGeneralParameters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) UpdateGeneralParameters(ctx context.Context, in *UpdateGeneralParametersRequest, opts ...grpc.CallOption) (*UpdateGeneralParametersReply, error) {
	out := new(UpdateGeneralParametersReply)
	err := c.cc.Invoke(ctx, "/api.settings.v1.Settings/UpdateGeneralParameters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) DeleteGeneralParameters(ctx context.Context, in *DeleteGeneralParametersRequest, opts ...grpc.CallOption) (*DeleteGeneralParametersReply, error) {
	out := new(DeleteGeneralParametersReply)
	err := c.cc.Invoke(ctx, "/api.settings.v1.Settings/DeleteGeneralParameters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) ListGeneralParameters(ctx context.Context, in *ListGeneralParametersRequest, opts ...grpc.CallOption) (*ListGeneralParametersReply, error) {
	out := new(ListGeneralParametersReply)
	err := c.cc.Invoke(ctx, "/api.settings.v1.Settings/ListGeneralParameters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SettingsServer is the server API for Settings service.
// All implementations must embed UnimplementedSettingsServer
// for forward compatibility
type SettingsServer interface {
	CreateMapConfig(context.Context, *CreateMapConfigRequest) (*CreateMapConfigReply, error)
	GetMapConfig(context.Context, *GetMapConfigRequest) (*GetMapConfigReply, error)
	CreateSystemInfo(context.Context, *CreateSystemInfoRequest) (*CreateSystemInfoReply, error)
	GetSystemInfo(context.Context, *GetSystemInfoRequest) (*GetSystemInfoReply, error)
	DeleteSystemInfo(context.Context, *DeleteSystemInfoRequest) (*DeleteSystemInfoReply, error)
	CreateModule(context.Context, *CreateModuleRequest) (*CreateModuleReply, error)
	DeleteModule(context.Context, *DeleteModuleRequest) (*DeleteModuleReply, error)
	UpdateModule(context.Context, *UpdateModuleRequest) (*UpdateModuleReply, error)
	GetModule(context.Context, *GetModuleRequest) (*GetModuleReply, error)
	ListModules(context.Context, *ListModulesRequest) (*ListModulesReply, error)
	CreateCameraAttr(context.Context, *CreateCameraAttrRequest) (*CreateCameraAttrReply, error)
	DeleteCameraAttr(context.Context, *DeleteCameraAttrRequest) (*DeleteCameraAttrReply, error)
	GetCameraAttr(context.Context, *GetCameraAttrRequest) (*GetCameraAttrReply, error)
	ListCameraAttr(context.Context, *ListCameraAttrRequest) (*ListCameraAttrReply, error)
	CreateGeneralParameters(context.Context, *CreateGeneralParametersRequest) (*CreateGeneralParametersReply, error)
	UpdateGeneralParameters(context.Context, *UpdateGeneralParametersRequest) (*UpdateGeneralParametersReply, error)
	DeleteGeneralParameters(context.Context, *DeleteGeneralParametersRequest) (*DeleteGeneralParametersReply, error)
	ListGeneralParameters(context.Context, *ListGeneralParametersRequest) (*ListGeneralParametersReply, error)
	mustEmbedUnimplementedSettingsServer()
}

// UnimplementedSettingsServer must be embedded to have forward compatible implementations.
type UnimplementedSettingsServer struct {
}

func (UnimplementedSettingsServer) CreateMapConfig(context.Context, *CreateMapConfigRequest) (*CreateMapConfigReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMapConfig not implemented")
}
func (UnimplementedSettingsServer) GetMapConfig(context.Context, *GetMapConfigRequest) (*GetMapConfigReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMapConfig not implemented")
}
func (UnimplementedSettingsServer) CreateSystemInfo(context.Context, *CreateSystemInfoRequest) (*CreateSystemInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSystemInfo not implemented")
}
func (UnimplementedSettingsServer) GetSystemInfo(context.Context, *GetSystemInfoRequest) (*GetSystemInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSystemInfo not implemented")
}
func (UnimplementedSettingsServer) DeleteSystemInfo(context.Context, *DeleteSystemInfoRequest) (*DeleteSystemInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSystemInfo not implemented")
}
func (UnimplementedSettingsServer) CreateModule(context.Context, *CreateModuleRequest) (*CreateModuleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateModule not implemented")
}
func (UnimplementedSettingsServer) DeleteModule(context.Context, *DeleteModuleRequest) (*DeleteModuleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteModule not implemented")
}
func (UnimplementedSettingsServer) UpdateModule(context.Context, *UpdateModuleRequest) (*UpdateModuleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateModule not implemented")
}
func (UnimplementedSettingsServer) GetModule(context.Context, *GetModuleRequest) (*GetModuleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModule not implemented")
}
func (UnimplementedSettingsServer) ListModules(context.Context, *ListModulesRequest) (*ListModulesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListModules not implemented")
}
func (UnimplementedSettingsServer) CreateCameraAttr(context.Context, *CreateCameraAttrRequest) (*CreateCameraAttrReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCameraAttr not implemented")
}
func (UnimplementedSettingsServer) DeleteCameraAttr(context.Context, *DeleteCameraAttrRequest) (*DeleteCameraAttrReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCameraAttr not implemented")
}
func (UnimplementedSettingsServer) GetCameraAttr(context.Context, *GetCameraAttrRequest) (*GetCameraAttrReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCameraAttr not implemented")
}
func (UnimplementedSettingsServer) ListCameraAttr(context.Context, *ListCameraAttrRequest) (*ListCameraAttrReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCameraAttr not implemented")
}
func (UnimplementedSettingsServer) CreateGeneralParameters(context.Context, *CreateGeneralParametersRequest) (*CreateGeneralParametersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGeneralParameters not implemented")
}
func (UnimplementedSettingsServer) UpdateGeneralParameters(context.Context, *UpdateGeneralParametersRequest) (*UpdateGeneralParametersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGeneralParameters not implemented")
}
func (UnimplementedSettingsServer) DeleteGeneralParameters(context.Context, *DeleteGeneralParametersRequest) (*DeleteGeneralParametersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGeneralParameters not implemented")
}
func (UnimplementedSettingsServer) ListGeneralParameters(context.Context, *ListGeneralParametersRequest) (*ListGeneralParametersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGeneralParameters not implemented")
}
func (UnimplementedSettingsServer) mustEmbedUnimplementedSettingsServer() {}

// UnsafeSettingsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SettingsServer will
// result in compilation errors.
type UnsafeSettingsServer interface {
	mustEmbedUnimplementedSettingsServer()
}

func RegisterSettingsServer(s grpc.ServiceRegistrar, srv SettingsServer) {
	s.RegisterService(&Settings_ServiceDesc, srv)
}

func _Settings_CreateMapConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMapConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).CreateMapConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.settings.v1.Settings/CreateMapConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).CreateMapConfig(ctx, req.(*CreateMapConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_GetMapConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMapConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).GetMapConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.settings.v1.Settings/GetMapConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).GetMapConfig(ctx, req.(*GetMapConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_CreateSystemInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSystemInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).CreateSystemInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.settings.v1.Settings/CreateSystemInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).CreateSystemInfo(ctx, req.(*CreateSystemInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_GetSystemInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSystemInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).GetSystemInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.settings.v1.Settings/GetSystemInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).GetSystemInfo(ctx, req.(*GetSystemInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_DeleteSystemInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSystemInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).DeleteSystemInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.settings.v1.Settings/DeleteSystemInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).DeleteSystemInfo(ctx, req.(*DeleteSystemInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_CreateModule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateModuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).CreateModule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.settings.v1.Settings/CreateModule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).CreateModule(ctx, req.(*CreateModuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_DeleteModule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteModuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).DeleteModule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.settings.v1.Settings/DeleteModule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).DeleteModule(ctx, req.(*DeleteModuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_UpdateModule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateModuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).UpdateModule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.settings.v1.Settings/UpdateModule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).UpdateModule(ctx, req.(*UpdateModuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_GetModule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).GetModule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.settings.v1.Settings/GetModule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).GetModule(ctx, req.(*GetModuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_ListModules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListModulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).ListModules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.settings.v1.Settings/ListModules",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).ListModules(ctx, req.(*ListModulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_CreateCameraAttr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCameraAttrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).CreateCameraAttr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.settings.v1.Settings/CreateCameraAttr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).CreateCameraAttr(ctx, req.(*CreateCameraAttrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_DeleteCameraAttr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCameraAttrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).DeleteCameraAttr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.settings.v1.Settings/DeleteCameraAttr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).DeleteCameraAttr(ctx, req.(*DeleteCameraAttrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_GetCameraAttr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCameraAttrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).GetCameraAttr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.settings.v1.Settings/GetCameraAttr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).GetCameraAttr(ctx, req.(*GetCameraAttrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_ListCameraAttr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCameraAttrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).ListCameraAttr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.settings.v1.Settings/ListCameraAttr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).ListCameraAttr(ctx, req.(*ListCameraAttrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_CreateGeneralParameters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGeneralParametersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).CreateGeneralParameters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.settings.v1.Settings/CreateGeneralParameters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).CreateGeneralParameters(ctx, req.(*CreateGeneralParametersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_UpdateGeneralParameters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGeneralParametersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).UpdateGeneralParameters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.settings.v1.Settings/UpdateGeneralParameters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).UpdateGeneralParameters(ctx, req.(*UpdateGeneralParametersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_DeleteGeneralParameters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGeneralParametersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).DeleteGeneralParameters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.settings.v1.Settings/DeleteGeneralParameters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).DeleteGeneralParameters(ctx, req.(*DeleteGeneralParametersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_ListGeneralParameters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGeneralParametersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).ListGeneralParameters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.settings.v1.Settings/ListGeneralParameters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).ListGeneralParameters(ctx, req.(*ListGeneralParametersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Settings_ServiceDesc is the grpc.ServiceDesc for Settings service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Settings_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.settings.v1.Settings",
	HandlerType: (*SettingsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMapConfig",
			Handler:    _Settings_CreateMapConfig_Handler,
		},
		{
			MethodName: "GetMapConfig",
			Handler:    _Settings_GetMapConfig_Handler,
		},
		{
			MethodName: "CreateSystemInfo",
			Handler:    _Settings_CreateSystemInfo_Handler,
		},
		{
			MethodName: "GetSystemInfo",
			Handler:    _Settings_GetSystemInfo_Handler,
		},
		{
			MethodName: "DeleteSystemInfo",
			Handler:    _Settings_DeleteSystemInfo_Handler,
		},
		{
			MethodName: "CreateModule",
			Handler:    _Settings_CreateModule_Handler,
		},
		{
			MethodName: "DeleteModule",
			Handler:    _Settings_DeleteModule_Handler,
		},
		{
			MethodName: "UpdateModule",
			Handler:    _Settings_UpdateModule_Handler,
		},
		{
			MethodName: "GetModule",
			Handler:    _Settings_GetModule_Handler,
		},
		{
			MethodName: "ListModules",
			Handler:    _Settings_ListModules_Handler,
		},
		{
			MethodName: "CreateCameraAttr",
			Handler:    _Settings_CreateCameraAttr_Handler,
		},
		{
			MethodName: "DeleteCameraAttr",
			Handler:    _Settings_DeleteCameraAttr_Handler,
		},
		{
			MethodName: "GetCameraAttr",
			Handler:    _Settings_GetCameraAttr_Handler,
		},
		{
			MethodName: "ListCameraAttr",
			Handler:    _Settings_ListCameraAttr_Handler,
		},
		{
			MethodName: "CreateGeneralParameters",
			Handler:    _Settings_CreateGeneralParameters_Handler,
		},
		{
			MethodName: "UpdateGeneralParameters",
			Handler:    _Settings_UpdateGeneralParameters_Handler,
		},
		{
			MethodName: "DeleteGeneralParameters",
			Handler:    _Settings_DeleteGeneralParameters_Handler,
		},
		{
			MethodName: "ListGeneralParameters",
			Handler:    _Settings_ListGeneralParameters_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "settings/v1/settings.proto",
}
