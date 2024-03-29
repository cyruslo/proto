// Code generated by protoc-gen-comet v0.1, DO NOT EDIT.
// source: api.proto

/*
Package api is a generated comet stub package.
This code was generated with kratos/tool/protobuf/protoc-gen-comet v0.1.

package 命名使用 {appid}.{version} 的方式, version 形如 v1, v2 ..

It is generated from these files:
	api.proto
*/
package api

import (
	"context"

	"github.com/cyruslo/library/pkg/net/comet"
	"github.com/golang/protobuf/proto"
)

// HalldrawCometServer is the server API for Halldraw service.
type HalldrawCometServer interface {
	SetCometChan(cl *comet.ChanList, cs *comet.Server)
	Player_DrawList(ctx context.Context, req *Player_DrawListReq) (resp *Player_DrawListRsp, err error)

	Player_Draw(ctx context.Context, req *Player_DrawReq) (resp *Player_DrawRsp, err error)

	Player_Compose(ctx context.Context, req *Player_ComposeReq) (resp *Player_ComposeRsp, err error)

	Player_Refresh(ctx context.Context, req *Player_RefreshReq) (resp *Player_RefreshRsp, err error)

	Player_Share(ctx context.Context, req *Player_ShareReq) (resp *Player_ShareRsp, err error)

	Player_Exchange(ctx context.Context, req *Player_ExchangeReq) (resp *Player_ExchangeRsp, err error)

	Player_Record(ctx context.Context, req *Player_RecordReq) (resp *Player_RecordRsp, err error)

	Player_GetActBoxReward(ctx context.Context, req *Player_GetActBoxRewardReq) (resp *Player_GetActBoxRewardRsp, err error)

	Player_GetActBoxHistory(ctx context.Context, req *Player_GetActBoxHistoryReq) (resp *Player_GetActBoxHistoryRsp, err error)

	Player_GetActBoxRewardStatus(ctx context.Context, req *Player_GetActBoxRewardStatusReq) (resp *Player_GetActBoxRewardStatusRsp, err error)
}

func RegisterHalldrawCometServer(s *comet.Server, srv HalldrawCometServer) {
	chanList := s.RegisterService(&_Comet_Halldraw_serviceDesc, srv)
	srv.SetCometChan(chanList, s)
}

