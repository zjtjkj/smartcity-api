// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.4.0
// - protoc             v3.20.0
// source: api/user-domain/v1/user-domain.proto

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

const OperationUserDomainCreateUserDomain = "/api.userdomain.v1.UserDomain/CreateUserDomain"
const OperationUserDomainDeleteUserDomain = "/api.userdomain.v1.UserDomain/DeleteUserDomain"
const OperationUserDomainFindUserDomain = "/api.userdomain.v1.UserDomain/FindUserDomain"
const OperationUserDomainGetUserDomain = "/api.userdomain.v1.UserDomain/GetUserDomain"
const OperationUserDomainListUserDomain = "/api.userdomain.v1.UserDomain/ListUserDomain"
const OperationUserDomainListUserDomainByUser = "/api.userdomain.v1.UserDomain/ListUserDomainByUser"
const OperationUserDomainUpdateUserDomain = "/api.userdomain.v1.UserDomain/UpdateUserDomain"

type UserDomainHTTPServer interface {
	CreateUserDomain(context.Context, *CreateUserDomainRequest) (*CreateUserDomainReply, error)
	DeleteUserDomain(context.Context, *DeleteUserDomainRequest) (*DeleteUserDomainReply, error)
	FindUserDomain(context.Context, *FindUserDomainRequest) (*FindUserDomainReply, error)
	GetUserDomain(context.Context, *GetUserDomainRequest) (*GetUserDomainReply, error)
	ListUserDomain(context.Context, *ListUserDomainRequest) (*ListUserDomainReply, error)
	ListUserDomainByUser(context.Context, *ListUserDomainByUserRequest) (*ListUserDomainByUserReply, error)
	UpdateUserDomain(context.Context, *UpdateUserDomainRequest) (*UpdateUserDomainReply, error)
}

func RegisterUserDomainHTTPServer(s *http.Server, srv UserDomainHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/userdomain/create", _UserDomain_CreateUserDomain0_HTTP_Handler(srv))
	r.POST("/api/v1/userdomain/update", _UserDomain_UpdateUserDomain0_HTTP_Handler(srv))
	r.POST("/api/v1/userdomain/delete", _UserDomain_DeleteUserDomain0_HTTP_Handler(srv))
	r.POST("/api/v1/userdomain/get", _UserDomain_GetUserDomain0_HTTP_Handler(srv))
	r.POST("/api/v1/userdomain/find", _UserDomain_FindUserDomain0_HTTP_Handler(srv))
	r.GET("/api/v1/user/list", _UserDomain_ListUserDomain0_HTTP_Handler(srv))
	r.GET("/api/v1/user/list/user", _UserDomain_ListUserDomainByUser0_HTTP_Handler(srv))
}

func _UserDomain_CreateUserDomain0_HTTP_Handler(srv UserDomainHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateUserDomainRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserDomainCreateUserDomain)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateUserDomain(ctx, req.(*CreateUserDomainRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateUserDomainReply)
		return ctx.Result(200, reply)
	}
}

func _UserDomain_UpdateUserDomain0_HTTP_Handler(srv UserDomainHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateUserDomainRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserDomainUpdateUserDomain)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUserDomain(ctx, req.(*UpdateUserDomainRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateUserDomainReply)
		return ctx.Result(200, reply)
	}
}

func _UserDomain_DeleteUserDomain0_HTTP_Handler(srv UserDomainHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteUserDomainRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserDomainDeleteUserDomain)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteUserDomain(ctx, req.(*DeleteUserDomainRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteUserDomainReply)
		return ctx.Result(200, reply)
	}
}

func _UserDomain_GetUserDomain0_HTTP_Handler(srv UserDomainHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserDomainRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserDomainGetUserDomain)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserDomain(ctx, req.(*GetUserDomainRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserDomainReply)
		return ctx.Result(200, reply)
	}
}

func _UserDomain_FindUserDomain0_HTTP_Handler(srv UserDomainHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FindUserDomainRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserDomainFindUserDomain)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FindUserDomain(ctx, req.(*FindUserDomainRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*FindUserDomainReply)
		return ctx.Result(200, reply)
	}
}

