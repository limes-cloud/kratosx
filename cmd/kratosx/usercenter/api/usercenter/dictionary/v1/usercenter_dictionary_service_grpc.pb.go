// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: usercenter_dictionary_service.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Usercenter_GetDictionary_FullMethodName         = "/usercenter.api.usercenter.dictionary.v1.Usercenter/GetDictionary"
	Usercenter_ListDictionary_FullMethodName        = "/usercenter.api.usercenter.dictionary.v1.Usercenter/ListDictionary"
	Usercenter_CreateDictionary_FullMethodName      = "/usercenter.api.usercenter.dictionary.v1.Usercenter/CreateDictionary"
	Usercenter_UpdateDictionary_FullMethodName      = "/usercenter.api.usercenter.dictionary.v1.Usercenter/UpdateDictionary"
	Usercenter_DeleteDictionary_FullMethodName      = "/usercenter.api.usercenter.dictionary.v1.Usercenter/DeleteDictionary"
	Usercenter_GetTrashDictionary_FullMethodName    = "/usercenter.api.usercenter.dictionary.v1.Usercenter/GetTrashDictionary"
	Usercenter_ListTrashDictionary_FullMethodName   = "/usercenter.api.usercenter.dictionary.v1.Usercenter/ListTrashDictionary"
	Usercenter_DeleteTrashDictionary_FullMethodName = "/usercenter.api.usercenter.dictionary.v1.Usercenter/DeleteTrashDictionary"
	Usercenter_RevertTrashDictionary_FullMethodName = "/usercenter.api.usercenter.dictionary.v1.Usercenter/RevertTrashDictionary"
)

// UsercenterClient is the client API for Usercenter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsercenterClient interface {
	// GetDictionary 获取指定的字典信息
	GetDictionary(ctx context.Context, in *GetDictionaryRequest, opts ...grpc.CallOption) (*GetDictionaryReply, error)
	// ListDictionary 获取字典信息列表
	ListDictionary(ctx context.Context, in *ListDictionaryRequest, opts ...grpc.CallOption) (*ListDictionaryReply, error)
	// CreateDictionary 创建字典信息
	CreateDictionary(ctx context.Context, in *CreateDictionaryRequest, opts ...grpc.CallOption) (*CreateDictionaryReply, error)
	// UpdateDictionary 更新字典信息
	UpdateDictionary(ctx context.Context, in *UpdateDictionaryRequest, opts ...grpc.CallOption) (*UpdateDictionaryReply, error)
	// DeleteDictionary 删除字典信息
	DeleteDictionary(ctx context.Context, in *DeleteDictionaryRequest, opts ...grpc.CallOption) (*DeleteDictionaryReply, error)
	// GetTrashDictionary 查看指定字典信息回收站数据
	GetTrashDictionary(ctx context.Context, in *GetTrashDictionaryRequest, opts ...grpc.CallOption) (*GetTrashDictionaryReply, error)
	// ListTrashDictionary 查看字典信息列表回收站数据
	ListTrashDictionary(ctx context.Context, in *ListTrashDictionaryRequest, opts ...grpc.CallOption) (*ListTrashDictionaryReply, error)
	// DeleteTrashDictionary 彻底删除字典信息
	DeleteTrashDictionary(ctx context.Context, in *DeleteTrashDictionaryRequest, opts ...grpc.CallOption) (*DeleteTrashDictionaryReply, error)
	// RevertTrashDictionary 还原字典信息
	RevertTrashDictionary(ctx context.Context, in *RevertTrashDictionaryRequest, opts ...grpc.CallOption) (*RevertTrashDictionaryReply, error)
}

type usercenterClient struct {
	cc grpc.ClientConnInterface
}

func NewUsercenterClient(cc grpc.ClientConnInterface) UsercenterClient {
	return &usercenterClient{cc}
}