func _Comet_Halldraw_Player_DrawList_Handler(srv interface{}, ctx context.Context, data []byte, interceptor comet.UnaryServerInterceptor) ([]byte, error) {
	in := new(Player_DrawListReq)
	err := proto.Unmarshal(data, in)
	if err = proto.Unmarshal(data, in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		out, err := srv.(HalldrawCometServer).Player_DrawList(ctx, in)
		data, _ := proto.Marshal(out)
		return data, err
	}
	info := &comet.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/halldraw.service.v1.halldraw/Player_DrawList",
	}
	handler := func(ctx context.Context, req interface{}) ([]byte, error) {
		out, err := srv.(HalldrawCometServer).Player_DrawList(ctx, req.(*Player_DrawListReq))
		if out != nil {
			data, _ := proto.Marshal(out)
			return data, err
		}
		return nil, err
	}
	return interceptor(ctx, in, info, handler)
}
func _Comet_Halldraw_Player_Draw_Handler(srv interface{}, ctx context.Context, data []byte, interceptor comet.UnaryServerInterceptor) ([]byte, error) {
	in := new(Player_DrawReq)
	err := proto.Unmarshal(data, in)
	if err = proto.Unmarshal(data, in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		out, err := srv.(HalldrawCometServer).Player_Draw(ctx, in)
		data, _ := proto.Marshal(out)
		return data, err
	}
	info := &comet.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/halldraw.service.v1.halldraw/Player_Draw",
	}
	handler := func(ctx context.Context, req interface{}) ([]byte, error) {
		out, err := srv.(HalldrawCometServer).Player_Draw(ctx, req.(*Player_DrawReq))
		if out != nil {
			data, _ := proto.Marshal(out)
			return data, err
		}
		return nil, err
	}
	return interceptor(ctx, in, info, handler)
}
func _Comet_Halldraw_Player_Compose_Handler(srv interface{}, ctx context.Context, data []byte, interceptor comet.UnaryServerInterceptor) ([]byte, error) {
	in := new(Player_ComposeReq)
	err := proto.Unmarshal(data, in)
	if err = proto.Unmarshal(data, in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		out, err := srv.(HalldrawCometServer).Player_Compose(ctx, in)
		data, _ := proto.Marshal(out)
		return data, err
	}
	info := &comet.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/halldraw.service.v1.halldraw/Player_Compose",
	}
	handler := func(ctx context.Context, req interface{}) ([]byte, error) {
		out, err := srv.(HalldrawCometServer).Player_Compose(ctx, req.(*Player_ComposeReq))
		if out != nil {
			data, _ := proto.Marshal(out)
			return data, err
		}
		return nil, err
	}
	return interceptor(ctx, in, info, handler)
}
func _Comet_Halldraw_Player_Refresh_Handler(srv interface{}, ctx context.Context, data []byte, interceptor comet.UnaryServerInterceptor) ([]byte, error) {
	in := new(Player_RefreshReq)
	err := proto.Unmarshal(data, in)
	if err = proto.Unmarshal(data, in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		out, err := srv.(HalldrawCometServer).Player_Refresh(ctx, in)
		data, _ := proto.Marshal(out)
		return data, err
	}
	info := &comet.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/halldraw.service.v1.halldraw/Player_Refresh",
	}
	handler := func(ctx context.Context, req interface{}) ([]byte, error) {
		out, err := srv.(HalldrawCometServer).Player_Refresh(ctx, req.(*Player_RefreshReq))
		if out != nil {
			data, _ := proto.Marshal(out)
			return data, err
		}
		return nil, err
	}
	return interceptor(ctx, in, info, handler)
}
func _Comet_Halldraw_Player_Share_Handler(srv interface{}, ctx context.Context, data []byte, interceptor comet.UnaryServerInterceptor) ([]byte, error) {
	in := new(Player_ShareReq)
	err := proto.Unmarshal(data, in)
	if err = proto.Unmarshal(data, in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		out, err := srv.(HalldrawCometServer).Player_Share(ctx, in)
		data, _ := proto.Marshal(out)
		return data, err
	}
	info := &comet.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/halldraw.service.v1.halldraw/Player_Share",
	}
	handler := func(ctx context.Context, req interface{}) ([]byte, error) {
		out, err := srv.(HalldrawCometServer).Player_Share(ctx, req.(*Player_ShareReq))
		if out != nil {
			data, _ := proto.Marshal(out)
			return data, err
		}
		return nil, err
	}
	return interceptor(ctx, in, info, handler)
}
func _Comet_Halldraw_Player_Exchange_Handler(srv interface{}, ctx context.Context, data []byte, interceptor comet.UnaryServerInterceptor) ([]byte, error) {
	in := new(Player_ExchangeReq)
	err := proto.Unmarshal(data, in)
	if err = proto.Unmarshal(data, in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		out, err := srv.(HalldrawCometServer).Player_Exchange(ctx, in)
		data, _ := proto.Marshal(out)
		return data, err
	}
	info := &comet.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/halldraw.service.v1.halldraw/Player_Exchange",
	}
	handler := func(ctx context.Context, req interface{}) ([]byte, error) {
		out, err := srv.(HalldrawCometServer).Player_Exchange(ctx, req.(*Player_ExchangeReq))
		if out != nil {
			data, _ := proto.Marshal(out)
			return data, err
		}
		return nil, err
	}
	return interceptor(ctx, in, info, handler)
}
func _Comet_Halldraw_Player_Record_Handler(srv interface{}, ctx context.Context, data []byte, interceptor comet.UnaryServerInterceptor) ([]byte, error) {
	in := new(Player_RecordReq)
	err := proto.Unmarshal(data, in)
	if err = proto.Unmarshal(data, in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		out, err := srv.(HalldrawCometServer).Player_Record(ctx, in)
		data, _ := proto.Marshal(out)
		return data, err
	}
	info := &comet.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/halldraw.service.v1.halldraw/Player_Record",
	}
	handler := func(ctx context.Context, req interface{}) ([]byte, error) {
		out, err := srv.(HalldrawCometServer).Player_Record(ctx, req.(*Player_RecordReq))
		if out != nil {
			data, _ := proto.Marshal(out)
			return data, err
		}
		return nil, err
	}
	return interceptor(ctx, in, info, handler)
}
func _Comet_Halldraw_Player_GetActBoxReward_Handler(srv interface{}, ctx context.Context, data []byte, interceptor comet.UnaryServerInterceptor) ([]byte, error) {
	in := new(Player_GetActBoxRewardReq)
	err := proto.Unmarshal(data, in)
	if err = proto.Unmarshal(data, in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		out, err := srv.(HalldrawCometServer).Player_GetActBoxReward(ctx, in)
		data, _ := proto.Marshal(out)
		return data, err
	}
	info := &comet.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/halldraw.service.v1.halldraw/Player_GetActBoxReward",
	}
	handler := func(ctx context.Context, req interface{}) ([]byte, error) {
		out, err := srv.(HalldrawCometServer).Player_GetActBoxReward(ctx, req.(*Player_GetActBoxRewardReq))
		if out != nil {
			data, _ := proto.Marshal(out)
			return data, err
		}
		return nil, err
	}
	return interceptor(ctx, in, info, handler)
}
func _Comet_Halldraw_Player_GetActBoxHistory_Handler(srv interface{}, ctx context.Context, data []byte, interceptor comet.UnaryServerInterceptor) ([]byte, error) {
	in := new(Player_GetActBoxHistoryReq)
	err := proto.Unmarshal(data, in)
	if err = proto.Unmarshal(data, in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		out, err := srv.(HalldrawCometServer).Player_GetActBoxHistory(ctx, in)
		data, _ := proto.Marshal(out)
		return data, err
	}
	info := &comet.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/halldraw.service.v1.halldraw/Player_GetActBoxHistory",
	}
	handler := func(ctx context.Context, req interface{}) ([]byte, error) {
		out, err := srv.(HalldrawCometServer).Player_GetActBoxHistory(ctx, req.(*Player_GetActBoxHistoryReq))
		if out != nil {
			data, _ := proto.Marshal(out)
			return data, err
		}
		return nil, err
	}
	return interceptor(ctx, in, info, handler)
}
func _Comet_Halldraw_Player_GetActBoxRewardStatus_Handler(srv interface{}, ctx context.Context, data []byte, interceptor comet.UnaryServerInterceptor) ([]byte, error) {
	in := new(Player_GetActBoxRewardStatusReq)
	err := proto.Unmarshal(data, in)
	if err = proto.Unmarshal(data, in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		out, err := srv.(HalldrawCometServer).Player_GetActBoxRewardStatus(ctx, in)
		data, _ := proto.Marshal(out)
		return data, err
	}
	info := &comet.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/halldraw.service.v1.halldraw/Player_GetActBoxRewardStatus",
	}
	handler := func(ctx context.Context, req interface{}) ([]byte, error) {
		out, err := srv.(HalldrawCometServer).Player_GetActBoxRewardStatus(ctx, req.(*Player_GetActBoxRewardStatusReq))
		if out != nil {
			data, _ := proto.Marshal(out)
			return data, err
		}
		return nil, err
	}
	return interceptor(ctx, in, info, handler)
}

