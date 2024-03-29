// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.3
// - protoc             v3.14.0
// source: api/retrieval/v1/retrieval.proto

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

const OperationRetrievalDeleteEvent = "/api.retrieval.Retrieval/DeleteEvent"
const OperationRetrievalDeleteTag = "/api.retrieval.Retrieval/DeleteTag"
const OperationRetrievalFindEventCount = "/api.retrieval.Retrieval/FindEventCount"
const OperationRetrievalFindEvents = "/api.retrieval.Retrieval/FindEvents"
const OperationRetrievalFindRegionEventCount = "/api.retrieval.Retrieval/FindRegionEventCount"
const OperationRetrievalGetChainEvidence = "/api.retrieval.Retrieval/GetChainEvidence"
const OperationRetrievalGetEvent = "/api.retrieval.Retrieval/GetEvent"
const OperationRetrievalGetEventByIndex = "/api.retrieval.Retrieval/GetEventByIndex"
const OperationRetrievalGetImage = "/api.retrieval.Retrieval/GetImage"
const OperationRetrievalListTag = "/api.retrieval.Retrieval/ListTag"
const OperationRetrievalMissionLatestInfo = "/api.retrieval.Retrieval/MissionLatestInfo"
const OperationRetrievalSetTags = "/api.retrieval.Retrieval/SetTags"

type RetrievalHTTPServer interface {
	DeleteEvent(context.Context, *DeleteEventRequest) (*DeleteEventReply, error)
	DeleteTag(context.Context, *DeleteTagRequest) (*DeleteTagReply, error)
	FindEventCount(context.Context, *FindEventCountRequest) (*FindEventCountReply, error)
	FindEvents(context.Context, *FindEventsRequest) (*FindEventsReply, error)
	FindRegionEventCount(context.Context, *FindRegionEventCountRequest) (*FindRegionEventCountReply, error)
	GetChainEvidence(context.Context, *GetChainEvidenceRequest) (*GetChainEvidenceReply, error)
	GetEvent(context.Context, *GetEventRequest) (*GetEventReply, error)
	GetEventByIndex(context.Context, *GetEventByIndexRequest) (*GetEventByIndexReply, error)
	GetImage(context.Context, *GetImageRequest) (*GetImageReply, error)
	ListTag(context.Context, *ListTagRequest) (*ListTagReply, error)
	MissionLatestInfo(context.Context, *MissionLatestInfoRequest) (*MissionLatestInfoReply, error)
	SetTags(context.Context, *SetTagsRequest) (*SetTagReply, error)
}

func RegisterRetrievalHTTPServer(s *http.Server, srv RetrievalHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/report/time/count", _Retrieval_FindEventCount0_HTTP_Handler(srv))
	r.POST("/api/v1/report/region/count", _Retrieval_FindRegionEventCount0_HTTP_Handler(srv))
	r.POST("/api/v1/event/find", _Retrieval_FindEvents0_HTTP_Handler(srv))
	r.POST("/api/v1/event/get", _Retrieval_GetEvent0_HTTP_Handler(srv))
	r.POST("/api/v1/event/chain/find", _Retrieval_GetChainEvidence0_HTTP_Handler(srv))
	r.POST("/api/v1/image/get", _Retrieval_GetImage0_HTTP_Handler(srv))
	r.POST("/api/v1/event/delete", _Retrieval_DeleteEvent0_HTTP_Handler(srv))
	r.POST("/api/v1/event/delete", _Retrieval_MissionLatestInfo0_HTTP_Handler(srv))
	r.POST("/api/v1/event/tags", _Retrieval_SetTags0_HTTP_Handler(srv))
	r.POST("/api/v1/event/tags/delete", _Retrieval_DeleteTag0_HTTP_Handler(srv))
	r.POST("/api/v1/event/tags/list", _Retrieval_ListTag0_HTTP_Handler(srv))
	r.POST("/api/v1/event/get/index", _Retrieval_GetEventByIndex0_HTTP_Handler(srv))
}

func _Retrieval_FindEventCount0_HTTP_Handler(srv RetrievalHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FindEventCountRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRetrievalFindEventCount)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FindEventCount(ctx, req.(*FindEventCountRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*FindEventCountReply)
		return ctx.Result(200, reply)
	}
}

