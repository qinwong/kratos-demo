// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: kratos-demo/api/api.proto

/*
Package api is a generated blademaster stub package.
This code was generated with kratos/tool/protobuf/protoc-gen-bm v0.1.

package 命名使用 {appid}.{version} 的方式, version 形如 v1, v2 ..

It is generated from these files:
	kratos-demo/api/api.proto
*/
package api

import (
	"context"

	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/bilibili/kratos/pkg/net/http/blademaster/binding"
)
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

var PathDemoSayHello = "/demo.service.v1.Demo/SayHello"
var PathDemoSayHelloURL = "/kratos-demo/say_hello"

// DemoBMServer is the server API for Demo service.
type DemoBMServer interface {
	SayHello(ctx context.Context, req *HelloReq) (resp *google_protobuf1.Empty, err error)

	SayHelloURL(ctx context.Context, req *HelloReq) (resp *HelloResp, err error)
}

var DemoSvc DemoBMServer

func demoSayHello(c *bm.Context) {
	p := new(HelloReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := DemoSvc.SayHello(c, p)
	c.JSON(resp, err)
}

func demoSayHelloURL(c *bm.Context) {
	p := new(HelloReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := DemoSvc.SayHelloURL(c, p)
	c.JSON(resp, err)
}

// RegisterDemoBMServer Register the blademaster route
func RegisterDemoBMServer(e *bm.Engine, server DemoBMServer) {
	DemoSvc = server
	e.GET("/demo.service.v1.Demo/SayHello", demoSayHello)
	e.GET("/kratos-demo/say_hello", demoSayHelloURL)
}
