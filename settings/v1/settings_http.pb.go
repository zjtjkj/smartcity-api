// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.3
// - protoc             v3.19.3
// source: api/settings/v1/settings.proto

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

const OperationSettingsCreateCameraAttr = "/api.settings.v1.Settings/CreateCameraAttr"
const OperationSettingsCreateGeneralParameters = "/api.settings.v1.Settings/CreateGeneralParameters"
const OperationSettingsCreateMapConfig = "/api.settings.v1.Settings/CreateMapConfig"
const OperationSettingsCreateModule = "/api.settings.v1.Settings/CreateModule"
const OperationSettingsCreateOperatorConfig = "/api.settings.v1.Settings/CreateOperatorConfig"
const OperationSettingsCreatePusherConfig = "/api.settings.v1.Settings/CreatePusherConfig"
const OperationSettingsCreateSystemInfo = "/api.settings.v1.Settings/CreateSystemInfo"
const OperationSettingsDeleteCameraAttr = "/api.settings.v1.Settings/DeleteCameraAttr"
const OperationSettingsDeleteGeneralParameters = "/api.settings.v1.Settings/DeleteGeneralParameters"
const OperationSettingsDeleteModule = "/api.settings.v1.Settings/DeleteModule"
const OperationSettingsDeleteSystemInfo = "/api.settings.v1.Settings/DeleteSystemInfo"
const OperationSettingsGetCameraAttr = "/api.settings.v1.Settings/GetCameraAttr"
const OperationSettingsGetMapConfig = "/api.settings.v1.Settings/GetMapConfig"
const OperationSettingsGetModule = "/api.settings.v1.Settings/GetModule"
const OperationSettingsGetOperatorConfig = "/api.settings.v1.Settings/GetOperatorConfig"
const OperationSettingsGetPusherConfig = "/api.settings.v1.Settings/GetPusherConfig"
const OperationSettingsGetSystemInfo = "/api.settings.v1.Settings/GetSystemInfo"
const OperationSettingsListCameraAttr = "/api.settings.v1.Settings/ListCameraAttr"
const OperationSettingsListCameraModules = "/api.settings.v1.Settings/ListCameraModules"
const OperationSettingsListGeneralParameters = "/api.settings.v1.Settings/ListGeneralParameters"
const OperationSettingsListModules = "/api.settings.v1.Settings/ListModules"
const OperationSettingsListRegionModules = "/api.settings.v1.Settings/ListRegionModules"
const OperationSettingsUpdateGeneralParameters = "/api.settings.v1.Settings/UpdateGeneralParameters"
const OperationSettingsUpdateModule = "/api.settings.v1.Settings/UpdateModule"

type SettingsHTTPServer interface {
	CreateCameraAttr(context.Context, *CreateCameraAttrRequest) (*CreateCameraAttrReply, error)
	CreateGeneralParameters(context.Context, *CreateGeneralParametersRequest) (*CreateGeneralParametersReply, error)
	CreateMapConfig(context.Context, *CreateMapConfigRequest) (*CreateMapConfigReply, error)
	CreateModule(context.Context, *CreateModuleRequest) (*CreateModuleReply, error)
	CreateOperatorConfig(context.Context, *CreateOperatorConfigRequest) (*CreateOperatorConfigReply, error)
	CreatePusherConfig(context.Context, *CreatePusherConfigRequest) (*CreatePusherConfigReply, error)
	CreateSystemInfo(context.Context, *CreateSystemInfoRequest) (*CreateSystemInfoReply, error)
	DeleteCameraAttr(context.Context, *DeleteCameraAttrRequest) (*DeleteCameraAttrReply, error)
	DeleteGeneralParameters(context.Context, *DeleteGeneralParametersRequest) (*DeleteGeneralParametersReply, error)
	DeleteModule(context.Context, *DeleteModuleRequest) (*DeleteModuleReply, error)
	DeleteSystemInfo(context.Context, *DeleteSystemInfoRequest) (*DeleteSystemInfoReply, error)
	GetCameraAttr(context.Context, *GetCameraAttrRequest) (*GetCameraAttrReply, error)
	GetMapConfig(context.Context, *GetMapConfigRequest) (*GetMapConfigReply, error)
	GetModule(context.Context, *GetModuleRequest) (*GetModuleReply, error)
	GetOperatorConfig(context.Context, *GetOperatorConfigRequest) (*GetOperatorConfigReply, error)
	GetPusherConfig(context.Context, *GetPusherConfigRequest) (*GetPusherConfigReply, error)
	GetSystemInfo(context.Context, *GetSystemInfoRequest) (*GetSystemInfoReply, error)
	ListCameraAttr(context.Context, *ListCameraAttrRequest) (*ListCameraAttrReply, error)
	ListCameraModules(context.Context, *ListCameraModulesRequest) (*ListCameraModulesReply, error)
	ListGeneralParameters(context.Context, *ListGeneralParametersRequest) (*ListGeneralParametersReply, error)
	ListModules(context.Context, *ListModulesRequest) (*ListModulesReply, error)
	ListRegionModules(context.Context, *ListRegionModulesRequest) (*ListRegionModulesReply, error)
	UpdateGeneralParameters(context.Context, *UpdateGeneralParametersRequest) (*UpdateGeneralParametersReply, error)
	UpdateModule(context.Context, *UpdateModuleRequest) (*UpdateModuleReply, error)
}

