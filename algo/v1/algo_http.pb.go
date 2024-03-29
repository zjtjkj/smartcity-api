// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.4.0
// - protoc             v3.20.1
// source: algo/v1/algo.proto

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

const OperationAlgoClear = "/api.algo.v1.Algo/Clear"
const OperationAlgoCreateMission = "/api.algo.v1.Algo/CreateMission"
const OperationAlgoDebug = "/api.algo.v1.Algo/Debug"
const OperationAlgoDeleteMission = "/api.algo.v1.Algo/DeleteMission"
const OperationAlgoDeleteMultiMission = "/api.algo.v1.Algo/DeleteMultiMission"
const OperationAlgoInfo = "/api.algo.v1.Algo/Info"
const OperationAlgoListAll = "/api.algo.v1.Algo/ListAll"
const OperationAlgoListMission = "/api.algo.v1.Algo/ListMission"

type AlgoHTTPServer interface {
	Clear(context.Context, *ClearRequest) (*ClearReply, error)
	CreateMission(context.Context, *CreateMissionRequest) (*CreateMissionReply, error)
	Debug(context.Context, *DebugRequest) (*DebugResponse, error)
	DeleteMission(context.Context, *DeleteMissionRequest) (*DeleteMissionReply, error)
	DeleteMultiMission(context.Context, *DeleteMissionMultiRequest) (*DeleteMissionMultiReply, error)
	Info(context.Context, *InfoRequest) (*InfoResponse, error)
	ListAll(context.Context, *ListAllRequest) (*ListAllReply, error)
	ListMission(context.Context, *ListMissionRequest) (*ListMissionReply, error)
}

func RegisterAlgoHTTPServer(s *http.Server, srv AlgoHTTPServer) {
	r := s.Route("/")
	r.PUT("/api/v1/algo/mission", _Algo_CreateMission0_HTTP_Handler(srv))
	r.DELETE("/api/v1/algo/mission/{id}", _Algo_DeleteMission0_HTTP_Handler(srv))
	r.DELETE("/api/v1/algo/missions/{id}", _Algo_DeleteMultiMission0_HTTP_Handler(srv))
	r.GET("/api/v1/algo/missions/{id}", _Algo_ListMission0_HTTP_Handler(srv))
	r.GET("/api/v1/algo/missions", _Algo_ListAll0_HTTP_Handler(srv))
	r.GET("/api/v1/algo/info", _Algo_Info0_HTTP_Handler(srv))
	r.GET("/api/v1/algo/debug", _Algo_Debug0_HTTP_Handler(srv))
	r.GET("/api/v1/algo/clear", _Algo_Clear0_HTTP_Handler(srv))
}

func _Algo_CreateMission0_HTTP_Handler(srv AlgoHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateMissionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAlgoCreateMission)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateMission(ctx, req.(*CreateMissionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateMissionReply)
		return ctx.Result(200, reply)
	}
}

func _Algo_DeleteMission0_HTTP_Handler(srv AlgoHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteMissionRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAlgoDeleteMission)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteMission(ctx, req.(*DeleteMissionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteMissionReply)
		return ctx.Result(200, reply)
	}
}

func _Algo_DeleteMultiMission0_HTTP_Handler(srv AlgoHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteMissionMultiRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAlgoDeleteMultiMission)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteMultiMission(ctx, req.(*DeleteMissionMultiRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteMissionMultiReply)
		return ctx.Result(200, reply)
	}
}

func _Algo_ListMission0_HTTP_Handler(srv AlgoHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListMissionRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAlgoListMission)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListMission(ctx, req.(*ListMissionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListMissionReply)
		return ctx.Result(200, reply)
	}
}

func _Algo_ListAll0_HTTP_Handler(srv AlgoHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListAllRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAlgoListAll)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListAll(ctx, req.(*ListAllRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListAllReply)
		return ctx.Result(200, reply)
	}
}

