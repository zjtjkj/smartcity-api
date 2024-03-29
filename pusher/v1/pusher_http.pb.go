// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.4.0
// - protoc             v3.20.1
// source: pusher/v1/pusher.proto

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

const OperationPusherCreateEndpoint = "/api.pusher.v1.Pusher/CreateEndpoint"
const OperationPusherCreatePusher = "/api.pusher.v1.Pusher/CreatePusher"
const OperationPusherDeleteEndpoint = "/api.pusher.v1.Pusher/DeleteEndpoint"
const OperationPusherDeletePusher = "/api.pusher.v1.Pusher/DeletePusher"
const OperationPusherGetEndpoint = "/api.pusher.v1.Pusher/GetEndpoint"
const OperationPusherGetPusher = "/api.pusher.v1.Pusher/GetPusher"
const OperationPusherListController = "/api.pusher.v1.Pusher/ListController"
const OperationPusherListEndpoint = "/api.pusher.v1.Pusher/ListEndpoint"
const OperationPusherListPusher = "/api.pusher.v1.Pusher/ListPusher"
const OperationPusherReport = "/api.pusher.v1.Pusher/Report"
const OperationPusherUpdateEndpoint = "/api.pusher.v1.Pusher/UpdateEndpoint"
const OperationPusherUpdatePusher = "/api.pusher.v1.Pusher/UpdatePusher"

type PusherHTTPServer interface {
	CreateEndpoint(context.Context, *CreateEndpointRequest) (*CreateEndpointReply, error)
	CreatePusher(context.Context, *CreatePusherRequest) (*CreatePusherReply, error)
	DeleteEndpoint(context.Context, *DeleteEndpointRequest) (*DeleteEndpointReply, error)
	DeletePusher(context.Context, *DeletePusherRequest) (*DeletePusherReply, error)
	GetEndpoint(context.Context, *GetEndpointRequest) (*GetEndpointReply, error)
	GetPusher(context.Context, *GetPusherRequest) (*GetPusherReply, error)
	ListController(context.Context, *ListControllerRequest) (*ListControllerReply, error)
	ListEndpoint(context.Context, *ListEndpointRequest) (*ListEndpointReply, error)
	ListPusher(context.Context, *ListPusherRequest) (*ListPusherReply, error)
	Report(context.Context, *ReportRequest) (*ReportReply, error)
	UpdateEndpoint(context.Context, *UpdateEndpointRequest) (*UpdateEndpointReply, error)
	UpdatePusher(context.Context, *UpdatePusherRequest) (*UpdatePusherReply, error)
}

func RegisterPusherHTTPServer(s *http.Server, srv PusherHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/pusher/report", _Pusher_Report0_HTTP_Handler(srv))
	r.GET("/api/v1/pushers", _Pusher_ListController0_HTTP_Handler(srv))
	r.PUT("/api/v1/pusher", _Pusher_CreatePusher0_HTTP_Handler(srv))
	r.POST("/api/v1/pusher/{id}", _Pusher_UpdatePusher0_HTTP_Handler(srv))
	r.DELETE("/api/v1/pusher", _Pusher_DeletePusher0_HTTP_Handler(srv))
	r.GET("/api/v1/pusher/{id}", _Pusher_GetPusher0_HTTP_Handler(srv))
	r.GET("/api/v1/pushers", _Pusher_ListPusher0_HTTP_Handler(srv))
	r.PUT("/api/v1/pusher/receiver", _Pusher_CreateEndpoint0_HTTP_Handler(srv))
	r.POST("/api/v1/push/receiver/{id}", _Pusher_UpdateEndpoint0_HTTP_Handler(srv))
	r.DELETE("/api/v1/push/receiver/{id}", _Pusher_DeleteEndpoint0_HTTP_Handler(srv))
	r.GET("/api/v1/push/receiver/{id}", _Pusher_GetEndpoint0_HTTP_Handler(srv))
	r.GET("/api/v1/push/receivers", _Pusher_ListEndpoint0_HTTP_Handler(srv))
}

