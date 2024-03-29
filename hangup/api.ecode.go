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

// HangUpErrCode ecode
var (
	SetLoginInfoFail         = ecode.New(30001)
	UserNotHangUp            = ecode.New(30002)
	PushOffOnline            = ecode.New(30003)
	OffOnline                = ecode.New(30004)
	GetLockFail              = ecode.New(30005)
	GetConfigFail            = ecode.New(30006)
	GetRewardFail            = ecode.New(30007)
	ClientGetHangupIPFail    = ecode.New(30008)
	HangupWorkOrNotGetReward = ecode.New(30009)
	RewardHaveGet            = ecode.New(30010)
	FinishHangUpFail         = ecode.New(30012)
)
