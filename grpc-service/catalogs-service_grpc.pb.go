// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc_service

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

// CatalogsServiceClient is the client API for CatalogsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CatalogsServiceClient interface {
	CreateCatalog(ctx context.Context, in *CreateCatalogRequest, opts ...grpc.CallOption) (*CreateCatalogResponse, error)
	AddCatalogItem(ctx context.Context, in *AddCatalogItemRequest, opts ...grpc.CallOption) (*AddCatalogItemResponse, error)
	GetCatalogItems(ctx context.Context, in *GetCatalogItemsRequest, opts ...grpc.CallOption) (*GetCatalogItemsResponse, error)
}

type catalogsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCatalogsServiceClient(cc grpc.ClientConnInterface) CatalogsServiceClient {
	return &catalogsServiceClient{cc}
}

func (c *catalogsServiceClient) CreateCatalog(ctx context.Context, in *CreateCatalogRequest, opts ...grpc.CallOption) (*CreateCatalogResponse, error) {
	out := new(CreateCatalogResponse)
	err := c.cc.Invoke(ctx, "/CatalogsService/CreateCatalog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogsServiceClient) AddCatalogItem(ctx context.Context, in *AddCatalogItemRequest, opts ...grpc.CallOption) (*AddCatalogItemResponse, error) {
	out := new(AddCatalogItemResponse)
	err := c.cc.Invoke(ctx, "/CatalogsService/AddCatalogItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogsServiceClient) GetCatalogItems(ctx context.Context, in *GetCatalogItemsRequest, opts ...grpc.CallOption) (*GetCatalogItemsResponse, error) {
	out := new(GetCatalogItemsResponse)
	err := c.cc.Invoke(ctx, "/CatalogsService/GetCatalogItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatalogsServiceServer is the server API for CatalogsService service.
// All implementations must embed UnimplementedCatalogsServiceServer
// for forward compatibility
type CatalogsServiceServer interface {
	CreateCatalog(context.Context, *CreateCatalogRequest) (*CreateCatalogResponse, error)
	AddCatalogItem(context.Context, *AddCatalogItemRequest) (*AddCatalogItemResponse, error)
	GetCatalogItems(context.Context, *GetCatalogItemsRequest) (*GetCatalogItemsResponse, error)
	mustEmbedUnimplementedCatalogsServiceServer()
}

// UnimplementedCatalogsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCatalogsServiceServer struct {
}

func (UnimplementedCatalogsServiceServer) CreateCatalog(context.Context, *CreateCatalogRequest) (*CreateCatalogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCatalog not implemented")
}
func (UnimplementedCatalogsServiceServer) AddCatalogItem(context.Context, *AddCatalogItemRequest) (*AddCatalogItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCatalogItem not implemented")
}
func (UnimplementedCatalogsServiceServer) GetCatalogItems(context.Context, *GetCatalogItemsRequest) (*GetCatalogItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCatalogItems not implemented")
}
func (UnimplementedCatalogsServiceServer) mustEmbedUnimplementedCatalogsServiceServer() {}

// UnsafeCatalogsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CatalogsServiceServer will
// result in compilation errors.
type UnsafeCatalogsServiceServer interface {
	mustEmbedUnimplementedCatalogsServiceServer()
}

func RegisterCatalogsServiceServer(s grpc.ServiceRegistrar, srv CatalogsServiceServer) {
	s.RegisterService(&CatalogsService_ServiceDesc, srv)
}

func _CatalogsService_CreateCatalog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCatalogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogsServiceServer).CreateCatalog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CatalogsService/CreateCatalog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogsServiceServer).CreateCatalog(ctx, req.(*CreateCatalogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogsService_AddCatalogItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCatalogItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogsServiceServer).AddCatalogItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CatalogsService/AddCatalogItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogsServiceServer).AddCatalogItem(ctx, req.(*AddCatalogItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogsService_GetCatalogItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCatalogItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogsServiceServer).GetCatalogItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CatalogsService/GetCatalogItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogsServiceServer).GetCatalogItems(ctx, req.(*GetCatalogItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CatalogsService_ServiceDesc is the grpc.ServiceDesc for CatalogsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CatalogsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CatalogsService",
	HandlerType: (*CatalogsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCatalog",
			Handler:    _CatalogsService_CreateCatalog_Handler,
		},
		{
			MethodName: "AddCatalogItem",
			Handler:    _CatalogsService_AddCatalogItem_Handler,
		},
		{
			MethodName: "GetCatalogItems",
			Handler:    _CatalogsService_GetCatalogItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc-service/catalogs-service.proto",
}
