// Code generated by Kitex v0.5.2. DO NOT EDIT.

package appservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	app "github.com/qml-123/app_log/kitex_gen/app"
)

func serviceInfo() *kitex.ServiceInfo {
	return appServiceServiceInfo
}

var appServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "AppService"
	handlerType := (*app.AppService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Ping":       kitex.NewMethodInfo(pingHandler, newAppServicePingArgs, newAppServicePingResult, false),
		"GetFile":    kitex.NewMethodInfo(getFileHandler, newAppServiceGetFileArgs, newAppServiceGetFileResult, false),
		"Upload":     kitex.NewMethodInfo(uploadHandler, newAppServiceUploadArgs, newAppServiceUploadResult, false),
		"GetFileKey": kitex.NewMethodInfo(getFileKeyHandler, newAppServiceGetFileKeyArgs, newAppServiceGetFileKeyResult, false),
		"Register":   kitex.NewMethodInfo(registerHandler, newAppServiceRegisterArgs, newAppServiceRegisterResult, false),
		"Login":      kitex.NewMethodInfo(loginHandler, newAppServiceLoginArgs, newAppServiceLoginResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "app",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.5.2",
		Extra:           extra,
	}
	return svcInfo
}

func pingHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*app.AppServicePingArgs)
	realResult := result.(*app.AppServicePingResult)
	success, err := handler.(app.AppService).Ping(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAppServicePingArgs() interface{} {
	return app.NewAppServicePingArgs()
}

func newAppServicePingResult() interface{} {
	return app.NewAppServicePingResult()
}

func getFileHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*app.AppServiceGetFileArgs)
	realResult := result.(*app.AppServiceGetFileResult)
	success, err := handler.(app.AppService).GetFile(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAppServiceGetFileArgs() interface{} {
	return app.NewAppServiceGetFileArgs()
}

func newAppServiceGetFileResult() interface{} {
	return app.NewAppServiceGetFileResult()
}

func uploadHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*app.AppServiceUploadArgs)
	realResult := result.(*app.AppServiceUploadResult)
	success, err := handler.(app.AppService).Upload(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAppServiceUploadArgs() interface{} {
	return app.NewAppServiceUploadArgs()
}

func newAppServiceUploadResult() interface{} {
	return app.NewAppServiceUploadResult()
}

func getFileKeyHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*app.AppServiceGetFileKeyArgs)
	realResult := result.(*app.AppServiceGetFileKeyResult)
	success, err := handler.(app.AppService).GetFileKey(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAppServiceGetFileKeyArgs() interface{} {
	return app.NewAppServiceGetFileKeyArgs()
}

func newAppServiceGetFileKeyResult() interface{} {
	return app.NewAppServiceGetFileKeyResult()
}

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*app.AppServiceRegisterArgs)
	realResult := result.(*app.AppServiceRegisterResult)
	success, err := handler.(app.AppService).Register(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAppServiceRegisterArgs() interface{} {
	return app.NewAppServiceRegisterArgs()
}

func newAppServiceRegisterResult() interface{} {
	return app.NewAppServiceRegisterResult()
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*app.AppServiceLoginArgs)
	realResult := result.(*app.AppServiceLoginResult)
	success, err := handler.(app.AppService).Login(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAppServiceLoginArgs() interface{} {
	return app.NewAppServiceLoginArgs()
}

func newAppServiceLoginResult() interface{} {
	return app.NewAppServiceLoginResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Ping(ctx context.Context, req *app.PingRequest) (r *app.PingResponse, err error) {
	var _args app.AppServicePingArgs
	_args.Req = req
	var _result app.AppServicePingResult
	if err = p.c.Call(ctx, "Ping", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFile(ctx context.Context, req *app.GetFileRequest) (r *app.GetFileResponse, err error) {
	var _args app.AppServiceGetFileArgs
	_args.Req = req
	var _result app.AppServiceGetFileResult
	if err = p.c.Call(ctx, "GetFile", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Upload(ctx context.Context, req *app.UploadFileRequest) (r *app.UploadFileResponse, err error) {
	var _args app.AppServiceUploadArgs
	_args.Req = req
	var _result app.AppServiceUploadResult
	if err = p.c.Call(ctx, "Upload", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFileKey(ctx context.Context, req *app.GetFileRequest) (r *app.GetFileKeyResponse, err error) {
	var _args app.AppServiceGetFileKeyArgs
	_args.Req = req
	var _result app.AppServiceGetFileKeyResult
	if err = p.c.Call(ctx, "GetFileKey", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Register(ctx context.Context, req *app.RegisteRequest) (r *app.RegisteResponse, err error) {
	var _args app.AppServiceRegisterArgs
	_args.Req = req
	var _result app.AppServiceRegisterResult
	if err = p.c.Call(ctx, "Register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Login(ctx context.Context, req *app.LoginRequest) (r *app.LoginResponse, err error) {
	var _args app.AppServiceLoginArgs
	_args.Req = req
	var _result app.AppServiceLoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
