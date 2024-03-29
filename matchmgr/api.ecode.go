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
	Failed             = ecode.New(1)
	TimeOut            = ecode.New(2)
	InvalidParameter   = ecode.New(3)
	DeserializedFailed = ecode.New(4)
	InvalidToken       = ecode.New(5)
	NotExistPlayer     = ecode.New(6)
	NotMatchTime       = ecode.New(7)
	NoGameConf         = ecode.New(8)
)