func _Algo_Info0_HTTP_Handler(srv AlgoHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in InfoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAlgoInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Info(ctx, req.(*InfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*InfoResponse)
		return ctx.Result(200, reply)
	}
}

func _Algo_Debug0_HTTP_Handler(srv AlgoHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DebugRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAlgoDebug)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Debug(ctx, req.(*DebugRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DebugResponse)
		return ctx.Result(200, reply)
	}
}

func _Algo_Clear0_HTTP_Handler(srv AlgoHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ClearRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAlgoClear)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Clear(ctx, req.(*ClearRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ClearReply)
		return ctx.Result(200, reply)
	}
}

type AlgoHTTPClient interface {
	Clear(ctx context.Context, req *ClearRequest, opts ...http.CallOption) (rsp *ClearReply, err error)
	CreateMission(ctx context.Context, req *CreateMissionRequest, opts ...http.CallOption) (rsp *CreateMissionReply, err error)
	Debug(ctx context.Context, req *DebugRequest, opts ...http.CallOption) (rsp *DebugResponse, err error)
	DeleteMission(ctx context.Context, req *DeleteMissionRequest, opts ...http.CallOption) (rsp *DeleteMissionReply, err error)
	DeleteMultiMission(ctx context.Context, req *DeleteMissionMultiRequest, opts ...http.CallOption) (rsp *DeleteMissionMultiReply, err error)
	Info(ctx context.Context, req *InfoRequest, opts ...http.CallOption) (rsp *InfoResponse, err error)
	ListAll(ctx context.Context, req *ListAllRequest, opts ...http.CallOption) (rsp *ListAllReply, err error)
	ListMission(ctx context.Context, req *ListMissionRequest, opts ...http.CallOption) (rsp *ListMissionReply, err error)
}

type AlgoHTTPClientImpl struct {
	cc *http.Client
}

func NewAlgoHTTPClient(client *http.Client) AlgoHTTPClient {
	return &AlgoHTTPClientImpl{client}
}

func (c *AlgoHTTPClientImpl) Clear(ctx context.Context, in *ClearRequest, opts ...http.CallOption) (*ClearReply, error) {
	var out ClearReply
	pattern := "/api/v1/algo/clear"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAlgoClear))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AlgoHTTPClientImpl) CreateMission(ctx context.Context, in *CreateMissionRequest, opts ...http.CallOption) (*CreateMissionReply, error) {
	var out CreateMissionReply
	pattern := "/api/v1/algo/mission"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAlgoCreateMission))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AlgoHTTPClientImpl) Debug(ctx context.Context, in *DebugRequest, opts ...http.CallOption) (*DebugResponse, error) {
	var out DebugResponse
	pattern := "/api/v1/algo/debug"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAlgoDebug))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AlgoHTTPClientImpl) DeleteMission(ctx context.Context, in *DeleteMissionRequest, opts ...http.CallOption) (*DeleteMissionReply, error) {
	var out DeleteMissionReply
	pattern := "/api/v1/algo/mission/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAlgoDeleteMission))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AlgoHTTPClientImpl) DeleteMultiMission(ctx context.Context, in *DeleteMissionMultiRequest, opts ...http.CallOption) (*DeleteMissionMultiReply, error) {
	var out DeleteMissionMultiReply
	pattern := "/api/v1/algo/missions/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAlgoDeleteMultiMission))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AlgoHTTPClientImpl) Info(ctx context.Context, in *InfoRequest, opts ...http.CallOption) (*InfoResponse, error) {
	var out InfoResponse
	pattern := "/api/v1/algo/info"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAlgoInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AlgoHTTPClientImpl) ListAll(ctx context.Context, in *ListAllRequest, opts ...http.CallOption) (*ListAllReply, error) {
	var out ListAllReply
	pattern := "/api/v1/algo/missions"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAlgoListAll))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AlgoHTTPClientImpl) ListMission(ctx context.Context, in *ListMissionRequest, opts ...http.CallOption) (*ListMissionReply, error) {
	var out ListMissionReply
	pattern := "/api/v1/algo/missions/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAlgoListMission))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
