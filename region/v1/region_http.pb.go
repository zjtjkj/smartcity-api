// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.4.0
// - protoc             v3.20.0
// source: api/region/v1/region.proto

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

const OperationRegionServiceCreateRegion = "/api.region.v1.RegionService/CreateRegion"
const OperationRegionServiceDeleteRegion = "/api.region.v1.RegionService/DeleteRegion"
const OperationRegionServiceGetRegion = "/api.region.v1.RegionService/GetRegion"
const OperationRegionServiceListRegion = "/api.region.v1.RegionService/ListRegion"
const OperationRegionServiceListRegionById = "/api.region.v1.RegionService/ListRegionById"
const OperationRegionServiceUpdateRegion = "/api.region.v1.RegionService/UpdateRegion"

type RegionServiceHTTPServer interface {
	CreateRegion(context.Context, *CreateRegionRequest) (*CreateRegionReply, error)
	DeleteRegion(context.Context, *DeleteRegionRequest) (*DeleteRegionReply, error)
	GetRegion(context.Context, *GetRegionRequest) (*GetRegionReply, error)
	ListRegion(context.Context, *ListRegionRequest) (*ListRegionReply, error)
	ListRegionById(context.Context, *ListRegionByIdRequest) (*ListRegionByIdReply, error)
	UpdateRegion(context.Context, *UpdateRegionRequest) (*UpdateRegionReply, error)
}

func RegisterRegionServiceHTTPServer(s *http.Server, srv RegionServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/region/create", _RegionService_CreateRegion0_HTTP_Handler(srv))
	r.POST("/api/v1/region/update", _RegionService_UpdateRegion0_HTTP_Handler(srv))
	r.POST("/api/v1/region/delete", _RegionService_DeleteRegion0_HTTP_Handler(srv))
	r.POST("/api/v1/region/get", _RegionService_GetRegion0_HTTP_Handler(srv))
	r.GET("/api/v1/region/list", _RegionService_ListRegion0_HTTP_Handler(srv))
	r.GET("/api/v1/region/list_id", _RegionService_ListRegionById0_HTTP_Handler(srv))
}

func _RegionService_CreateRegion0_HTTP_Handler(srv RegionServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateRegionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRegionServiceCreateRegion)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateRegion(ctx, req.(*CreateRegionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateRegionReply)
		return ctx.Result(200, reply)
	}
}

func _RegionService_UpdateRegion0_HTTP_Handler(srv RegionServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateRegionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRegionServiceUpdateRegion)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateRegion(ctx, req.(*UpdateRegionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateRegionReply)
		return ctx.Result(200, reply)
	}
}

func _RegionService_DeleteRegion0_HTTP_Handler(srv RegionServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteRegionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRegionServiceDeleteRegion)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteRegion(ctx, req.(*DeleteRegionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteRegionReply)
		return ctx.Result(200, reply)
	}
}

func _RegionService_GetRegion0_HTTP_Handler(srv RegionServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetRegionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRegionServiceGetRegion)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetRegion(ctx, req.(*GetRegionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetRegionReply)
		return ctx.Result(200, reply)
	}
}

func _RegionService_ListRegion0_HTTP_Handler(srv RegionServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListRegionRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRegionServiceListRegion)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListRegion(ctx, req.(*ListRegionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListRegionReply)
		return ctx.Result(200, reply)
	}
}

func _RegionService_ListRegionById0_HTTP_Handler(srv RegionServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListRegionByIdRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRegionServiceListRegionById)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListRegionById(ctx, req.(*ListRegionByIdRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListRegionByIdReply)
		return ctx.Result(200, reply)
	}
}

type RegionServiceHTTPClient interface {
	CreateRegion(ctx context.Context, req *CreateRegionRequest, opts ...http.CallOption) (rsp *CreateRegionReply, err error)
	DeleteRegion(ctx context.Context, req *DeleteRegionRequest, opts ...http.CallOption) (rsp *DeleteRegionReply, err error)
	GetRegion(ctx context.Context, req *GetRegionRequest, opts ...http.CallOption) (rsp *GetRegionReply, err error)
	ListRegion(ctx context.Context, req *ListRegionRequest, opts ...http.CallOption) (rsp *ListRegionReply, err error)
	ListRegionById(ctx context.Context, req *ListRegionByIdRequest, opts ...http.CallOption) (rsp *ListRegionByIdReply, err error)
	UpdateRegion(ctx context.Context, req *UpdateRegionRequest, opts ...http.CallOption) (rsp *UpdateRegionReply, err error)
}

type RegionServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewRegionServiceHTTPClient(client *http.Client) RegionServiceHTTPClient {
	return &RegionServiceHTTPClientImpl{client}
}

func (c *RegionServiceHTTPClientImpl) CreateRegion(ctx context.Context, in *CreateRegionRequest, opts ...http.CallOption) (*CreateRegionReply, error) {
	var out CreateRegionReply
	pattern := "/api/v1/region/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRegionServiceCreateRegion))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RegionServiceHTTPClientImpl) DeleteRegion(ctx context.Context, in *DeleteRegionRequest, opts ...http.CallOption) (*DeleteRegionReply, error) {
	var out DeleteRegionReply
	pattern := "/api/v1/region/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRegionServiceDeleteRegion))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RegionServiceHTTPClientImpl) GetRegion(ctx context.Context, in *GetRegionRequest, opts ...http.CallOption) (*GetRegionReply, error) {
	var out GetRegionReply
	pattern := "/api/v1/region/get"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRegionServiceGetRegion))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RegionServiceHTTPClientImpl) ListRegion(ctx context.Context, in *ListRegionRequest, opts ...http.CallOption) (*ListRegionReply, error) {
	var out ListRegionReply
	pattern := "/api/v1/region/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRegionServiceListRegion))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RegionServiceHTTPClientImpl) ListRegionById(ctx context.Context, in *ListRegionByIdRequest, opts ...http.CallOption) (*ListRegionByIdReply, error) {
	var out ListRegionByIdReply
	pattern := "/api/v1/region/list_id"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRegionServiceListRegionById))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RegionServiceHTTPClientImpl) UpdateRegion(ctx context.Context, in *UpdateRegionRequest, opts ...http.CallOption) (*UpdateRegionReply, error) {
	var out UpdateRegionReply
	pattern := "/api/v1/region/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRegionServiceUpdateRegion))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
