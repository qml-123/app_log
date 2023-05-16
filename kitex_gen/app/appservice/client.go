// Code generated by Kitex v0.5.2. DO NOT EDIT.

package appservice

import (
	"context"

	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	"github.com/qml-123/app_log/kitex_gen/app"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Ping(ctx context.Context, req *app.PingRequest, callOptions ...callopt.Option) (r *app.PingResponse, err error)
	GetFile(ctx context.Context, req *app.GetFileRequest, callOptions ...callopt.Option) (r *app.GetFileResponse, err error)
	Upload(ctx context.Context, req *app.UploadFileRequest, callOptions ...callopt.Option) (r *app.UploadFileResponse, err error)
	Register(ctx context.Context, req *app.RegisteRequest, callOptions ...callopt.Option) (r *app.RegisteResponse, err error)
	Login(ctx context.Context, req *app.LoginRequest, callOptions ...callopt.Option) (r *app.LoginResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kAppServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kAppServiceClient struct {
	*kClient
}

func (p *kAppServiceClient) Ping(ctx context.Context, req *app.PingRequest, callOptions ...callopt.Option) (r *app.PingResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Ping(ctx, req)
}

func (p *kAppServiceClient) GetFile(ctx context.Context, req *app.GetFileRequest, callOptions ...callopt.Option) (r *app.GetFileResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFile(ctx, req)
}

func (p *kAppServiceClient) Upload(ctx context.Context, req *app.UploadFileRequest, callOptions ...callopt.Option) (r *app.UploadFileResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Upload(ctx, req)
}

func (p *kAppServiceClient) Register(ctx context.Context, req *app.RegisteRequest, callOptions ...callopt.Option) (r *app.RegisteResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Register(ctx, req)
}

func (p *kAppServiceClient) Login(ctx context.Context, req *app.LoginRequest, callOptions ...callopt.Option) (r *app.LoginResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Login(ctx, req)
}
