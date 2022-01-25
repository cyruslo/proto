package api

import (
	"context"
	"fmt"
	"github.com/bilibili/kratos/pkg/log"
	xtime "github.com/bilibili/kratos/pkg/time"
	"sync"
	"time"

	"github.com/bilibili/kratos/pkg/net/rpc/warden"

	"google.golang.org/grpc"
)

// AppID .
const AppID = "hallservice"

// default client
var (
	_cli  HallserviceClient
	_once sync.Once
)

// NewClient new grpc client
func NewClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (HallserviceClient, error) {
	client := warden.NewClient(cfg, opts...)
	cc, err := client.Dial(context.Background(), fmt.Sprintf("discovery://default/%s", AppID))
	if err != nil {
		return nil, err
	}
	return NewHallserviceClient(cc), nil
}

// DefaultClient 默认客户端初始化
func DefaultClient() HallserviceClient {
	_once.Do(func() {
		cfg := &warden.ClientConfig{
			Timeout: xtime.Duration(time.Millisecond * 4000),
		}

		var err error
		_cli, err = NewClient(cfg)
		if err != nil {
			log.Error("client Dial error:err%s", err)
			return
		}

		log.Info("init %s client ok .", AppID)

	})

	return _cli
}

// 生成 gRPC 代码
//go:generate kratos tool protoc --service api.proto