func (c *usercenterClient) GetDictionary(ctx context.Context, in *GetDictionaryRequest, opts ...grpc.CallOption) (*GetDictionaryReply, error) {
	out := new(GetDictionaryReply)
	err := c.cc.Invoke(ctx, Usercenter_GetDictionary_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) ListDictionary(ctx context.Context, in *ListDictionaryRequest, opts ...grpc.CallOption) (*ListDictionaryReply, error) {
	out := new(ListDictionaryReply)
	err := c.cc.Invoke(ctx, Usercenter_ListDictionary_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) CreateDictionary(ctx context.Context, in *CreateDictionaryRequest, opts ...grpc.CallOption) (*CreateDictionaryReply, error) {
	out := new(CreateDictionaryReply)
	err := c.cc.Invoke(ctx, Usercenter_CreateDictionary_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) UpdateDictionary(ctx context.Context, in *UpdateDictionaryRequest, opts ...grpc.CallOption) (*UpdateDictionaryReply, error) {
	out := new(UpdateDictionaryReply)
	err := c.cc.Invoke(ctx, Usercenter_UpdateDictionary_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) DeleteDictionary(ctx context.Context, in *DeleteDictionaryRequest, opts ...grpc.CallOption) (*DeleteDictionaryReply, error) {
	out := new(DeleteDictionaryReply)
	err := c.cc.Invoke(ctx, Usercenter_DeleteDictionary_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) GetTrashDictionary(ctx context.Context, in *GetTrashDictionaryRequest, opts ...grpc.CallOption) (*GetTrashDictionaryReply, error) {
	out := new(GetTrashDictionaryReply)
	err := c.cc.Invoke(ctx, Usercenter_GetTrashDictionary_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) ListTrashDictionary(ctx context.Context, in *ListTrashDictionaryRequest, opts ...grpc.CallOption) (*ListTrashDictionaryReply, error) {
	out := new(ListTrashDictionaryReply)
	err := c.cc.Invoke(ctx, Usercenter_ListTrashDictionary_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) DeleteTrashDictionary(ctx context.Context, in *DeleteTrashDictionaryRequest, opts ...grpc.CallOption) (*DeleteTrashDictionaryReply, error) {
	out := new(DeleteTrashDictionaryReply)
	err := c.cc.Invoke(ctx, Usercenter_DeleteTrashDictionary_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) RevertTrashDictionary(ctx context.Context, in *RevertTrashDictionaryRequest, opts ...grpc.CallOption) (*RevertTrashDictionaryReply, error) {
	out := new(RevertTrashDictionaryReply)
	err := c.cc.Invoke(ctx, Usercenter_RevertTrashDictionary_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsercenterServer is the server API for Usercenter service.
// All implementations must embed UnimplementedUsercenterServer
// for forward compatibility
type UsercenterServer interface {
	// GetDictionary 获取指定的字典信息
	GetDictionary(context.Context, *GetDictionaryRequest) (*GetDictionaryReply, error)
	// ListDictionary 获取字典信息列表
	ListDictionary(context.Context, *ListDictionaryRequest) (*ListDictionaryReply, error)
	// CreateDictionary 创建字典信息
	CreateDictionary(context.Context, *CreateDictionaryRequest) (*CreateDictionaryReply, error)
	// UpdateDictionary 更新字典信息
	UpdateDictionary(context.Context, *UpdateDictionaryRequest) (*UpdateDictionaryReply, error)
	// DeleteDictionary 删除字典信息
	DeleteDictionary(context.Context, *DeleteDictionaryRequest) (*DeleteDictionaryReply, error)
	// GetTrashDictionary 查看指定字典信息回收站数据
	GetTrashDictionary(context.Context, *GetTrashDictionaryRequest) (*GetTrashDictionaryReply, error)
	// ListTrashDictionary 查看字典信息列表回收站数据
	ListTrashDictionary(context.Context, *ListTrashDictionaryRequest) (*ListTrashDictionaryReply, error)
	// DeleteTrashDictionary 彻底删除字典信息
	DeleteTrashDictionary(context.Context, *DeleteTrashDictionaryRequest) (*DeleteTrashDictionaryReply, error)
	// RevertTrashDictionary 还原字典信息
	RevertTrashDictionary(context.Context, *RevertTrashDictionaryRequest) (*RevertTrashDictionaryReply, error)
	mustEmbedUnimplementedUsercenterServer()
}

// UnimplementedUsercenterServer must be embedded to have forward compatible implementations.
type UnimplementedUsercenterServer struct {
}

func (UnimplementedUsercenterServer) GetDictionary(context.Context, *GetDictionaryRequest) (*GetDictionaryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDictionary not implemented")
}
func (UnimplementedUsercenterServer) ListDictionary(context.Context, *ListDictionaryRequest) (*ListDictionaryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDictionary not implemented")
}
func (UnimplementedUsercenterServer) CreateDictionary(context.Context, *CreateDictionaryRequest) (*CreateDictionaryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDictionary not implemented")
}
func (UnimplementedUsercenterServer) UpdateDictionary(context.Context, *UpdateDictionaryRequest) (*UpdateDictionaryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDictionary not implemented")
}
func (UnimplementedUsercenterServer) DeleteDictionary(context.Context, *DeleteDictionaryRequest) (*DeleteDictionaryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDictionary not implemented")
}
func (UnimplementedUsercenterServer) GetTrashDictionary(context.Context, *GetTrashDictionaryRequest) (*GetTrashDictionaryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrashDictionary not implemented")
}
func (UnimplementedUsercenterServer) ListTrashDictionary(context.Context, *ListTrashDictionaryRequest) (*ListTrashDictionaryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTrashDictionary not implemented")
}
func (UnimplementedUsercenterServer) DeleteTrashDictionary(context.Context, *DeleteTrashDictionaryRequest) (*DeleteTrashDictionaryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTrashDictionary not implemented")
}
func (UnimplementedUsercenterServer) RevertTrashDictionary(context.Context, *RevertTrashDictionaryRequest) (*RevertTrashDictionaryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevertTrashDictionary not implemented")
}
func (UnimplementedUsercenterServer) mustEmbedUnimplementedUsercenterServer() {}

// UnsafeUsercenterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsercenterServer will
// result in compilation errors.
type UnsafeUsercenterServer interface {
	mustEmbedUnimplementedUsercenterServer()
}

func RegisterUsercenterServer(s grpc.ServiceRegistrar, srv UsercenterServer) {
	s.RegisterService(&Usercenter_ServiceDesc, srv)
}

func _Usercenter_GetDictionary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDictionaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).GetDictionary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_GetDictionary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).GetDictionary(ctx, req.(*GetDictionaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_ListDictionary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDictionaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).ListDictionary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_ListDictionary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).ListDictionary(ctx, req.(*ListDictionaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_CreateDictionary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDictionaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).CreateDictionary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_CreateDictionary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).CreateDictionary(ctx, req.(*CreateDictionaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_UpdateDictionary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDictionaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).UpdateDictionary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_UpdateDictionary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).UpdateDictionary(ctx, req.(*UpdateDictionaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_DeleteDictionary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDictionaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).DeleteDictionary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_DeleteDictionary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).DeleteDictionary(ctx, req.(*DeleteDictionaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_GetTrashDictionary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTrashDictionaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).GetTrashDictionary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_GetTrashDictionary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).GetTrashDictionary(ctx, req.(*GetTrashDictionaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_ListTrashDictionary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTrashDictionaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).ListTrashDictionary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_ListTrashDictionary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).ListTrashDictionary(ctx, req.(*ListTrashDictionaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_DeleteTrashDictionary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTrashDictionaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).DeleteTrashDictionary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_DeleteTrashDictionary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).DeleteTrashDictionary(ctx, req.(*DeleteTrashDictionaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_RevertTrashDictionary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevertTrashDictionaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).RevertTrashDictionary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_RevertTrashDictionary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).RevertTrashDictionary(ctx, req.(*RevertTrashDictionaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Usercenter_ServiceDesc is the grpc.ServiceDesc for Usercenter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Usercenter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "usercenter.api.usercenter.dictionary.v1.Usercenter",
	HandlerType: (*UsercenterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDictionary",
			Handler:    _Usercenter_GetDictionary_Handler,
		},
		{
			MethodName: "ListDictionary",
			Handler:    _Usercenter_ListDictionary_Handler,
		},
		{
			MethodName: "CreateDictionary",
			Handler:    _Usercenter_CreateDictionary_Handler,
		},
		{
			MethodName: "UpdateDictionary",
			Handler:    _Usercenter_UpdateDictionary_Handler,
		},
		{
			MethodName: "DeleteDictionary",
			Handler:    _Usercenter_DeleteDictionary_Handler,
		},
		{
			MethodName: "GetTrashDictionary",
			Handler:    _Usercenter_GetTrashDictionary_Handler,
		},
		{
			MethodName: "ListTrashDictionary",
			Handler:    _Usercenter_ListTrashDictionary_Handler,
		},
		{
			MethodName: "DeleteTrashDictionary",
			Handler:    _Usercenter_DeleteTrashDictionary_Handler,
		},
		{
			MethodName: "RevertTrashDictionary",
			Handler:    _Usercenter_RevertTrashDictionary_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "usercenter_dictionary_service.proto",
}