func _UserDomain_ListUserDomain0_HTTP_Handler(srv UserDomainHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListUserDomainRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserDomainListUserDomain)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListUserDomain(ctx, req.(*ListUserDomainRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListUserDomainReply)
		return ctx.Result(200, reply)
	}
}

func _UserDomain_ListUserDomainByUser0_HTTP_Handler(srv UserDomainHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListUserDomainByUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserDomainListUserDomainByUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListUserDomainByUser(ctx, req.(*ListUserDomainByUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListUserDomainByUserReply)
		return ctx.Result(200, reply)
	}
}

type UserDomainHTTPClient interface {
	CreateUserDomain(ctx context.Context, req *CreateUserDomainRequest, opts ...http.CallOption) (rsp *CreateUserDomainReply, err error)
	DeleteUserDomain(ctx context.Context, req *DeleteUserDomainRequest, opts ...http.CallOption) (rsp *DeleteUserDomainReply, err error)
	FindUserDomain(ctx context.Context, req *FindUserDomainRequest, opts ...http.CallOption) (rsp *FindUserDomainReply, err error)
	GetUserDomain(ctx context.Context, req *GetUserDomainRequest, opts ...http.CallOption) (rsp *GetUserDomainReply, err error)
	ListUserDomain(ctx context.Context, req *ListUserDomainRequest, opts ...http.CallOption) (rsp *ListUserDomainReply, err error)
	ListUserDomainByUser(ctx context.Context, req *ListUserDomainByUserRequest, opts ...http.CallOption) (rsp *ListUserDomainByUserReply, err error)
	UpdateUserDomain(ctx context.Context, req *UpdateUserDomainRequest, opts ...http.CallOption) (rsp *UpdateUserDomainReply, err error)
}

type UserDomainHTTPClientImpl struct {
	cc *http.Client
}

func NewUserDomainHTTPClient(client *http.Client) UserDomainHTTPClient {
	return &UserDomainHTTPClientImpl{client}
}

func (c *UserDomainHTTPClientImpl) CreateUserDomain(ctx context.Context, in *CreateUserDomainRequest, opts ...http.CallOption) (*CreateUserDomainReply, error) {
	var out CreateUserDomainReply
	pattern := "/api/v1/userdomain/create"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserDomainCreateUserDomain))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserDomainHTTPClientImpl) DeleteUserDomain(ctx context.Context, in *DeleteUserDomainRequest, opts ...http.CallOption) (*DeleteUserDomainReply, error) {
	var out DeleteUserDomainReply
	pattern := "/api/v1/userdomain/delete"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserDomainDeleteUserDomain))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserDomainHTTPClientImpl) FindUserDomain(ctx context.Context, in *FindUserDomainRequest, opts ...http.CallOption) (*FindUserDomainReply, error) {
	var out FindUserDomainReply
	pattern := "/api/v1/userdomain/find"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserDomainFindUserDomain))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserDomainHTTPClientImpl) GetUserDomain(ctx context.Context, in *GetUserDomainRequest, opts ...http.CallOption) (*GetUserDomainReply, error) {
	var out GetUserDomainReply
	pattern := "/api/v1/userdomain/get"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserDomainGetUserDomain))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserDomainHTTPClientImpl) ListUserDomain(ctx context.Context, in *ListUserDomainRequest, opts ...http.CallOption) (*ListUserDomainReply, error) {
	var out ListUserDomainReply
	pattern := "/api/v1/user/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserDomainListUserDomain))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserDomainHTTPClientImpl) ListUserDomainByUser(ctx context.Context, in *ListUserDomainByUserRequest, opts ...http.CallOption) (*ListUserDomainByUserReply, error) {
	var out ListUserDomainByUserReply
	pattern := "/api/v1/user/list/user"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserDomainListUserDomainByUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserDomainHTTPClientImpl) UpdateUserDomain(ctx context.Context, in *UpdateUserDomainRequest, opts ...http.CallOption) (*UpdateUserDomainReply, error) {
	var out UpdateUserDomainReply
	pattern := "/api/v1/userdomain/update"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserDomainUpdateUserDomain))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
