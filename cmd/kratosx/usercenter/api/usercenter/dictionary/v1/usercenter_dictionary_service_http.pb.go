// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v4.24.4
// source: usercenter_dictionary_service.proto

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

const OperationUsercenterCreateDictionary = "/usercenter.api.usercenter.dictionary.v1.Usercenter/CreateDictionary"
const OperationUsercenterDeleteDictionary = "/usercenter.api.usercenter.dictionary.v1.Usercenter/DeleteDictionary"
const OperationUsercenterDeleteTrashDictionary = "/usercenter.api.usercenter.dictionary.v1.Usercenter/DeleteTrashDictionary"
const OperationUsercenterGetDictionary = "/usercenter.api.usercenter.dictionary.v1.Usercenter/GetDictionary"
const OperationUsercenterGetTrashDictionary = "/usercenter.api.usercenter.dictionary.v1.Usercenter/GetTrashDictionary"
const OperationUsercenterListDictionary = "/usercenter.api.usercenter.dictionary.v1.Usercenter/ListDictionary"
const OperationUsercenterListTrashDictionary = "/usercenter.api.usercenter.dictionary.v1.Usercenter/ListTrashDictionary"
const OperationUsercenterRevertTrashDictionary = "/usercenter.api.usercenter.dictionary.v1.Usercenter/RevertTrashDictionary"
const OperationUsercenterUpdateDictionary = "/usercenter.api.usercenter.dictionary.v1.Usercenter/UpdateDictionary"

type UsercenterHTTPServer interface {
	// CreateDictionary CreateDictionary 创建字典信息
	CreateDictionary(context.Context, *CreateDictionaryRequest) (*CreateDictionaryReply, error)
	// DeleteDictionary DeleteDictionary 删除字典信息
	DeleteDictionary(context.Context, *DeleteDictionaryRequest) (*DeleteDictionaryReply, error)
	// DeleteTrashDictionary DeleteTrashDictionary 彻底删除字典信息
	DeleteTrashDictionary(context.Context, *DeleteTrashDictionaryRequest) (*DeleteTrashDictionaryReply, error)
	// GetDictionary GetDictionary 获取指定的字典信息
	GetDictionary(context.Context, *GetDictionaryRequest) (*GetDictionaryReply, error)
	// GetTrashDictionary GetTrashDictionary 查看指定字典信息回收站数据
	GetTrashDictionary(context.Context, *GetTrashDictionaryRequest) (*GetTrashDictionaryReply, error)
	// ListDictionary ListDictionary 获取字典信息列表
	ListDictionary(context.Context, *ListDictionaryRequest) (*ListDictionaryReply, error)
	// ListTrashDictionary ListTrashDictionary 查看字典信息列表回收站数据
	ListTrashDictionary(context.Context, *ListTrashDictionaryRequest) (*ListTrashDictionaryReply, error)
	// RevertTrashDictionary RevertTrashDictionary 还原字典信息
	RevertTrashDictionary(context.Context, *RevertTrashDictionaryRequest) (*RevertTrashDictionaryReply, error)
	// UpdateDictionary UpdateDictionary 更新字典信息
	UpdateDictionary(context.Context, *UpdateDictionaryRequest) (*UpdateDictionaryReply, error)
}

func RegisterUsercenterHTTPServer(s *http.Server, srv UsercenterHTTPServer) {
	r := s.Route("/")
	r.GET("/usercenter/api/v1/dictionary", _Usercenter_GetDictionary0_HTTP_Handler(srv))
	r.GET("/usercenter/api/v1/dictionaries", _Usercenter_ListDictionary0_HTTP_Handler(srv))
	r.POST("/usercenter/api/v1/dictionary", _Usercenter_CreateDictionary0_HTTP_Handler(srv))
	r.PUT("/usercenter/api/v1/dictionary", _Usercenter_UpdateDictionary0_HTTP_Handler(srv))
	r.DELETE("/usercenter/api/v1/dictionary", _Usercenter_DeleteDictionary0_HTTP_Handler(srv))
	r.GET("/usercenter/api/v1/dictionary/trash", _Usercenter_GetTrashDictionary0_HTTP_Handler(srv))
	r.GET("/usercenter/api/v1/dictionary/trashes", _Usercenter_ListTrashDictionary0_HTTP_Handler(srv))
	r.DELETE("/usercenter/api/v1/dictionary/trash", _Usercenter_DeleteTrashDictionary0_HTTP_Handler(srv))
	r.PUT("/usercenter/api/v1/dictionary/trash", _Usercenter_RevertTrashDictionary0_HTTP_Handler(srv))
}