var _Comet_Halldraw_serviceDesc = comet.ServiceDesc{
	ServiceName: "halldraw.service.v1.halldraw",
	HandlerType: (*HalldrawCometServer)(nil),
	Methods: []comet.MethodDesc{
		{
			MethodName: "Player_DrawList",
			Handler:    _Comet_Halldraw_Player_DrawList_Handler,
			Ops:        4,
		},
		{
			MethodName: "Player_Draw",
			Handler:    _Comet_Halldraw_Player_Draw_Handler,
			Ops:        1,
		},
		{
			MethodName: "Player_Compose",
			Handler:    _Comet_Halldraw_Player_Compose_Handler,
			Ops:        2,
		},
		{
			MethodName: "Player_Refresh",
			Handler:    _Comet_Halldraw_Player_Refresh_Handler,
			Ops:        3,
		},
		{
			MethodName: "Player_Share",
			Handler:    _Comet_Halldraw_Player_Share_Handler,
			Ops:        5,
		},
		{
			MethodName: "Player_Exchange",
			Handler:    _Comet_Halldraw_Player_Exchange_Handler,
			Ops:        6,
		},
		{
			MethodName: "Player_Record",
			Handler:    _Comet_Halldraw_Player_Record_Handler,
			Ops:        7,
		},
		{
			MethodName: "Player_GetActBoxReward",
			Handler:    _Comet_Halldraw_Player_GetActBoxReward_Handler,
			Ops:        8,
		},
		{
			MethodName: "Player_GetActBoxHistory",
			Handler:    _Comet_Halldraw_Player_GetActBoxHistory_Handler,
			Ops:        9,
		},
		{
			MethodName: "Player_GetActBoxRewardStatus",
			Handler:    _Comet_Halldraw_Player_GetActBoxRewardStatus_Handler,
			Ops:        10,
		},
	},
}
