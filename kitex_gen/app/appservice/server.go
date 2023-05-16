// Code generated by Kitex v0.5.2. DO NOT EDIT.
package appservice

import (
	server "github.com/cloudwego/kitex/server"
	app "github.com/qml-123/app_log/kitex_gen/app"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler app.AppService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
