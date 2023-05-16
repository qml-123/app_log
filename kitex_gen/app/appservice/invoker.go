// Code generated by Kitex v0.5.2. DO NOT EDIT.

package appservice

import (
	server "github.com/cloudwego/kitex/server"
	app "github.com/qml-123/app_log/kitex_gen/app"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler app.AppService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