func RegisterSettingsHTTPServer(s *http.Server, srv SettingsHTTPServer) {
	r := s.Route("/")
	r.PUT("/api/v1/settings/pusher", _Settings_CreatePusherConfig0_HTTP_Handler(srv))
	r.GET("/api/v1/settings/pusher", _Settings_GetPusherConfig0_HTTP_Handler(srv))
	r.PUT("/api/v1/settings/operator", _Settings_CreateOperatorConfig0_HTTP_Handler(srv))
	r.GET("/api/v1/settings/operator", _Settings_GetOperatorConfig0_HTTP_Handler(srv))
	r.PUT("/api/v1/settings/map", _Settings_CreateMapConfig0_HTTP_Handler(srv))
	r.GET("/api/v1/settings/map", _Settings_GetMapConfig0_HTTP_Handler(srv))
	r.PUT("/api/v1/settings/system", _Settings_CreateSystemInfo0_HTTP_Handler(srv))
	r.GET("/api/v1/settings/system", _Settings_GetSystemInfo0_HTTP_Handler(srv))
	r.DELETE("/api/v1/settings/system", _Settings_DeleteSystemInfo0_HTTP_Handler(srv))
	r.PUT("/api/v1/settings/module", _Settings_CreateModule0_HTTP_Handler(srv))
	r.DELETE("/api/v1/settings/module/{id}", _Settings_DeleteModule0_HTTP_Handler(srv))
	r.POST("/api/v1/setting/module/{id}", _Settings_UpdateModule0_HTTP_Handler(srv))
	r.GET("/api/v1/settings/module/{id}", _Settings_GetModule0_HTTP_Handler(srv))
	r.GET("/api/v1/settings/module", _Settings_ListModules0_HTTP_Handler(srv))
	r.PUT("/api/v1/settings/icon", _Settings_CreateCameraAttr0_HTTP_Handler(srv))
	r.DELETE("/api/v1/settings/icon/{id}", _Settings_DeleteCameraAttr0_HTTP_Handler(srv))
	r.GET("/api/v1/settings/icon/{id}", _Settings_GetCameraAttr0_HTTP_Handler(srv))
	r.GET("/api/v1/settings/icons", _Settings_ListCameraAttr0_HTTP_Handler(srv))
	r.PUT("/api/v1/settings/parameter", _Settings_CreateGeneralParameters0_HTTP_Handler(srv))
	r.POST("/api/v1/settings/parameter/{id}", _Settings_UpdateGeneralParameters0_HTTP_Handler(srv))
	r.DELETE("/api/v1/settings/parameter/{id}", _Settings_DeleteGeneralParameters0_HTTP_Handler(srv))
	r.GET("/api/v1/settings/parameters", _Settings_ListGeneralParameters0_HTTP_Handler(srv))
	r.GET("/api/v1/settings/region/module", _Settings_ListRegionModules0_HTTP_Handler(srv))
	r.GET("/api/v1/settings/camera/module", _Settings_ListCameraModules0_HTTP_Handler(srv))
}

func _Settings_CreatePusherConfig0_HTTP_Handler(srv SettingsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreatePusherConfigRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSettingsCreatePusherConfig)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreatePusherConfig(ctx, req.(*CreatePusherConfigRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreatePusherConfigReply)
		return ctx.Result(200, reply)
	}
}