func _Retrieval_FindRegionEventCount0_HTTP_Handler(srv RetrievalHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FindRegionEventCountRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRetrievalFindRegionEventCount)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FindRegionEventCount(ctx, req.(*FindRegionEventCountRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*FindRegionEventCountReply)
		return ctx.Result(200, reply)
	}
}

func _Retrieval_FindEvents0_HTTP_Handler(srv RetrievalHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FindEventsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRetrievalFindEvents)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FindEvents(ctx, req.(*FindEventsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*FindEventsReply)
		return ctx.Result(200, reply)
	}
}

func _Retrieval_GetEvent0_HTTP_Handler(srv RetrievalHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetEventRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRetrievalGetEvent)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetEvent(ctx, req.(*GetEventRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetEventReply)
		return ctx.Result(200, reply)
	}
}

func _Retrieval_GetChainEvidence0_HTTP_Handler(srv RetrievalHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetChainEvidenceRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRetrievalGetChainEvidence)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetChainEvidence(ctx, req.(*GetChainEvidenceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetChainEvidenceReply)
		return ctx.Result(200, reply)
	}
}

func _Retrieval_GetImage0_HTTP_Handler(srv RetrievalHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetImageRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRetrievalGetImage)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetImage(ctx, req.(*GetImageRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetImageReply)
		return ctx.Result(200, reply)
	}
}

func _Retrieval_DeleteEvent0_HTTP_Handler(srv RetrievalHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteEventRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRetrievalDeleteEvent)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteEvent(ctx, req.(*DeleteEventRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteEventReply)
		return ctx.Result(200, reply)
	}
}

func _Retrieval_MissionLatestInfo0_HTTP_Handler(srv RetrievalHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in MissionLatestInfoRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRetrievalMissionLatestInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.MissionLatestInfo(ctx, req.(*MissionLatestInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MissionLatestInfoReply)
		return ctx.Result(200, reply)
	}
}

func _Retrieval_SetTags0_HTTP_Handler(srv RetrievalHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SetTagsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRetrievalSetTags)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SetTags(ctx, req.(*SetTagsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SetTagReply)
		return ctx.Result(200, reply)
	}
}

func _Retrieval_DeleteTag0_HTTP_Handler(srv RetrievalHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteTagRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRetrievalDeleteTag)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteTag(ctx, req.(*DeleteTagRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteTagReply)
		return ctx.Result(200, reply)
	}
}

func _Retrieval_ListTag0_HTTP_Handler(srv RetrievalHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListTagRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRetrievalListTag)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListTag(ctx, req.(*ListTagRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListTagReply)
		return ctx.Result(200, reply)
	}
}

func _Retrieval_GetEventByIndex0_HTTP_Handler(srv RetrievalHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetEventByIndexRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRetrievalGetEventByIndex)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetEventByIndex(ctx, req.(*GetEventByIndexRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetEventByIndexReply)
		return ctx.Result(200, reply)
	}
}

type RetrievalHTTPClient interface {
	DeleteEvent(ctx context.Context, req *DeleteEventRequest, opts ...http.CallOption) (rsp *DeleteEventReply, err error)
	DeleteTag(ctx context.Context, req *DeleteTagRequest, opts ...http.CallOption) (rsp *DeleteTagReply, err error)
	FindEventCount(ctx context.Context, req *FindEventCountRequest, opts ...http.CallOption) (rsp *FindEventCountReply, err error)
	FindEvents(ctx context.Context, req *FindEventsRequest, opts ...http.CallOption) (rsp *FindEventsReply, err error)
	FindRegionEventCount(ctx context.Context, req *FindRegionEventCountRequest, opts ...http.CallOption) (rsp *FindRegionEventCountReply, err error)
	GetChainEvidence(ctx context.Context, req *GetChainEvidenceRequest, opts ...http.CallOption) (rsp *GetChainEvidenceReply, err error)
	GetEvent(ctx context.Context, req *GetEventRequest, opts ...http.CallOption) (rsp *GetEventReply, err error)
	GetEventByIndex(ctx context.Context, req *GetEventByIndexRequest, opts ...http.CallOption) (rsp *GetEventByIndexReply, err error)
	GetImage(ctx context.Context, req *GetImageRequest, opts ...http.CallOption) (rsp *GetImageReply, err error)
	ListTag(ctx context.Context, req *ListTagRequest, opts ...http.CallOption) (rsp *ListTagReply, err error)
	MissionLatestInfo(ctx context.Context, req *MissionLatestInfoRequest, opts ...http.CallOption) (rsp *MissionLatestInfoReply, err error)
	SetTags(ctx context.Context, req *SetTagsRequest, opts ...http.CallOption) (rsp *SetTagReply, err error)
}

