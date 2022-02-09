package api

import (
	"context"
	"fmt"

	"github.com/bilibili/kratos/pkg/net/rpc/warden"

	"google.golang.org/grpc"
)

// AppID .
const AppID = "hallactivity"

// NewClient new grpc client
func newClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (HallActivityClient, error) {
	client := warden.NewClient(cfg, opts...)
	cc, err := client.Dial(context.Background(), fmt.Sprintf("discovery://default/%s", AppID))
	if err != nil {
		return nil, err
	}
	return NewHallActivityClient(cc), nil
}

// 生成 gRPC 代码
//go:generate kratos tool protoc --service api.proto