func _Settings_GetPusherConfig0_HTTP_Handler(srv SettingsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetPusherConfigRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSettingsGetPusherConfig)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPusherConfig(ctx, req.(*GetPusherConfigRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetPusherConfigReply)
		return ctx.Result(200, reply)
	}
}

func _Settings_CreateOperatorConfig0_HTTP_Handler(srv SettingsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateOperatorConfigRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSettingsCreateOperatorConfig)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateOperatorConfig(ctx, req.(*CreateOperatorConfigRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateOperatorConfigReply)
		return ctx.Result(200, reply)
	}
}

func _Settings_GetOperatorConfig0_HTTP_Handler(srv SettingsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetOperatorConfigRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSettingsGetOperatorConfig)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetOperatorConfig(ctx, req.(*GetOperatorConfigRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetOperatorConfigReply)
		return ctx.Result(200, reply)
	}
}

func _Settings_CreateMapConfig0_HTTP_Handler(srv SettingsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateMapConfigRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSettingsCreateMapConfig)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateMapConfig(ctx, req.(*CreateMapConfigRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateMapConfigReply)
		return ctx.Result(200, reply)
	}
}

func _Settings_GetMapConfig0_HTTP_Handler(srv SettingsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetMapConfigRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSettingsGetMapConfig)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetMapConfig(ctx, req.(*GetMapConfigRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetMapConfigReply)
		return ctx.Result(200, reply)
	}
}

func _Settings_CreateSystemInfo0_HTTP_Handler(srv SettingsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateSystemInfoRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSettingsCreateSystemInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateSystemInfo(ctx, req.(*CreateSystemInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateSystemInfoReply)
		return ctx.Result(200, reply)
	}
}

func _Settings_GetSystemInfo0_HTTP_Handler(srv SettingsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetSystemInfoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSettingsGetSystemInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSystemInfo(ctx, req.(*GetSystemInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetSystemInfoReply)
		return ctx.Result(200, reply)
	}
}

func _Settings_DeleteSystemInfo0_HTTP_Handler(srv SettingsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteSystemInfoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSettingsDeleteSystemInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteSystemInfo(ctx, req.(*DeleteSystemInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteSystemInfoReply)
		return ctx.Result(200, reply)
	}
}

func _Settings_CreateModule0_HTTP_Handler(srv SettingsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateModuleRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSettingsCreateModule)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateModule(ctx, req.(*CreateModuleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateModuleReply)
		return ctx.Result(200, reply)
	}
}

func _Settings_DeleteModule0_HTTP_Handler(srv SettingsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteModuleRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSettingsDeleteModule)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteModule(ctx, req.(*DeleteModuleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteModuleReply)
		return ctx.Result(200, reply)
	}
}

func _Settings_UpdateModule0_HTTP_Handler(srv SettingsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateModuleRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSettingsUpdateModule)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateModule(ctx, req.(*UpdateModuleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateModuleReply)
		return ctx.Result(200, reply)
	}
}

func _Settings_GetModule0_HTTP_Handler(srv SettingsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetModuleRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSettingsGetModule)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetModule(ctx, req.(*GetModuleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetModuleReply)
		return ctx.Result(200, reply)
	}
}

func _Settings_ListModules0_HTTP_Handler(srv SettingsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListModulesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSettingsListModules)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListModules(ctx, req.(*ListModulesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListModulesReply)
		return ctx.Result(200, reply)
	}
}

func _Settings_CreateCameraAttr0_HTTP_Handler(srv SettingsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateCameraAttrRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSettingsCreateCameraAttr)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateCameraAttr(ctx, req.(*CreateCameraAttrRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateCameraAttrReply)
		return ctx.Result(200, reply)
	}
}

func _Settings_DeleteCameraAttr0_HTTP_Handler(srv SettingsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteCameraAttrRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSettingsDeleteCameraAttr)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteCameraAttr(ctx, req.(*DeleteCameraAttrRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteCameraAttrReply)
		return ctx.Result(200, reply)
	}
}

func _Settings_GetCameraAttr0_HTTP_Handler(srv SettingsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCameraAttrRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSettingsGetCameraAttr)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCameraAttr(ctx, req.(*GetCameraAttrRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetCameraAttrReply)
		return ctx.Result(200, reply)
	}
}

func _Settings_ListCameraAttr0_HTTP_Handler(srv SettingsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListCameraAttrRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSettingsListCameraAttr)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListCameraAttr(ctx, req.(*ListCameraAttrRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListCameraAttrReply)
		return ctx.Result(200, reply)
	}
}