func _Usercenter_GetDictionary0_HTTP_Handler(srv UsercenterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetDictionaryRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUsercenterGetDictionary)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.GetDictionary(ctx, req.(*GetDictionaryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetDictionaryReply)
		return ctx.Result(200, reply)
	}
}

func _Usercenter_ListDictionary0_HTTP_Handler(srv UsercenterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListDictionaryRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUsercenterListDictionary)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.ListDictionary(ctx, req.(*ListDictionaryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListDictionaryReply)
		return ctx.Result(200, reply)
	}
}

func _Usercenter_CreateDictionary0_HTTP_Handler(srv UsercenterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateDictionaryRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUsercenterCreateDictionary)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.CreateDictionary(ctx, req.(*CreateDictionaryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateDictionaryReply)
		return ctx.Result(200, reply)
	}
}

func _Usercenter_UpdateDictionary0_HTTP_Handler(srv UsercenterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateDictionaryRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUsercenterUpdateDictionary)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.UpdateDictionary(ctx, req.(*UpdateDictionaryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateDictionaryReply)
		return ctx.Result(200, reply)
	}
}

func _Usercenter_DeleteDictionary0_HTTP_Handler(srv UsercenterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteDictionaryRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUsercenterDeleteDictionary)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.DeleteDictionary(ctx, req.(*DeleteDictionaryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteDictionaryReply)
		return ctx.Result(200, reply)
	}
}

func _Usercenter_GetTrashDictionary0_HTTP_Handler(srv UsercenterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetTrashDictionaryRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUsercenterGetTrashDictionary)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.GetTrashDictionary(ctx, req.(*GetTrashDictionaryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetTrashDictionaryReply)
		return ctx.Result(200, reply)
	}
}

func _Usercenter_ListTrashDictionary0_HTTP_Handler(srv UsercenterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListTrashDictionaryRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUsercenterListTrashDictionary)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.ListTrashDictionary(ctx, req.(*ListTrashDictionaryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListTrashDictionaryReply)
		return ctx.Result(200, reply)
	}
}

func _Usercenter_DeleteTrashDictionary0_HTTP_Handler(srv UsercenterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteTrashDictionaryRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUsercenterDeleteTrashDictionary)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.DeleteTrashDictionary(ctx, req.(*DeleteTrashDictionaryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteTrashDictionaryReply)
		return ctx.Result(200, reply)
	}
}

func _Usercenter_RevertTrashDictionary0_HTTP_Handler(srv UsercenterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RevertTrashDictionaryRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUsercenterRevertTrashDictionary)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.RevertTrashDictionary(ctx, req.(*RevertTrashDictionaryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RevertTrashDictionaryReply)
		return ctx.Result(200, reply)
	}
}

