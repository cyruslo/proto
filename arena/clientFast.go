/*
 * @Author: your name
 * @Date: 2019-12-08 13:33:45
 * @LastEditTime: 2019-12-08 13:40:38
 * @LastEditors: your name
 * @Description: In User Settings Edit
 * @FilePath: \arenaServer\internal\api\client.go
 */
package api

import (
	"context"
	"fmt"

	"github.com/bilibili/kratos/pkg/net/rpc/warden"

	"google.golang.org/grpc"
)

// AppID .
const AppIDFast = "arenafast"

// NewClient new grpc client
func NewFastClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (ArenaServerClient, error) {
	client := warden.NewClient(cfg, opts...)
	cc, err := client.Dial(context.Background(), fmt.Sprintf("discovery://default/%s", AppIDFast))
	if err != nil {
		return nil, err
	}
	return NewArenaServerClient(cc), nil
}


