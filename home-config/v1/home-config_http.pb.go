// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.4.0
// - protoc             v3.20.0
// source: api/home-config/v1/home-config.proto

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

const OperationHomeConfigGetHomeConfig = "/api.home.v1.HomeConfig/GetHomeConfig"
const OperationHomeConfigGetOptionById = "/api.home.v1.HomeConfig/GetOptionById"
const OperationHomeConfigGetOptions = "/api.home.v1.HomeConfig/GetOptions"
const OperationHomeConfigSaveHomeConfig = "/api.home.v1.HomeConfig/SaveHomeConfig"

type HomeConfigHTTPServer interface {
	GetHomeConfig(context.Context, *GetHomeConfigRequest) (*GetHomeConfigReply, error)
	GetOptionById(context.Context, *GetOptionByIdRequest) (*GetOptionByIdReply, error)
	GetOptions(context.Context, *GetOptionsRequest) (*GetOptionsReply, error)
	SaveHomeConfig(context.Context, *SaveHomeConfigRequest) (*SaveHomeConfigReply, error)
}

func RegisterHomeConfigHTTPServer(s *http.Server, srv HomeConfigHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/home/config/save", _HomeConfig_SaveHomeConfig0_HTTP_Handler(srv))
	r.POST("/api/v1/home/config/get", _HomeConfig_GetHomeConfig0_HTTP_Handler(srv))
	r.POST("/api/v1/home/option/get", _HomeConfig_GetOptions0_HTTP_Handler(srv))
	r.POST("/api/v1/home/option/get", _HomeConfig_GetOptionById0_HTTP_Handler(srv))
}

func _HomeConfig_SaveHomeConfig0_HTTP_Handler(srv HomeConfigHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SaveHomeConfigRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationHomeConfigSaveHomeConfig)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SaveHomeConfig(ctx, req.(*SaveHomeConfigRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SaveHomeConfigReply)
		return ctx.Result(200, reply)
	}
}

func _HomeConfig_GetHomeConfig0_HTTP_Handler(srv HomeConfigHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetHomeConfigRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationHomeConfigGetHomeConfig)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetHomeConfig(ctx, req.(*GetHomeConfigRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetHomeConfigReply)
		return ctx.Result(200, reply)
	}
}

func _HomeConfig_GetOptions0_HTTP_Handler(srv HomeConfigHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetOptionsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationHomeConfigGetOptions)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetOptions(ctx, req.(*GetOptionsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetOptionsReply)
		return ctx.Result(200, reply)
	}
}

func _HomeConfig_GetOptionById0_HTTP_Handler(srv HomeConfigHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetOptionByIdRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationHomeConfigGetOptionById)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetOptionById(ctx, req.(*GetOptionByIdRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetOptionByIdReply)
		return ctx.Result(200, reply)
	}
}

type HomeConfigHTTPClient interface {
	GetHomeConfig(ctx context.Context, req *GetHomeConfigRequest, opts ...http.CallOption) (rsp *GetHomeConfigReply, err error)
	GetOptionById(ctx context.Context, req *GetOptionByIdRequest, opts ...http.CallOption) (rsp *GetOptionByIdReply, err error)
	GetOptions(ctx context.Context, req *GetOptionsRequest, opts ...http.CallOption) (rsp *GetOptionsReply, err error)
	SaveHomeConfig(ctx context.Context, req *SaveHomeConfigRequest, opts ...http.CallOption) (rsp *SaveHomeConfigReply, err error)
}

type HomeConfigHTTPClientImpl struct {
	cc *http.Client
}

func NewHomeConfigHTTPClient(client *http.Client) HomeConfigHTTPClient {
	return &HomeConfigHTTPClientImpl{client}
}

func (c *HomeConfigHTTPClientImpl) GetHomeConfig(ctx context.Context, in *GetHomeConfigRequest, opts ...http.CallOption) (*GetHomeConfigReply, error) {
	var out GetHomeConfigReply
	pattern := "/api/v1/home/config/get"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationHomeConfigGetHomeConfig))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *HomeConfigHTTPClientImpl) GetOptionById(ctx context.Context, in *GetOptionByIdRequest, opts ...http.CallOption) (*GetOptionByIdReply, error) {
	var out GetOptionByIdReply
	pattern := "/api/v1/home/option/get"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationHomeConfigGetOptionById))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *HomeConfigHTTPClientImpl) GetOptions(ctx context.Context, in *GetOptionsRequest, opts ...http.CallOption) (*GetOptionsReply, error) {
	var out GetOptionsReply
	pattern := "/api/v1/home/option/get"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationHomeConfigGetOptions))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *HomeConfigHTTPClientImpl) SaveHomeConfig(ctx context.Context, in *SaveHomeConfigRequest, opts ...http.CallOption) (*SaveHomeConfigReply, error) {
	var out SaveHomeConfigReply
	pattern := "/api/v1/home/config/save"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationHomeConfigSaveHomeConfig))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}