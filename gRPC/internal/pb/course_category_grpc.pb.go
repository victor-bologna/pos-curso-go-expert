// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: proto/course_category.proto

package pb

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

// CategoryServiceClient is the client API for CategoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CategoryServiceClient interface {
	CreateCategory(ctx context.Context, in *CategoryRequest, opts ...grpc.CallOption) (*Category, error)
	ListCategories(ctx context.Context, in *Blank, opts ...grpc.CallOption) (*ListCategoryResponse, error)
	GetCategory(ctx context.Context, in *CategoryGetRequest, opts ...grpc.CallOption) (*Category, error)
	StreamCategories(ctx context.Context, opts ...grpc.CallOption) (CategoryService_StreamCategoriesClient, error)
	BiStreamCategories(ctx context.Context, opts ...grpc.CallOption) (CategoryService_BiStreamCategoriesClient, error)
}

type categoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCategoryServiceClient(cc grpc.ClientConnInterface) CategoryServiceClient {
	return &categoryServiceClient{cc}
}

func (c *categoryServiceClient) CreateCategory(ctx context.Context, in *CategoryRequest, opts ...grpc.CallOption) (*Category, error) {
	out := new(Category)
	err := c.cc.Invoke(ctx, "/pb.CategoryService/CreateCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) ListCategories(ctx context.Context, in *Blank, opts ...grpc.CallOption) (*ListCategoryResponse, error) {
	out := new(ListCategoryResponse)
	err := c.cc.Invoke(ctx, "/pb.CategoryService/ListCategories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) GetCategory(ctx context.Context, in *CategoryGetRequest, opts ...grpc.CallOption) (*Category, error) {
	out := new(Category)
	err := c.cc.Invoke(ctx, "/pb.CategoryService/GetCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) StreamCategories(ctx context.Context, opts ...grpc.CallOption) (CategoryService_StreamCategoriesClient, error) {
	stream, err := c.cc.NewStream(ctx, &CategoryService_ServiceDesc.Streams[0], "/pb.CategoryService/StreamCategories", opts...)
	if err != nil {
		return nil, err
	}
	x := &categoryServiceStreamCategoriesClient{stream}
	return x, nil
}

type CategoryService_StreamCategoriesClient interface {
	Send(*CategoryRequest) error
	CloseAndRecv() (*ListCategoryResponse, error)
	grpc.ClientStream
}

type categoryServiceStreamCategoriesClient struct {
	grpc.ClientStream
}

func (x *categoryServiceStreamCategoriesClient) Send(m *CategoryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *categoryServiceStreamCategoriesClient) CloseAndRecv() (*ListCategoryResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ListCategoryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *categoryServiceClient) BiStreamCategories(ctx context.Context, opts ...grpc.CallOption) (CategoryService_BiStreamCategoriesClient, error) {
	stream, err := c.cc.NewStream(ctx, &CategoryService_ServiceDesc.Streams[1], "/pb.CategoryService/BiStreamCategories", opts...)
	if err != nil {
		return nil, err
	}
	x := &categoryServiceBiStreamCategoriesClient{stream}
	return x, nil
}

type CategoryService_BiStreamCategoriesClient interface {
	Send(*CategoryRequest) error
	Recv() (*Category, error)
	grpc.ClientStream
}

type categoryServiceBiStreamCategoriesClient struct {
	grpc.ClientStream
}

func (x *categoryServiceBiStreamCategoriesClient) Send(m *CategoryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *categoryServiceBiStreamCategoriesClient) Recv() (*Category, error) {
	m := new(Category)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CategoryServiceServer is the server API for CategoryService service.
// All implementations must embed UnimplementedCategoryServiceServer
// for forward compatibility
type CategoryServiceServer interface {
	CreateCategory(context.Context, *CategoryRequest) (*Category, error)
	ListCategories(context.Context, *Blank) (*ListCategoryResponse, error)
	GetCategory(context.Context, *CategoryGetRequest) (*Category, error)
	StreamCategories(CategoryService_StreamCategoriesServer) error
	BiStreamCategories(CategoryService_BiStreamCategoriesServer) error
	mustEmbedUnimplementedCategoryServiceServer()
}

// UnimplementedCategoryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCategoryServiceServer struct {
}

func (UnimplementedCategoryServiceServer) CreateCategory(context.Context, *CategoryRequest) (*Category, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCategory not implemented")
}
func (UnimplementedCategoryServiceServer) ListCategories(context.Context, *Blank) (*ListCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCategories not implemented")
}
func (UnimplementedCategoryServiceServer) GetCategory(context.Context, *CategoryGetRequest) (*Category, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategory not implemented")
}
func (UnimplementedCategoryServiceServer) StreamCategories(CategoryService_StreamCategoriesServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamCategories not implemented")
}
func (UnimplementedCategoryServiceServer) BiStreamCategories(CategoryService_BiStreamCategoriesServer) error {
	return status.Errorf(codes.Unimplemented, "method BiStreamCategories not implemented")
}
func (UnimplementedCategoryServiceServer) mustEmbedUnimplementedCategoryServiceServer() {}

// UnsafeCategoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CategoryServiceServer will
// result in compilation errors.
type UnsafeCategoryServiceServer interface {
	mustEmbedUnimplementedCategoryServiceServer()
}

func RegisterCategoryServiceServer(s grpc.ServiceRegistrar, srv CategoryServiceServer) {
	s.RegisterService(&CategoryService_ServiceDesc, srv)
}

func _CategoryService_CreateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).CreateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CategoryService/CreateCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).CreateCategory(ctx, req.(*CategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_ListCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Blank)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).ListCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CategoryService/ListCategories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).ListCategories(ctx, req.(*Blank))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_GetCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).GetCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CategoryService/GetCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).GetCategory(ctx, req.(*CategoryGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_StreamCategories_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CategoryServiceServer).StreamCategories(&categoryServiceStreamCategoriesServer{stream})
}

type CategoryService_StreamCategoriesServer interface {
	SendAndClose(*ListCategoryResponse) error
	Recv() (*CategoryRequest, error)
	grpc.ServerStream
}

type categoryServiceStreamCategoriesServer struct {
	grpc.ServerStream
}

func (x *categoryServiceStreamCategoriesServer) SendAndClose(m *ListCategoryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *categoryServiceStreamCategoriesServer) Recv() (*CategoryRequest, error) {
	m := new(CategoryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CategoryService_BiStreamCategories_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CategoryServiceServer).BiStreamCategories(&categoryServiceBiStreamCategoriesServer{stream})
}

type CategoryService_BiStreamCategoriesServer interface {
	Send(*Category) error
	Recv() (*CategoryRequest, error)
	grpc.ServerStream
}

type categoryServiceBiStreamCategoriesServer struct {
	grpc.ServerStream
}

func (x *categoryServiceBiStreamCategoriesServer) Send(m *Category) error {
	return x.ServerStream.SendMsg(m)
}

func (x *categoryServiceBiStreamCategoriesServer) Recv() (*CategoryRequest, error) {
	m := new(CategoryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CategoryService_ServiceDesc is the grpc.ServiceDesc for CategoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CategoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CategoryService",
	HandlerType: (*CategoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCategory",
			Handler:    _CategoryService_CreateCategory_Handler,
		},
		{
			MethodName: "ListCategories",
			Handler:    _CategoryService_ListCategories_Handler,
		},
		{
			MethodName: "GetCategory",
			Handler:    _CategoryService_GetCategory_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamCategories",
			Handler:       _CategoryService_StreamCategories_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BiStreamCategories",
			Handler:       _CategoryService_BiStreamCategories_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/course_category.proto",
}