func _Settings_CreateGeneralParameters0_HTTP_Handler(srv SettingsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateGeneralParametersRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSettingsCreateGeneralParameters)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateGeneralParameters(ctx, req.(*CreateGeneralParametersRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateGeneralParametersReply)
		return ctx.Result(200, reply)
	}
}

func _Settings_UpdateGeneralParameters0_HTTP_Handler(srv SettingsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateGeneralParametersRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSettingsUpdateGeneralParameters)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateGeneralParameters(ctx, req.(*UpdateGeneralParametersRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateGeneralParametersReply)
		return ctx.Result(200, reply)
	}
}

func _Settings_DeleteGeneralParameters0_HTTP_Handler(srv SettingsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteGeneralParametersRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSettingsDeleteGeneralParameters)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteGeneralParameters(ctx, req.(*DeleteGeneralParametersRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteGeneralParametersReply)
		return ctx.Result(200, reply)
	}
}

func _Settings_ListGeneralParameters0_HTTP_Handler(srv SettingsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListGeneralParametersRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSettingsListGeneralParameters)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListGeneralParameters(ctx, req.(*ListGeneralParametersRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListGeneralParametersReply)
		return ctx.Result(200, reply)
	}
}

func _Settings_ListRegionModules0_HTTP_Handler(srv SettingsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListRegionModulesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSettingsListRegionModules)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListRegionModules(ctx, req.(*ListRegionModulesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListRegionModulesReply)
		return ctx.Result(200, reply)
	}
}

func _Settings_ListCameraModules0_HTTP_Handler(srv SettingsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListCameraModulesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSettingsListCameraModules)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListCameraModules(ctx, req.(*ListCameraModulesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListCameraModulesReply)
		return ctx.Result(200, reply)
	}
}

type SettingsHTTPClient interface {
	CreateCameraAttr(ctx context.Context, req *CreateCameraAttrRequest, opts ...http.CallOption) (rsp *CreateCameraAttrReply, err error)
	CreateGeneralParameters(ctx context.Context, req *CreateGeneralParametersRequest, opts ...http.CallOption) (rsp *CreateGeneralParametersReply, err error)
	CreateMapConfig(ctx context.Context, req *CreateMapConfigRequest, opts ...http.CallOption) (rsp *CreateMapConfigReply, err error)
	CreateModule(ctx context.Context, req *CreateModuleRequest, opts ...http.CallOption) (rsp *CreateModuleReply, err error)
	CreateOperatorConfig(ctx context.Context, req *CreateOperatorConfigRequest, opts ...http.CallOption) (rsp *CreateOperatorConfigReply, err error)
	CreatePusherConfig(ctx context.Context, req *CreatePusherConfigRequest, opts ...http.CallOption) (rsp *CreatePusherConfigReply, err error)
	CreateSystemInfo(ctx context.Context, req *CreateSystemInfoRequest, opts ...http.CallOption) (rsp *CreateSystemInfoReply, err error)
	DeleteCameraAttr(ctx context.Context, req *DeleteCameraAttrRequest, opts ...http.CallOption) (rsp *DeleteCameraAttrReply, err error)
	DeleteGeneralParameters(ctx context.Context, req *DeleteGeneralParametersRequest, opts ...http.CallOption) (rsp *DeleteGeneralParametersReply, err error)
	DeleteModule(ctx context.Context, req *DeleteModuleRequest, opts ...http.CallOption) (rsp *DeleteModuleReply, err error)
	DeleteSystemInfo(ctx context.Context, req *DeleteSystemInfoRequest, opts ...http.CallOption) (rsp *DeleteSystemInfoReply, err error)
	GetCameraAttr(ctx context.Context, req *GetCameraAttrRequest, opts ...http.CallOption) (rsp *GetCameraAttrReply, err error)
	GetMapConfig(ctx context.Context, req *GetMapConfigRequest, opts ...http.CallOption) (rsp *GetMapConfigReply, err error)
	GetModule(ctx context.Context, req *GetModuleRequest, opts ...http.CallOption) (rsp *GetModuleReply, err error)
	GetOperatorConfig(ctx context.Context, req *GetOperatorConfigRequest, opts ...http.CallOption) (rsp *GetOperatorConfigReply, err error)
	GetPusherConfig(ctx context.Context, req *GetPusherConfigRequest, opts ...http.CallOption) (rsp *GetPusherConfigReply, err error)
	GetSystemInfo(ctx context.Context, req *GetSystemInfoRequest, opts ...http.CallOption) (rsp *GetSystemInfoReply, err error)
	ListCameraAttr(ctx context.Context, req *ListCameraAttrRequest, opts ...http.CallOption) (rsp *ListCameraAttrReply, err error)
	ListCameraModules(ctx context.Context, req *ListCameraModulesRequest, opts ...http.CallOption) (rsp *ListCameraModulesReply, err error)
	ListGeneralParameters(ctx context.Context, req *ListGeneralParametersRequest, opts ...http.CallOption) (rsp *ListGeneralParametersReply, err error)
	ListModules(ctx context.Context, req *ListModulesRequest, opts ...http.CallOption) (rsp *ListModulesReply, err error)
	ListRegionModules(ctx context.Context, req *ListRegionModulesRequest, opts ...http.CallOption) (rsp *ListRegionModulesReply, err error)
	UpdateGeneralParameters(ctx context.Context, req *UpdateGeneralParametersRequest, opts ...http.CallOption) (rsp *UpdateGeneralParametersReply, err error)
	UpdateModule(ctx context.Context, req *UpdateModuleRequest, opts ...http.CallOption) (rsp *UpdateModuleReply, err error)
}