func _Pusher_Report0_HTTP_Handler(srv PusherHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ReportRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPusherReport)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Report(ctx, req.(*ReportRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ReportReply)
		return ctx.Result(200, reply)
	}
}

func _Pusher_ListController0_HTTP_Handler(srv PusherHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListControllerRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPusherListController)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListController(ctx, req.(*ListControllerRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListControllerReply)
		return ctx.Result(200, reply)
	}
}

func _Pusher_CreatePusher0_HTTP_Handler(srv PusherHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreatePusherRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPusherCreatePusher)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreatePusher(ctx, req.(*CreatePusherRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreatePusherReply)
		return ctx.Result(200, reply)
	}
}

func _Pusher_UpdatePusher0_HTTP_Handler(srv PusherHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdatePusherRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPusherUpdatePusher)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdatePusher(ctx, req.(*UpdatePusherRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdatePusherReply)
		return ctx.Result(200, reply)
	}
}

func _Pusher_DeletePusher0_HTTP_Handler(srv PusherHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeletePusherRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPusherDeletePusher)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeletePusher(ctx, req.(*DeletePusherRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeletePusherReply)
		return ctx.Result(200, reply)
	}
}

func _Pusher_GetPusher0_HTTP_Handler(srv PusherHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetPusherRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPusherGetPusher)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPusher(ctx, req.(*GetPusherRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetPusherReply)
		return ctx.Result(200, reply)
	}
}

func _Pusher_ListPusher0_HTTP_Handler(srv PusherHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListPusherRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPusherListPusher)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListPusher(ctx, req.(*ListPusherRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListPusherReply)
		return ctx.Result(200, reply)
	}
}

func _Pusher_CreateEndpoint0_HTTP_Handler(srv PusherHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateEndpointRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPusherCreateEndpoint)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateEndpoint(ctx, req.(*CreateEndpointRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateEndpointReply)
		return ctx.Result(200, reply)
	}
}

func _Pusher_UpdateEndpoint0_HTTP_Handler(srv PusherHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateEndpointRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPusherUpdateEndpoint)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateEndpoint(ctx, req.(*UpdateEndpointRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateEndpointReply)
		return ctx.Result(200, reply)
	}
}

func _Pusher_DeleteEndpoint0_HTTP_Handler(srv PusherHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteEndpointRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPusherDeleteEndpoint)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteEndpoint(ctx, req.(*DeleteEndpointRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteEndpointReply)
		return ctx.Result(200, reply)
	}
}

func _Pusher_GetEndpoint0_HTTP_Handler(srv PusherHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetEndpointRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPusherGetEndpoint)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetEndpoint(ctx, req.(*GetEndpointRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetEndpointReply)
		return ctx.Result(200, reply)
	}
}

func _Pusher_ListEndpoint0_HTTP_Handler(srv PusherHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListEndpointRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPusherListEndpoint)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListEndpoint(ctx, req.(*ListEndpointRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListEndpointReply)
		return ctx.Result(200, reply)
	}
}

type PusherHTTPClient interface {
	CreateEndpoint(ctx context.Context, req *CreateEndpointRequest, opts ...http.CallOption) (rsp *CreateEndpointReply, err error)
	CreatePusher(ctx context.Context, req *CreatePusherRequest, opts ...http.CallOption) (rsp *CreatePusherReply, err error)
	DeleteEndpoint(ctx context.Context, req *DeleteEndpointRequest, opts ...http.CallOption) (rsp *DeleteEndpointReply, err error)
	DeletePusher(ctx context.Context, req *DeletePusherRequest, opts ...http.CallOption) (rsp *DeletePusherReply, err error)
	GetEndpoint(ctx context.Context, req *GetEndpointRequest, opts ...http.CallOption) (rsp *GetEndpointReply, err error)
	GetPusher(ctx context.Context, req *GetPusherRequest, opts ...http.CallOption) (rsp *GetPusherReply, err error)
	ListController(ctx context.Context, req *ListControllerRequest, opts ...http.CallOption) (rsp *ListControllerReply, err error)
	ListEndpoint(ctx context.Context, req *ListEndpointRequest, opts ...http.CallOption) (rsp *ListEndpointReply, err error)
	ListPusher(ctx context.Context, req *ListPusherRequest, opts ...http.CallOption) (rsp *ListPusherReply, err error)
	Report(ctx context.Context, req *ReportRequest, opts ...http.CallOption) (rsp *ReportReply, err error)
	UpdateEndpoint(ctx context.Context, req *UpdateEndpointRequest, opts ...http.CallOption) (rsp *UpdateEndpointReply, err error)
	UpdatePusher(ctx context.Context, req *UpdatePusherRequest, opts ...http.CallOption) (rsp *UpdatePusherReply, err error)
}

