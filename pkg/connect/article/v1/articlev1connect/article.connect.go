// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: article/v1/article.proto

package articlev1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/morning-night-dream/platform-app/pkg/connect/article/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// ArticleServiceName is the fully-qualified name of the ArticleService service.
	ArticleServiceName = "article.v1.ArticleService"
)

// ArticleServiceClient is a client for the article.v1.ArticleService service.
type ArticleServiceClient interface {
	// 共有
	// Need X-API-KEY Header
	Share(context.Context, *connect_go.Request[v1.ShareRequest]) (*connect_go.Response[v1.ShareResponse], error)
	// 一覧取得
	// Need Authorization Header
	List(context.Context, *connect_go.Request[v1.ListRequest]) (*connect_go.Response[v1.ListResponse], error)
	// 削除
	// Need Authorization Header
	Delete(context.Context, *connect_go.Request[v1.DeleteRequest]) (*connect_go.Response[v1.DeleteResponse], error)
	// 既読
	// Need Authorization Header
	Read(context.Context, *connect_go.Request[v1.ReadRequest]) (*connect_go.Response[v1.ReadResponse], error)
	// タグ追加
	AddTag(context.Context, *connect_go.Request[v1.AddTagRequest]) (*connect_go.Response[v1.AddTagResponse], error)
	// タグ一覧
	ListTag(context.Context, *connect_go.Request[v1.ListTagRequest]) (*connect_go.Response[v1.ListTagResponse], error)
}

// NewArticleServiceClient constructs a client for the article.v1.ArticleService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewArticleServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ArticleServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &articleServiceClient{
		share: connect_go.NewClient[v1.ShareRequest, v1.ShareResponse](
			httpClient,
			baseURL+"/article.v1.ArticleService/Share",
			opts...,
		),
		list: connect_go.NewClient[v1.ListRequest, v1.ListResponse](
			httpClient,
			baseURL+"/article.v1.ArticleService/List",
			opts...,
		),
		delete: connect_go.NewClient[v1.DeleteRequest, v1.DeleteResponse](
			httpClient,
			baseURL+"/article.v1.ArticleService/Delete",
			opts...,
		),
		read: connect_go.NewClient[v1.ReadRequest, v1.ReadResponse](
			httpClient,
			baseURL+"/article.v1.ArticleService/Read",
			opts...,
		),
		addTag: connect_go.NewClient[v1.AddTagRequest, v1.AddTagResponse](
			httpClient,
			baseURL+"/article.v1.ArticleService/AddTag",
			opts...,
		),
		listTag: connect_go.NewClient[v1.ListTagRequest, v1.ListTagResponse](
			httpClient,
			baseURL+"/article.v1.ArticleService/ListTag",
			opts...,
		),
	}
}

// articleServiceClient implements ArticleServiceClient.
type articleServiceClient struct {
	share   *connect_go.Client[v1.ShareRequest, v1.ShareResponse]
	list    *connect_go.Client[v1.ListRequest, v1.ListResponse]
	delete  *connect_go.Client[v1.DeleteRequest, v1.DeleteResponse]
	read    *connect_go.Client[v1.ReadRequest, v1.ReadResponse]
	addTag  *connect_go.Client[v1.AddTagRequest, v1.AddTagResponse]
	listTag *connect_go.Client[v1.ListTagRequest, v1.ListTagResponse]
}

// Share calls article.v1.ArticleService.Share.
func (c *articleServiceClient) Share(ctx context.Context, req *connect_go.Request[v1.ShareRequest]) (*connect_go.Response[v1.ShareResponse], error) {
	return c.share.CallUnary(ctx, req)
}

// List calls article.v1.ArticleService.List.
func (c *articleServiceClient) List(ctx context.Context, req *connect_go.Request[v1.ListRequest]) (*connect_go.Response[v1.ListResponse], error) {
	return c.list.CallUnary(ctx, req)
}

// Delete calls article.v1.ArticleService.Delete.
func (c *articleServiceClient) Delete(ctx context.Context, req *connect_go.Request[v1.DeleteRequest]) (*connect_go.Response[v1.DeleteResponse], error) {
	return c.delete.CallUnary(ctx, req)
}

// Read calls article.v1.ArticleService.Read.
func (c *articleServiceClient) Read(ctx context.Context, req *connect_go.Request[v1.ReadRequest]) (*connect_go.Response[v1.ReadResponse], error) {
	return c.read.CallUnary(ctx, req)
}

