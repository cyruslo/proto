/*
 * @Author: your name
 * @Date: 2021-05-19 14:29:39
 * @LastEditTime: 2021-05-27 14:21:21
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \halldraw\api\client.go
 */
package api

import (
	"context"
	"fmt"

	"github.com/bilibili/kratos/pkg/net/rpc/warden"

	"google.golang.org/grpc"
)

// AppID .
const AppID = "halldraw"

// NewClient new grpc client
func NewClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (HalldrawClient, error) {
	client := warden.NewClient(cfg, opts...)
	cc, err := client.Dial(context.Background(), fmt.Sprintf("discovery://default/%s", AppID))
	if err != nil {
		return nil, err
	}
	return NewHalldrawClient(cc), nil
}

// 生成 gRPC 代码
//go:generate kratos tool protoc --service api.proto