type RetrievalHTTPClientImpl struct {
	cc *http.Client
}

func NewRetrievalHTTPClient(client *http.Client) RetrievalHTTPClient {
	return &RetrievalHTTPClientImpl{client}
}

func (c *RetrievalHTTPClientImpl) DeleteEvent(ctx context.Context, in *DeleteEventRequest, opts ...http.CallOption) (*DeleteEventReply, error) {
	var out DeleteEventReply
	pattern := "/api/v1/event/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRetrievalDeleteEvent))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RetrievalHTTPClientImpl) DeleteTag(ctx context.Context, in *DeleteTagRequest, opts ...http.CallOption) (*DeleteTagReply, error) {
	var out DeleteTagReply
	pattern := "/api/v1/event/tags/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRetrievalDeleteTag))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RetrievalHTTPClientImpl) FindEventCount(ctx context.Context, in *FindEventCountRequest, opts ...http.CallOption) (*FindEventCountReply, error) {
	var out FindEventCountReply
	pattern := "/api/v1/report/time/count"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRetrievalFindEventCount))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RetrievalHTTPClientImpl) FindEvents(ctx context.Context, in *FindEventsRequest, opts ...http.CallOption) (*FindEventsReply, error) {
	var out FindEventsReply
	pattern := "/api/v1/event/find"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRetrievalFindEvents))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RetrievalHTTPClientImpl) FindRegionEventCount(ctx context.Context, in *FindRegionEventCountRequest, opts ...http.CallOption) (*FindRegionEventCountReply, error) {
	var out FindRegionEventCountReply
	pattern := "/api/v1/report/region/count"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRetrievalFindRegionEventCount))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RetrievalHTTPClientImpl) GetChainEvidence(ctx context.Context, in *GetChainEvidenceRequest, opts ...http.CallOption) (*GetChainEvidenceReply, error) {
	var out GetChainEvidenceReply
	pattern := "/api/v1/event/chain/find"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRetrievalGetChainEvidence))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RetrievalHTTPClientImpl) GetEvent(ctx context.Context, in *GetEventRequest, opts ...http.CallOption) (*GetEventReply, error) {
	var out GetEventReply
	pattern := "/api/v1/event/get"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRetrievalGetEvent))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RetrievalHTTPClientImpl) GetEventByIndex(ctx context.Context, in *GetEventByIndexRequest, opts ...http.CallOption) (*GetEventByIndexReply, error) {
	var out GetEventByIndexReply
	pattern := "/api/v1/event/get/index"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRetrievalGetEventByIndex))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RetrievalHTTPClientImpl) GetImage(ctx context.Context, in *GetImageRequest, opts ...http.CallOption) (*GetImageReply, error) {
	var out GetImageReply
	pattern := "/api/v1/image/get"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRetrievalGetImage))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RetrievalHTTPClientImpl) ListTag(ctx context.Context, in *ListTagRequest, opts ...http.CallOption) (*ListTagReply, error) {
	var out ListTagReply
	pattern := "/api/v1/event/tags/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRetrievalListTag))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RetrievalHTTPClientImpl) MissionLatestInfo(ctx context.Context, in *MissionLatestInfoRequest, opts ...http.CallOption) (*MissionLatestInfoReply, error) {
	var out MissionLatestInfoReply
	pattern := "/api/v1/event/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRetrievalMissionLatestInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RetrievalHTTPClientImpl) SetTags(ctx context.Context, in *SetTagsRequest, opts ...http.CallOption) (*SetTagReply, error) {
	var out SetTagReply
	pattern := "/api/v1/event/tags"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRetrievalSetTags))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