// AddTag calls article.v1.ArticleService.AddTag.
func (c *articleServiceClient) AddTag(ctx context.Context, req *connect_go.Request[v1.AddTagRequest]) (*connect_go.Response[v1.AddTagResponse], error) {
	return c.addTag.CallUnary(ctx, req)
}

// ListTag calls article.v1.ArticleService.ListTag.
func (c *articleServiceClient) ListTag(ctx context.Context, req *connect_go.Request[v1.ListTagRequest]) (*connect_go.Response[v1.ListTagResponse], error) {
	return c.listTag.CallUnary(ctx, req)
}

// ArticleServiceHandler is an implementation of the article.v1.ArticleService service.
type ArticleServiceHandler interface {
	// 共有
	// Need X-API-KEY Header
	Share(context.Context, *connect_go.Request[v1.ShareRequest]) (*connect_go.Response[v1.ShareResponse], error)
	// 一覧取得
	// Need Authorization Header
	List(context.Context, *connect_go.Request[v1.ListRequest]) (*connect_go.Response[v1.ListResponse], error)
	// 削除
	// Need Authorization Header
	Delete(context.Context, *connect_go.Request[v1.DeleteRequest]) (*connect_go.Response[v1.DeleteResponse], error)
	// 既読
	// Need Authorization Header
	Read(context.Context, *connect_go.Request[v1.ReadRequest]) (*connect_go.Response[v1.ReadResponse], error)
	// タグ追加
	AddTag(context.Context, *connect_go.Request[v1.AddTagRequest]) (*connect_go.Response[v1.AddTagResponse], error)
	// タグ一覧
	ListTag(context.Context, *connect_go.Request[v1.ListTagRequest]) (*connect_go.Response[v1.ListTagResponse], error)
}

// NewArticleServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewArticleServiceHandler(svc ArticleServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/article.v1.ArticleService/Share", connect_go.NewUnaryHandler(
		"/article.v1.ArticleService/Share",
		svc.Share,
		opts...,
	))
	mux.Handle("/article.v1.ArticleService/List", connect_go.NewUnaryHandler(
		"/article.v1.ArticleService/List",
		svc.List,
		opts...,
	))
	mux.Handle("/article.v1.ArticleService/Delete", connect_go.NewUnaryHandler(
		"/article.v1.ArticleService/Delete",
		svc.Delete,
		opts...,
	))
	mux.Handle("/article.v1.ArticleService/Read", connect_go.NewUnaryHandler(
		"/article.v1.ArticleService/Read",
		svc.Read,
		opts...,
	))
	mux.Handle("/article.v1.ArticleService/AddTag", connect_go.NewUnaryHandler(
		"/article.v1.ArticleService/AddTag",
		svc.AddTag,
		opts...,
	))
	mux.Handle("/article.v1.ArticleService/ListTag", connect_go.NewUnaryHandler(
		"/article.v1.ArticleService/ListTag",
		svc.ListTag,
		opts...,
	))
	return "/article.v1.ArticleService/", mux
}

// UnimplementedArticleServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedArticleServiceHandler struct{}

func (UnimplementedArticleServiceHandler) Share(context.Context, *connect_go.Request[v1.ShareRequest]) (*connect_go.Response[v1.ShareResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("article.v1.ArticleService.Share is not implemented"))
}

func (UnimplementedArticleServiceHandler) List(context.Context, *connect_go.Request[v1.ListRequest]) (*connect_go.Response[v1.ListResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("article.v1.ArticleService.List is not implemented"))
}

func (UnimplementedArticleServiceHandler) Delete(context.Context, *connect_go.Request[v1.DeleteRequest]) (*connect_go.Response[v1.DeleteResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("article.v1.ArticleService.Delete is not implemented"))
}

func (UnimplementedArticleServiceHandler) Read(context.Context, *connect_go.Request[v1.ReadRequest]) (*connect_go.Response[v1.ReadResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("article.v1.ArticleService.Read is not implemented"))
}

func (UnimplementedArticleServiceHandler) AddTag(context.Context, *connect_go.Request[v1.AddTagRequest]) (*connect_go.Response[v1.AddTagResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("article.v1.ArticleService.AddTag is not implemented"))
}

func (UnimplementedArticleServiceHandler) ListTag(context.Context, *connect_go.Request[v1.ListTagRequest]) (*connect_go.Response[v1.ListTagResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("article.v1.ArticleService.ListTag is not implemented"))
}