// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.4.0
// - protoc             v3.20.1
// source: camera/v1/camera.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationCameraCreateCamera = "/api.camera.v1.Camera/CreateCamera"
const OperationCameraDeleteCamera = "/api.camera.v1.Camera/DeleteCamera"
const OperationCameraGetCamera = "/api.camera.v1.Camera/GetCamera"
const OperationCameraListCameraByKey = "/api.camera.v1.Camera/ListCameraByKey"
const OperationCameraListCameraByRegion = "/api.camera.v1.Camera/ListCameraByRegion"
const OperationCameraUpdateCamera = "/api.camera.v1.Camera/UpdateCamera"

type CameraHTTPServer interface {
	CreateCamera(context.Context, *CreateCameraRequest) (*CreateCameraReply, error)
	DeleteCamera(context.Context, *DeleteCameraRequest) (*DeleteCameraReply, error)
	GetCamera(context.Context, *GetCameraRequest) (*GetCameraReply, error)
	ListCameraByKey(context.Context, *ListCameraByKeyRequest) (*ListCameraByKeyReply, error)
	ListCameraByRegion(context.Context, *ListCameraByRegionRequest) (*ListCameraByRegionReply, error)
	UpdateCamera(context.Context, *UpdateCameraRequest) (*UpdateCameraReply, error)
}

func RegisterCameraHTTPServer(s *http.Server, srv CameraHTTPServer) {
	r := s.Route("/")
	r.PUT("/api/v1/camera", _Camera_CreateCamera0_HTTP_Handler(srv))
	r.POST("/api/v1/camera/{id}", _Camera_UpdateCamera0_HTTP_Handler(srv))
	r.DELETE("/api/v1/camera/{id}", _Camera_DeleteCamera0_HTTP_Handler(srv))
	r.GET("/api/v1/camera/{id}", _Camera_GetCamera0_HTTP_Handler(srv))
	r.GET("/api/v1/cameras/{id}", _Camera_ListCameraByRegion0_HTTP_Handler(srv))
	r.GET("/api/v1/cameras/key/{key}", _Camera_ListCameraByKey0_HTTP_Handler(srv))
}

func _Camera_CreateCamera0_HTTP_Handler(srv CameraHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateCameraRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCameraCreateCamera)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateCamera(ctx, req.(*CreateCameraRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateCameraReply)
		return ctx.Result(200, reply)
	}
}

func _Camera_UpdateCamera0_HTTP_Handler(srv CameraHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateCameraRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCameraUpdateCamera)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateCamera(ctx, req.(*UpdateCameraRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateCameraReply)
		return ctx.Result(200, reply)
	}
}

func _Camera_DeleteCamera0_HTTP_Handler(srv CameraHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteCameraRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCameraDeleteCamera)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteCamera(ctx, req.(*DeleteCameraRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteCameraReply)
		return ctx.Result(200, reply)
	}
}

func _Camera_GetCamera0_HTTP_Handler(srv CameraHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCameraRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCameraGetCamera)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCamera(ctx, req.(*GetCameraRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetCameraReply)
		return ctx.Result(200, reply)
	}
}

func _Camera_ListCameraByRegion0_HTTP_Handler(srv CameraHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListCameraByRegionRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCameraListCameraByRegion)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListCameraByRegion(ctx, req.(*ListCameraByRegionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListCameraByRegionReply)
		return ctx.Result(200, reply)
	}
}

func _Camera_ListCameraByKey0_HTTP_Handler(srv CameraHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListCameraByKeyRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCameraListCameraByKey)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListCameraByKey(ctx, req.(*ListCameraByKeyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListCameraByKeyReply)
		return ctx.Result(200, reply)
	}
}

type CameraHTTPClient interface {
	CreateCamera(ctx context.Context, req *CreateCameraRequest, opts ...http.CallOption) (rsp *CreateCameraReply, err error)
	DeleteCamera(ctx context.Context, req *DeleteCameraRequest, opts ...http.CallOption) (rsp *DeleteCameraReply, err error)
	GetCamera(ctx context.Context, req *GetCameraRequest, opts ...http.CallOption) (rsp *GetCameraReply, err error)
	ListCameraByKey(ctx context.Context, req *ListCameraByKeyRequest, opts ...http.CallOption) (rsp *ListCameraByKeyReply, err error)
	ListCameraByRegion(ctx context.Context, req *ListCameraByRegionRequest, opts ...http.CallOption) (rsp *ListCameraByRegionReply, err error)
	UpdateCamera(ctx context.Context, req *UpdateCameraRequest, opts ...http.CallOption) (rsp *UpdateCameraReply, err error)
}

type CameraHTTPClientImpl struct {
	cc *http.Client
}

func NewCameraHTTPClient(client *http.Client) CameraHTTPClient {
	return &CameraHTTPClientImpl{client}
}

func (c *CameraHTTPClientImpl) CreateCamera(ctx context.Context, in *CreateCameraRequest, opts ...http.CallOption) (*CreateCameraReply, error) {
	var out CreateCameraReply
	pattern := "/api/v1/camera"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCameraCreateCamera))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CameraHTTPClientImpl) DeleteCamera(ctx context.Context, in *DeleteCameraRequest, opts ...http.CallOption) (*DeleteCameraReply, error) {
	var out DeleteCameraReply
	pattern := "/api/v1/camera/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCameraDeleteCamera))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CameraHTTPClientImpl) GetCamera(ctx context.Context, in *GetCameraRequest, opts ...http.CallOption) (*GetCameraReply, error) {
	var out GetCameraReply
	pattern := "/api/v1/camera/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCameraGetCamera))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CameraHTTPClientImpl) ListCameraByKey(ctx context.Context, in *ListCameraByKeyRequest, opts ...http.CallOption) (*ListCameraByKeyReply, error) {
	var out ListCameraByKeyReply
	pattern := "/api/v1/cameras/key/{key}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCameraListCameraByKey))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CameraHTTPClientImpl) ListCameraByRegion(ctx context.Context, in *ListCameraByRegionRequest, opts ...http.CallOption) (*ListCameraByRegionReply, error) {
	var out ListCameraByRegionReply
	pattern := "/api/v1/cameras/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCameraListCameraByRegion))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CameraHTTPClientImpl) UpdateCamera(ctx context.Context, in *UpdateCameraRequest, opts ...http.CallOption) (*UpdateCameraReply, error) {
	var out UpdateCameraReply
	pattern := "/api/v1/camera/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCameraUpdateCamera))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