type PusherHTTPClientImpl struct {
	cc *http.Client
}

func NewPusherHTTPClient(client *http.Client) PusherHTTPClient {
	return &PusherHTTPClientImpl{client}
}

func (c *PusherHTTPClientImpl) CreateEndpoint(ctx context.Context, in *CreateEndpointRequest, opts ...http.CallOption) (*CreateEndpointReply, error) {
	var out CreateEndpointReply
	pattern := "/api/v1/pusher/receiver"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPusherCreateEndpoint))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PusherHTTPClientImpl) CreatePusher(ctx context.Context, in *CreatePusherRequest, opts ...http.CallOption) (*CreatePusherReply, error) {
	var out CreatePusherReply
	pattern := "/api/v1/pusher"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPusherCreatePusher))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PusherHTTPClientImpl) DeleteEndpoint(ctx context.Context, in *DeleteEndpointRequest, opts ...http.CallOption) (*DeleteEndpointReply, error) {
	var out DeleteEndpointReply
	pattern := "/api/v1/push/receiver/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPusherDeleteEndpoint))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PusherHTTPClientImpl) DeletePusher(ctx context.Context, in *DeletePusherRequest, opts ...http.CallOption) (*DeletePusherReply, error) {
	var out DeletePusherReply
	pattern := "/api/v1/pusher"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPusherDeletePusher))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PusherHTTPClientImpl) GetEndpoint(ctx context.Context, in *GetEndpointRequest, opts ...http.CallOption) (*GetEndpointReply, error) {
	var out GetEndpointReply
	pattern := "/api/v1/push/receiver/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPusherGetEndpoint))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PusherHTTPClientImpl) GetPusher(ctx context.Context, in *GetPusherRequest, opts ...http.CallOption) (*GetPusherReply, error) {
	var out GetPusherReply
	pattern := "/api/v1/pusher/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPusherGetPusher))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PusherHTTPClientImpl) ListController(ctx context.Context, in *ListControllerRequest, opts ...http.CallOption) (*ListControllerReply, error) {
	var out ListControllerReply
	pattern := "/api/v1/pushers"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPusherListController))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PusherHTTPClientImpl) ListEndpoint(ctx context.Context, in *ListEndpointRequest, opts ...http.CallOption) (*ListEndpointReply, error) {
	var out ListEndpointReply
	pattern := "/api/v1/push/receivers"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPusherListEndpoint))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PusherHTTPClientImpl) ListPusher(ctx context.Context, in *ListPusherRequest, opts ...http.CallOption) (*ListPusherReply, error) {
	var out ListPusherReply
	pattern := "/api/v1/pushers"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPusherListPusher))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PusherHTTPClientImpl) Report(ctx context.Context, in *ReportRequest, opts ...http.CallOption) (*ReportReply, error) {
	var out ReportReply
	pattern := "/api/v1/pusher/report"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPusherReport))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PusherHTTPClientImpl) UpdateEndpoint(ctx context.Context, in *UpdateEndpointRequest, opts ...http.CallOption) (*UpdateEndpointReply, error) {
	var out UpdateEndpointReply
	pattern := "/api/v1/push/receiver/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPusherUpdateEndpoint))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PusherHTTPClientImpl) UpdatePusher(ctx context.Context, in *UpdatePusherRequest, opts ...http.CallOption) (*UpdatePusherReply, error) {
	var out UpdatePusherReply
	pattern := "/api/v1/pusher/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPusherUpdatePusher))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
