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

// RewardPoolErrCode ecode
var (
	GetRewardPoolConfigFail = ecode.New(1000300)
	RewardPoolNumOverLimit  = ecode.New(1000301)
)