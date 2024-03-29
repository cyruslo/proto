// Code generated by protoc-gen-ecode v0.1, DO NOT EDIT.
// source: api.proto

/*
Package api is a generated ecode package.
This code was generated with kratos/tool/protobuf/protoc-gen-ecode v0.1.

package 命名使用 {appid}.{version} 的方式, version 形如 v1, v2 ..

It is generated from these files:
	api.proto
*/
package api

import (
	"github.com/bilibili/kratos/pkg/ecode"
)

// to suppressed 'imported but not used warning'
var _ ecode.Codes

// ErrCode ecode
var (
	Fragment_Not_Enough     = ecode.New(10000)
	Fragment_Update         = ecode.New(10001)
	Fragment_Deduct         = ecode.New(10002)
	Fragment_Compose        = ecode.New(10003)
	Draw_List_NotExist      = ecode.New(10004)
	Prop_Update             = ecode.New(10005)
	Coin_Deduct             = ecode.New(10006)
	Draw_InGame             = ecode.New(10007)
	Refresh_Times_NotEnough = ecode.New(10008)
	IsShared                = ecode.New(10009)
	DrawInterval            = ecode.New(10010)
	DrawParameter           = ecode.New(10011)
	ExchangeFail            = ecode.New(10012)
	ActiveIsColse           = ecode.New(10013)
	TokenFail               = ecode.New(10014)
	ActBox_OutTime          = ecode.New(10015)
	ActBox_Had_Got          = ecode.New(10016)
	Get_Box_Fail            = ecode.New(10017)
)
