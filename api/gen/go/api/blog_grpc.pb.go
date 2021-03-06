// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: api/blog.proto

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

// BlogServiceClient is the client API for BlogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlogServiceClient interface {
	GetArticle(ctx context.Context, in *GetArticleRequest, opts ...grpc.CallOption) (*Article, error)
	ListArticle(ctx context.Context, in *ListArticleRequest, opts ...grpc.CallOption) (*ListArticleResponse, error)
}

type blogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlogServiceClient(cc grpc.ClientConnInterface) BlogServiceClient {
	return &blogServiceClient{cc}
}

func (c *blogServiceClient) GetArticle(ctx context.Context, in *GetArticleRequest, opts ...grpc.CallOption) (*Article, error) {
	out := new(Article)
	err := c.cc.Invoke(ctx, "/blog.v1.BlogService/GetArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) ListArticle(ctx context.Context, in *ListArticleRequest, opts ...grpc.CallOption) (*ListArticleResponse, error) {
	out := new(ListArticleResponse)
	err := c.cc.Invoke(ctx, "/blog.v1.BlogService/ListArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlogServiceServer is the server API for BlogService service.
// All implementations should embed UnimplementedBlogServiceServer
// for forward compatibility
type BlogServiceServer interface {
	GetArticle(context.Context, *GetArticleRequest) (*Article, error)
	ListArticle(context.Context, *ListArticleRequest) (*ListArticleResponse, error)
}

// UnimplementedBlogServiceServer should be embedded to have forward compatible implementations.
type UnimplementedBlogServiceServer struct {
}

func (UnimplementedBlogServiceServer) GetArticle(context.Context, *GetArticleRequest) (*Article, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticle not implemented")
}
func (UnimplementedBlogServiceServer) ListArticle(context.Context, *ListArticleRequest) (*ListArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArticle not implemented")
}

// UnsafeBlogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlogServiceServer will
// result in compilation errors.
type UnsafeBlogServiceServer interface {
	mustEmbedUnimplementedBlogServiceServer()
}

func RegisterBlogServiceServer(s grpc.ServiceRegistrar, srv BlogServiceServer) {
	s.RegisterService(&BlogService_ServiceDesc, srv)
}

func _BlogService_GetArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).GetArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.v1.BlogService/GetArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).GetArticle(ctx, req.(*GetArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_ListArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).ListArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.v1.BlogService/ListArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).ListArticle(ctx, req.(*ListArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BlogService_ServiceDesc is the grpc.ServiceDesc for BlogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blog.v1.BlogService",
	HandlerType: (*BlogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetArticle",
			Handler:    _BlogService_GetArticle_Handler,
		},
		{
			MethodName: "ListArticle",
			Handler:    _BlogService_ListArticle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/blog.proto",
}
