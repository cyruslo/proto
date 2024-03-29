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

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

var PathRankingListUpdateRankingList = "/RankingList/UpdateRankingList"
var PathRankingListGetRankingList = "/RankingList/GetRankingList"
var PathRankingListGetHallRankingList = "/RankingList/GetHallRankingList"

// RankingListBMServer is the server API for RankingList service.
type RankingListBMServer interface {
	UpdateRankingList(ctx context.Context, req *UpdateRankingListReq) (resp *UpdateRankingListRsp, err error)

	GetRankingList(ctx context.Context, req *GetRankingListReq) (resp *GetRankingListRsp, err error)

	GetHallRankingList(ctx context.Context, req *GetHallRankingListReq) (resp *GetHallRankingListRsp, err error)
}

var RankingListSvc RankingListBMServer

func rankingListUpdateRankingList(c *bm.Context) {
	p := new(UpdateRankingListReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := RankingListSvc.UpdateRankingList(c, p)
	c.JSON(resp, err)
}

func rankingListGetRankingList(c *bm.Context) {
	p := new(GetRankingListReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := RankingListSvc.GetRankingList(c, p)
	c.JSON(resp, err)
}

func rankingListGetHallRankingList(c *bm.Context) {
	p := new(GetHallRankingListReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := RankingListSvc.GetHallRankingList(c, p)
	c.JSON(resp, err)
}

// RegisterRankingListBMServer Register the blademaster route
func RegisterRankingListBMServer(e *bm.Engine, server RankingListBMServer) {
	RankingListSvc = server
	e.POST("/RankingList/UpdateRankingList", rankingListUpdateRankingList)
	e.POST("/RankingList/GetRankingList", rankingListGetRankingList)
	e.POST("/RankingList/GetHallRankingList", rankingListGetHallRankingList)
}