type SettingsHTTPClientImpl struct {
	cc *http.Client
}

func NewSettingsHTTPClient(client *http.Client) SettingsHTTPClient {
	return &SettingsHTTPClientImpl{client}
}

func (c *SettingsHTTPClientImpl) CreateCameraAttr(ctx context.Context, in *CreateCameraAttrRequest, opts ...http.CallOption) (*CreateCameraAttrReply, error) {
	var out CreateCameraAttrReply
	pattern := "/api/v1/settings/icon"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSettingsCreateCameraAttr))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SettingsHTTPClientImpl) CreateGeneralParameters(ctx context.Context, in *CreateGeneralParametersRequest, opts ...http.CallOption) (*CreateGeneralParametersReply, error) {
	var out CreateGeneralParametersReply
	pattern := "/api/v1/settings/parameter"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSettingsCreateGeneralParameters))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SettingsHTTPClientImpl) CreateMapConfig(ctx context.Context, in *CreateMapConfigRequest, opts ...http.CallOption) (*CreateMapConfigReply, error) {
	var out CreateMapConfigReply
	pattern := "/api/v1/settings/map"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSettingsCreateMapConfig))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SettingsHTTPClientImpl) CreateModule(ctx context.Context, in *CreateModuleRequest, opts ...http.CallOption) (*CreateModuleReply, error) {
	var out CreateModuleReply
	pattern := "/api/v1/settings/module"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSettingsCreateModule))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SettingsHTTPClientImpl) CreateOperatorConfig(ctx context.Context, in *CreateOperatorConfigRequest, opts ...http.CallOption) (*CreateOperatorConfigReply, error) {
	var out CreateOperatorConfigReply
	pattern := "/api/v1/settings/operator"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSettingsCreateOperatorConfig))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SettingsHTTPClientImpl) CreatePusherConfig(ctx context.Context, in *CreatePusherConfigRequest, opts ...http.CallOption) (*CreatePusherConfigReply, error) {
	var out CreatePusherConfigReply
	pattern := "/api/v1/settings/pusher"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSettingsCreatePusherConfig))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SettingsHTTPClientImpl) CreateSystemInfo(ctx context.Context, in *CreateSystemInfoRequest, opts ...http.CallOption) (*CreateSystemInfoReply, error) {
	var out CreateSystemInfoReply
	pattern := "/api/v1/settings/system"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSettingsCreateSystemInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SettingsHTTPClientImpl) DeleteCameraAttr(ctx context.Context, in *DeleteCameraAttrRequest, opts ...http.CallOption) (*DeleteCameraAttrReply, error) {
	var out DeleteCameraAttrReply
	pattern := "/api/v1/settings/icon/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSettingsDeleteCameraAttr))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SettingsHTTPClientImpl) DeleteGeneralParameters(ctx context.Context, in *DeleteGeneralParametersRequest, opts ...http.CallOption) (*DeleteGeneralParametersReply, error) {
	var out DeleteGeneralParametersReply
	pattern := "/api/v1/settings/parameter/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSettingsDeleteGeneralParameters))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SettingsHTTPClientImpl) DeleteModule(ctx context.Context, in *DeleteModuleRequest, opts ...http.CallOption) (*DeleteModuleReply, error) {
	var out DeleteModuleReply
	pattern := "/api/v1/settings/module/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSettingsDeleteModule))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SettingsHTTPClientImpl) DeleteSystemInfo(ctx context.Context, in *DeleteSystemInfoRequest, opts ...http.CallOption) (*DeleteSystemInfoReply, error) {
	var out DeleteSystemInfoReply
	pattern := "/api/v1/settings/system"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSettingsDeleteSystemInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SettingsHTTPClientImpl) GetCameraAttr(ctx context.Context, in *GetCameraAttrRequest, opts ...http.CallOption) (*GetCameraAttrReply, error) {
	var out GetCameraAttrReply
	pattern := "/api/v1/settings/icon/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSettingsGetCameraAttr))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SettingsHTTPClientImpl) GetMapConfig(ctx context.Context, in *GetMapConfigRequest, opts ...http.CallOption) (*GetMapConfigReply, error) {
	var out GetMapConfigReply
	pattern := "/api/v1/settings/map"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSettingsGetMapConfig))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SettingsHTTPClientImpl) GetModule(ctx context.Context, in *GetModuleRequest, opts ...http.CallOption) (*GetModuleReply, error) {
	var out GetModuleReply
	pattern := "/api/v1/settings/module/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSettingsGetModule))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SettingsHTTPClientImpl) GetOperatorConfig(ctx context.Context, in *GetOperatorConfigRequest, opts ...http.CallOption) (*GetOperatorConfigReply, error) {
	var out GetOperatorConfigReply
	pattern := "/api/v1/settings/operator"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSettingsGetOperatorConfig))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SettingsHTTPClientImpl) GetPusherConfig(ctx context.Context, in *GetPusherConfigRequest, opts ...http.CallOption) (*GetPusherConfigReply, error) {
	var out GetPusherConfigReply
	pattern := "/api/v1/settings/pusher"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSettingsGetPusherConfig))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SettingsHTTPClientImpl) GetSystemInfo(ctx context.Context, in *GetSystemInfoRequest, opts ...http.CallOption) (*GetSystemInfoReply, error) {
	var out GetSystemInfoReply
	pattern := "/api/v1/settings/system"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSettingsGetSystemInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SettingsHTTPClientImpl) ListCameraAttr(ctx context.Context, in *ListCameraAttrRequest, opts ...http.CallOption) (*ListCameraAttrReply, error) {
	var out ListCameraAttrReply
	pattern := "/api/v1/settings/icons"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSettingsListCameraAttr))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SettingsHTTPClientImpl) ListCameraModules(ctx context.Context, in *ListCameraModulesRequest, opts ...http.CallOption) (*ListCameraModulesReply, error) {
	var out ListCameraModulesReply
	pattern := "/api/v1/settings/camera/module"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSettingsListCameraModules))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SettingsHTTPClientImpl) ListGeneralParameters(ctx context.Context, in *ListGeneralParametersRequest, opts ...http.CallOption) (*ListGeneralParametersReply, error) {
	var out ListGeneralParametersReply
	pattern := "/api/v1/settings/parameters"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSettingsListGeneralParameters))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SettingsHTTPClientImpl) ListModules(ctx context.Context, in *ListModulesRequest, opts ...http.CallOption) (*ListModulesReply, error) {
	var out ListModulesReply
	pattern := "/api/v1/settings/module"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSettingsListModules))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SettingsHTTPClientImpl) ListRegionModules(ctx context.Context, in *ListRegionModulesRequest, opts ...http.CallOption) (*ListRegionModulesReply, error) {
	var out ListRegionModulesReply
	pattern := "/api/v1/settings/region/module"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSettingsListRegionModules))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SettingsHTTPClientImpl) UpdateGeneralParameters(ctx context.Context, in *UpdateGeneralParametersRequest, opts ...http.CallOption) (*UpdateGeneralParametersReply, error) {
	var out UpdateGeneralParametersReply
	pattern := "/api/v1/settings/parameter/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSettingsUpdateGeneralParameters))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SettingsHTTPClientImpl) UpdateModule(ctx context.Context, in *UpdateModuleRequest, opts ...http.CallOption) (*UpdateModuleReply, error) {
	var out UpdateModuleReply
	pattern := "/api/v1/setting/module/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSettingsUpdateModule))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
