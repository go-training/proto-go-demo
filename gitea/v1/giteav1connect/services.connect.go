// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: gitea/v1/services.proto

package giteav1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/go-training/proto-go-demo/gitea/v1"
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
	// GiteaServiceName is the fully-qualified name of the GiteaService service.
	GiteaServiceName = "gitea.v1.GiteaService"
)

// GiteaServiceClient is a client for the gitea.v1.GiteaService service.
type GiteaServiceClient interface {
	Gitea(context.Context, *connect_go.Request[v1.GiteaRequest]) (*connect_go.Response[v1.GiteaResponse], error)
}

// NewGiteaServiceClient constructs a client for the gitea.v1.GiteaService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewGiteaServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) GiteaServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &giteaServiceClient{
		gitea: connect_go.NewClient[v1.GiteaRequest, v1.GiteaResponse](
			httpClient,
			baseURL+"/gitea.v1.GiteaService/Gitea",
			opts...,
		),
	}
}

// giteaServiceClient implements GiteaServiceClient.
type giteaServiceClient struct {
	gitea *connect_go.Client[v1.GiteaRequest, v1.GiteaResponse]
}

// Gitea calls gitea.v1.GiteaService.Gitea.
func (c *giteaServiceClient) Gitea(ctx context.Context, req *connect_go.Request[v1.GiteaRequest]) (*connect_go.Response[v1.GiteaResponse], error) {
	return c.gitea.CallUnary(ctx, req)
}

// GiteaServiceHandler is an implementation of the gitea.v1.GiteaService service.
type GiteaServiceHandler interface {
	Gitea(context.Context, *connect_go.Request[v1.GiteaRequest]) (*connect_go.Response[v1.GiteaResponse], error)
}

// NewGiteaServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewGiteaServiceHandler(svc GiteaServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/gitea.v1.GiteaService/Gitea", connect_go.NewUnaryHandler(
		"/gitea.v1.GiteaService/Gitea",
		svc.Gitea,
		opts...,
	))
	return "/gitea.v1.GiteaService/", mux
}

// UnimplementedGiteaServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedGiteaServiceHandler struct{}

func (UnimplementedGiteaServiceHandler) Gitea(context.Context, *connect_go.Request[v1.GiteaRequest]) (*connect_go.Response[v1.GiteaResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitea.v1.GiteaService.Gitea is not implemented"))
}
