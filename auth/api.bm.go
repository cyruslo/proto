// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: api.proto

/*
Package api is a generated blademaster stub package.
This code was generated with kratos/tool/protobuf/protoc-gen-bm v0.1.

package 命名使用 {appid}.{version} 的方式, version 形如 v1, v2 ..

It is generated from these files:
	api.proto
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

var PathAuthGetToken = "/hys_auth.v1.Auth/GetToken"
var PathAuthUpdateToken = "/hys_auth.v1.Auth/UpdateToken"
var PathAuthVerifyToken = "/hys_auth.v1.Auth/VerifyToken"

// AuthBMServer is the server API for Auth service.
type AuthBMServer interface {
	GetToken(ctx context.Context, req *Userinfo) (resp *Token, err error)

	UpdateToken(ctx context.Context, req *google_protobuf1.Empty) (resp *Token, err error)

	VerifyToken(ctx context.Context, req *google_protobuf1.Empty) (resp *VerifyResult, err error)
}

var AuthSvc AuthBMServer

func authGetToken(c *bm.Context) {
	p := new(Userinfo)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := AuthSvc.GetToken(c, p)
	c.JSON(resp, err)
}

func authUpdateToken(c *bm.Context) {
	p := new(google_protobuf1.Empty)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := AuthSvc.UpdateToken(c, p)
	c.JSON(resp, err)
}

func authVerifyToken(c *bm.Context) {
	p := new(google_protobuf1.Empty)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := AuthSvc.VerifyToken(c, p)
	c.JSON(resp, err)
}

// RegisterAuthBMServer Register the blademaster route
func RegisterAuthBMServer(e *bm.Engine, server AuthBMServer) {
	AuthSvc = server
	e.GET("/hys_auth.v1.Auth/GetToken", authGetToken)
	e.GET("/hys_auth.v1.Auth/UpdateToken", authUpdateToken)
	e.GET("/hys_auth.v1.Auth/VerifyToken", authVerifyToken)
}
