// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.3
// - protoc             v3.14.0
// source: api/events/v1/events.proto

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

const OperationEventsCreateEvent = "/api.events.v1.Events/CreateEvent"

type EventsHTTPServer interface {
	CreateEvent(context.Context, *CreateEventRequest) (*CreateEventReply, error)
}

func RegisterEventsHTTPServer(s *http.Server, srv EventsHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/events", _Events_CreateEvent0_HTTP_Handler(srv))
}

func _Events_CreateEvent0_HTTP_Handler(srv EventsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateEventRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationEventsCreateEvent)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateEvent(ctx, req.(*CreateEventRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateEventReply)
		return ctx.Result(200, reply)
	}
}

type EventsHTTPClient interface {
	CreateEvent(ctx context.Context, req *CreateEventRequest, opts ...http.CallOption) (rsp *CreateEventReply, err error)
}

type EventsHTTPClientImpl struct {
	cc *http.Client
}

func NewEventsHTTPClient(client *http.Client) EventsHTTPClient {
	return &EventsHTTPClientImpl{client}
}

func (c *EventsHTTPClientImpl) CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...http.CallOption) (*CreateEventReply, error) {
	var out CreateEventReply
	pattern := "/api/v1/events"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationEventsCreateEvent))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}