type UsercenterHTTPClient interface {
	CreateDictionary(ctx context.Context, req *CreateDictionaryRequest, opts ...http.CallOption) (rsp *CreateDictionaryReply, err error)
	DeleteDictionary(ctx context.Context, req *DeleteDictionaryRequest, opts ...http.CallOption) (rsp *DeleteDictionaryReply, err error)
	DeleteTrashDictionary(ctx context.Context, req *DeleteTrashDictionaryRequest, opts ...http.CallOption) (rsp *DeleteTrashDictionaryReply, err error)
	GetDictionary(ctx context.Context, req *GetDictionaryRequest, opts ...http.CallOption) (rsp *GetDictionaryReply, err error)
	GetTrashDictionary(ctx context.Context, req *GetTrashDictionaryRequest, opts ...http.CallOption) (rsp *GetTrashDictionaryReply, err error)
	ListDictionary(ctx context.Context, req *ListDictionaryRequest, opts ...http.CallOption) (rsp *ListDictionaryReply, err error)
	ListTrashDictionary(ctx context.Context, req *ListTrashDictionaryRequest, opts ...http.CallOption) (rsp *ListTrashDictionaryReply, err error)
	RevertTrashDictionary(ctx context.Context, req *RevertTrashDictionaryRequest, opts ...http.CallOption) (rsp *RevertTrashDictionaryReply, err error)
	UpdateDictionary(ctx context.Context, req *UpdateDictionaryRequest, opts ...http.CallOption) (rsp *UpdateDictionaryReply, err error)
}

type UsercenterHTTPClientImpl struct {
	cc *http.Client
}

func NewUsercenterHTTPClient(client *http.Client) UsercenterHTTPClient {
	return &UsercenterHTTPClientImpl{client}
}

func (c *UsercenterHTTPClientImpl) CreateDictionary(ctx context.Context, in *CreateDictionaryRequest, opts ...http.CallOption) (*CreateDictionaryReply, error) {
	var out CreateDictionaryReply
	pattern := "/usercenter/api/v1/dictionary"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUsercenterCreateDictionary))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UsercenterHTTPClientImpl) DeleteDictionary(ctx context.Context, in *DeleteDictionaryRequest, opts ...http.CallOption) (*DeleteDictionaryReply, error) {
	var out DeleteDictionaryReply
	pattern := "/usercenter/api/v1/dictionary"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUsercenterDeleteDictionary))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UsercenterHTTPClientImpl) DeleteTrashDictionary(ctx context.Context, in *DeleteTrashDictionaryRequest, opts ...http.CallOption) (*DeleteTrashDictionaryReply, error) {
	var out DeleteTrashDictionaryReply
	pattern := "/usercenter/api/v1/dictionary/trash"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUsercenterDeleteTrashDictionary))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UsercenterHTTPClientImpl) GetDictionary(ctx context.Context, in *GetDictionaryRequest, opts ...http.CallOption) (*GetDictionaryReply, error) {
	var out GetDictionaryReply
	pattern := "/usercenter/api/v1/dictionary"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUsercenterGetDictionary))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UsercenterHTTPClientImpl) GetTrashDictionary(ctx context.Context, in *GetTrashDictionaryRequest, opts ...http.CallOption) (*GetTrashDictionaryReply, error) {
	var out GetTrashDictionaryReply
	pattern := "/usercenter/api/v1/dictionary/trash"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUsercenterGetTrashDictionary))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UsercenterHTTPClientImpl) ListDictionary(ctx context.Context, in *ListDictionaryRequest, opts ...http.CallOption) (*ListDictionaryReply, error) {
	var out ListDictionaryReply
	pattern := "/usercenter/api/v1/dictionaries"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUsercenterListDictionary))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UsercenterHTTPClientImpl) ListTrashDictionary(ctx context.Context, in *ListTrashDictionaryRequest, opts ...http.CallOption) (*ListTrashDictionaryReply, error) {
	var out ListTrashDictionaryReply
	pattern := "/usercenter/api/v1/dictionary/trashes"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUsercenterListTrashDictionary))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UsercenterHTTPClientImpl) RevertTrashDictionary(ctx context.Context, in *RevertTrashDictionaryRequest, opts ...http.CallOption) (*RevertTrashDictionaryReply, error) {
	var out RevertTrashDictionaryReply
	pattern := "/usercenter/api/v1/dictionary/trash"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUsercenterRevertTrashDictionary))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UsercenterHTTPClientImpl) UpdateDictionary(ctx context.Context, in *UpdateDictionaryRequest, opts ...http.CallOption) (*UpdateDictionaryReply, error) {
	var out UpdateDictionaryReply
	pattern := "/usercenter/api/v1/dictionary"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUsercenterUpdateDictionary))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
