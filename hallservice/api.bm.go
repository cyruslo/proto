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

var PathHallserviceOnGetPassportID = "/hallservice/OnGetPassportID"
var PathHallserviceOnGetModifyFlag = "/hallservice/OnGetModifyFlag"
var PathHallserviceOnGetUserTicket = "/hallservice/OnGetUserTicket"
var PathHallserviceOnGetUserMoney = "/hallservice/OnGetUserMoney"
var PathHallserviceOnGetByPartnerUserid = "/hallservice/OnGetByPartnerUserid"
var PathHallserviceOnPartnerUserLogin = "/hallservice/OnPartnerUserLogin"
var PathHallserviceOnUserInfoUserEnterGame = "/hallservice/OnUserInfoUserEnterGame"
var PathHallserviceOnWechatSign = "/hallservice/OnWechatSign"
var PathHallserviceOnUserAccountLogin = "/hallservice/OnUserAccountLogin"
var PathHallserviceOnGetGameProps = "/hallservice/OnGetGameProps"
var PathHallserviceOnExchangeGameProps = "/hallservice/OnExchangeGameProps"
var PathHallserviceOnGetGameQuest = "/hallservice/OnGetGameQuest"
var PathHallserviceOnGetGameReward = "/hallservice/OnGetGameReward"
var PathHallserviceOnGetActReward = "/hallservice/OnGetActReward"
var PathHallserviceOnGetPayMoney = "/hallservice/OnGetPayMoney"
var PathHallserviceOnTimeToCall = "/hallservice/OnTimeToCall"
var PathHallserviceOnUserInfoUserExitGame = "/hallservice/OnUserInfoUserExitGame"
var PathHallserviceOnGameTransferStatAdd = "/hallservice/OnGameTransferStatAdd"
var PathHallserviceOnGetGamePropsForServer = "/hallservice/OnGetGamePropsForServer"
var PathHallserviceOnConsumeGameProps = "/hallservice/OnConsumeGameProps"
var PathHallserviceOnGetCardGame = "/hallservice/OnGetCardGame"
var PathHallserviceOnGetCardType = "/hallservice/OnGetCardType"
var PathHallserviceOnUpdateGameSignConfig = "/hallservice/OnUpdateGameSignConfig"
var PathHallserviceOnGetGameSignConfig = "/hallservice/OnGetGameSignConfig"
var PathHallserviceOnCommonUserSignLogAdd = "/hallservice/OnCommonUserSignLogAdd"
var PathHallserviceOnGetVipInfos = "/hallservice/OnGetVipInfos"
var PathHallserviceOnGetVipInfo = "/hallservice/OnGetVipInfo"
var PathHallserviceOnGetIsBind = "/hallservice/OnGetIsBind"
var PathHallserviceOnGetTableUserRecord = "/hallservice/OnGetTableUserRecord"
var PathHallserviceGetExtendCurrency = "/hallservice/GetExtendCurrency"
var PathHallserviceUpdateExtendCurrency = "/hallservice/UpdateExtendCurrency"
var PathHallserviceGetPetShareData = "/hallservice/GetPetShareData"
var PathHallserviceUpdatePetShareData = "/hallservice/UpdatePetShareData"
var PathHallserviceOnAwardUserTelephoneBill = "/hallservice/OnAwardUserTelephoneBill"
var PathHallserviceGetJumpInfo = "/hallservice/GetJumpInfo"
var PathHallserviceUpdateJumpInfo = "/hallservice/UpdateJumpInfo"
var PathHallserviceGetWingState = "/hallservice/GetWingState"
var PathHallserviceApplyGunSerialNumber = "/hallservice/ApplyGunSerialNumber"
var PathHallserviceGetGunSerialNumInfo = "/hallservice/GetGunSerialNumInfo"

// HallserviceBMServer is the server API for Hallservice service.
type HallserviceBMServer interface {
	OnGetPassportID(ctx context.Context, req *GetPassportIDReq) (resp *GetPassportIDRsp, err error)

	OnGetModifyFlag(ctx context.Context, req *GetModifyFlagReq) (resp *GetModifyFlagRsp, err error)

	OnGetUserTicket(ctx context.Context, req *GetUserTicketReq) (resp *GetUserTicketRsp, err error)

	OnGetUserMoney(ctx context.Context, req *GetUserMoneyReq) (resp *GetUserMoneyRsp, err error)

	OnGetByPartnerUserid(ctx context.Context, req *GetByPartnerUseridReq) (resp *GetByPartnerUseridRsp, err error)

	OnPartnerUserLogin(ctx context.Context, req *PartnerUserLoginReq) (resp *PartnerUserLoginRsp, err error)

	OnUserInfoUserEnterGame(ctx context.Context, req *UserInfoUserEnterGameReq) (resp *CommonResp, err error)

	OnWechatSign(ctx context.Context, req *WechatSignReq) (resp *WechatSignRsp, err error)

	OnUserAccountLogin(ctx context.Context, req *UserAccountLoginReq) (resp *UserAccountLoginRsp, err error)

	OnGetGameProps(ctx context.Context, req *GetGamePropsReq) (resp *GetGamePropsRsp, err error)

	OnExchangeGameProps(ctx context.Context, req *ExchangeGamePropsReq) (resp *CommonResp, err error)

	OnGetGameQuest(ctx context.Context, req *GetGameQuestReq) (resp *GetGameQuestRsp, err error)

	OnGetGameReward(ctx context.Context, req *GetGameRewardReq) (resp *GetGameRewardRsp, err error)

	OnGetActReward(ctx context.Context, req *GetActRewardReq) (resp *GetActRewardRsp, err error)

	OnGetPayMoney(ctx context.Context, req *GetPayMoneyReq) (resp *CommonResp, err error)

	OnTimeToCall(ctx context.Context, req *TimeToCallReq) (resp *TimeToCallRsp, err error)

	OnUserInfoUserExitGame(ctx context.Context, req *UserInfoUserExitGameReq) (resp *CommonResp, err error)

	OnGameTransferStatAdd(ctx context.Context, req *GameTransferStatAddReq) (resp *CommonResp, err error)

	OnGetGamePropsForServer(ctx context.Context, req *GetGamePropsForServerReq) (resp *CommonResp, err error)

	OnConsumeGameProps(ctx context.Context, req *ConsumeGamePropsReq) (resp *CommonResp, err error)

	OnGetCardGame(ctx context.Context, req *GetCardGameReq) (resp *CommonResp, err error)

	OnGetCardType(ctx context.Context, req *GetCardTypeReq) (resp *CommonResp, err error)

	// 让服务器重载GameSignConfig配置
	OnUpdateGameSignConfig(ctx context.Context, req *UpdateGameSignConfigReq) (resp *CommonResp, err error)

	// 获取GameSignConfig配置
	OnGetGameSignConfig(ctx context.Context, req *GetGameSignConfigReq) (resp *GetGameSignConfigRsp, err error)

	// 签到
	OnCommonUserSignLogAdd(ctx context.Context, req *CommonUserSignLogAddReq) (resp *CommonUserSignLogAddRsp, err error)

	// 获取Vip天数信息
	OnGetVipInfos(ctx context.Context, req *GetVipInfosReq) (resp *GetVipInfosRsp, err error)

	// 查询vip点数信息
	OnGetVipInfo(ctx context.Context, req *GetVipInfoReq) (resp *GetVipInfoRsp, err error)

	// 查询是否绑定(微信等合作商)
	OnGetIsBind(ctx context.Context, req *GetIsBindReq) (resp *GetIsBindRsp, err error)

	// 查询房间信息
	OnGetTableUserRecord(ctx context.Context, req *GetTableUserRecordReq) (resp *GetTableUserRecordRsp, err error)

	// 获取扩展货币信息
	GetExtendCurrency(ctx context.Context, req *GetExtendCurrencyReq) (resp *GetExtendCurrencyRsp, err error)

	// 更新扩展货币信息
	UpdateExtendCurrency(ctx context.Context, req *UpdateExtendCurrencyReq) (resp *UpdateExtendCurrencyRsp, err error)

	// 获取共享数据（宠物相关）
	GetPetShareData(ctx context.Context, req *GetPetShareDataReq) (resp *GetPetShareDataRsp, err error)

	// 更新共享数据
	UpdatePetShareData(ctx context.Context, req *UpdatePetShareDataReq) (resp *CommonResp, err error)

	OnAwardUserTelephoneBill(ctx context.Context, req *AwardUserTelephoneBillReq) (resp *CommonResp, err error)

	// 获取玩家的跳转信息
	GetJumpInfo(ctx context.Context, req *GetJumpInfoReq) (resp *GetJumpInfoRsp, err error)

	// 获取玩家的跳转信息
	UpdateJumpInfo(ctx context.Context, req *UpdateJumpInfoReq) (resp *UpdateJumpInfoRsp, err error)

	// 获取玩家是否拥有翅膀
	GetWingState(ctx context.Context, req *GetWingStateReq) (resp *GetWingStateRsp, err error)

	// 请求申请一个炮台的编号
	ApplyGunSerialNumber(ctx context.Context, req *ApplyGunSerialNumberReq) (resp *ApplyGunSerialNumberRsp, err error)

	// 玩家获取自己的炮台信息
	GetGunSerialNumInfo(ctx context.Context, req *GetGunSerialNumInfoReq) (resp *GetGunSerialNumInfoRsp, err error)
}

var HallserviceSvc HallserviceBMServer

func hallserviceOnGetPassportID(c *bm.Context) {
	p := new(GetPassportIDReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := HallserviceSvc.OnGetPassportID(c, p)
	c.JSON(resp, err)
}

func hallserviceOnGetModifyFlag(c *bm.Context) {
	p := new(GetModifyFlagReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := HallserviceSvc.OnGetModifyFlag(c, p)
	c.JSON(resp, err)
}

func hallserviceOnGetUserTicket(c *bm.Context) {
	p := new(GetUserTicketReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := HallserviceSvc.OnGetUserTicket(c, p)
	c.JSON(resp, err)
}

func hallserviceOnGetUserMoney(c *bm.Context) {
	p := new(GetUserMoneyReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := HallserviceSvc.OnGetUserMoney(c, p)
	c.JSON(resp, err)
}

func hallserviceOnGetByPartnerUserid(c *bm.Context) {
	p := new(GetByPartnerUseridReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := HallserviceSvc.OnGetByPartnerUserid(c, p)
	c.JSON(resp, err)
}

func hallserviceOnPartnerUserLogin(c *bm.Context) {
	p := new(PartnerUserLoginReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := HallserviceSvc.OnPartnerUserLogin(c, p)
	c.JSON(resp, err)
}

func hallserviceOnUserInfoUserEnterGame(c *bm.Context) {
	p := new(UserInfoUserEnterGameReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := HallserviceSvc.OnUserInfoUserEnterGame(c, p)
	c.JSON(resp, err)
}

func hallserviceOnWechatSign(c *bm.Context) {
	p := new(WechatSignReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := HallserviceSvc.OnWechatSign(c, p)
	c.JSON(resp, err)
}

func hallserviceOnUserAccountLogin(c *bm.Context) {
	p := new(UserAccountLoginReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := HallserviceSvc.OnUserAccountLogin(c, p)
	c.JSON(resp, err)
}

func hallserviceOnGetGameProps(c *bm.Context) {
	p := new(GetGamePropsReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := HallserviceSvc.OnGetGameProps(c, p)
	c.JSON(resp, err)
}

func hallserviceOnExchangeGameProps(c *bm.Context) {
	p := new(ExchangeGamePropsReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := HallserviceSvc.OnExchangeGameProps(c, p)
	c.JSON(resp, err)
}

func hallserviceOnGetGameQuest(c *bm.Context) {
	p := new(GetGameQuestReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := HallserviceSvc.OnGetGameQuest(c, p)
	c.JSON(resp, err)
}

func hallserviceOnGetGameReward(c *bm.Context) {
	p := new(GetGameRewardReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := HallserviceSvc.OnGetGameReward(c, p)
	c.JSON(resp, err)
}

func hallserviceOnGetActReward(c *bm.Context) {
	p := new(GetActRewardReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := HallserviceSvc.OnGetActReward(c, p)
	c.JSON(resp, err)
}

func hallserviceOnGetPayMoney(c *bm.Context) {
	p := new(GetPayMoneyReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := HallserviceSvc.OnGetPayMoney(c, p)
	c.JSON(resp, err)
}

func hallserviceOnTimeToCall(c *bm.Context) {
	p := new(TimeToCallReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := HallserviceSvc.OnTimeToCall(c, p)
	c.JSON(resp, err)
}

func hallserviceOnUserInfoUserExitGame(c *bm.Context) {
	p := new(UserInfoUserExitGameReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := HallserviceSvc.OnUserInfoUserExitGame(c, p)
	c.JSON(resp, err)
}

func hallserviceOnGameTransferStatAdd(c *bm.Context) {
	p := new(GameTransferStatAddReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := HallserviceSvc.OnGameTransferStatAdd(c, p)
	c.JSON(resp, err)
}

func hallserviceOnGetGamePropsForServer(c *bm.Context) {
	p := new(GetGamePropsForServerReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := HallserviceSvc.OnGetGamePropsForServer(c, p)
	c.JSON(resp, err)
}

func hallserviceOnConsumeGameProps(c *bm.Context) {
	p := new(ConsumeGamePropsReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := HallserviceSvc.OnConsumeGameProps(c, p)
	c.JSON(resp, err)
}

func hallserviceOnGetCardGame(c *bm.Context) {
	p := new(GetCardGameReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := HallserviceSvc.OnGetCardGame(c, p)
	c.JSON(resp, err)
}

func hallserviceOnGetCardType(c *bm.Context) {
	p := new(GetCardTypeReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := HallserviceSvc.OnGetCardType(c, p)
	c.JSON(resp, err)
}

func hallserviceOnUpdateGameSignConfig(c *bm.Context) {
	p := new(UpdateGameSignConfigReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := HallserviceSvc.OnUpdateGameSignConfig(c, p)
	c.JSON(resp, err)
}

func hallserviceOnGetGameSignConfig(c *bm.Context) {
	p := new(GetGameSignConfigReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := HallserviceSvc.OnGetGameSignConfig(c, p)
	c.JSON(resp, err)
}

func hallserviceOnCommonUserSignLogAdd(c *bm.Context) {
	p := new(CommonUserSignLogAddReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := HallserviceSvc.OnCommonUserSignLogAdd(c, p)
	c.JSON(resp, err)
}

func hallserviceOnGetVipInfos(c *bm.Context) {
	p := new(GetVipInfosReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := HallserviceSvc.OnGetVipInfos(c, p)
	c.JSON(resp, err)
}

func hallserviceOnGetVipInfo(c *bm.Context) {
	p := new(GetVipInfoReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := HallserviceSvc.OnGetVipInfo(c, p)
	c.JSON(resp, err)
}

func hallserviceOnGetIsBind(c *bm.Context) {
	p := new(GetIsBindReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := HallserviceSvc.OnGetIsBind(c, p)
	c.JSON(resp, err)
}

func hallserviceOnGetTableUserRecord(c *bm.Context) {
	p := new(GetTableUserRecordReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := HallserviceSvc.OnGetTableUserRecord(c, p)
	c.JSON(resp, err)
}

func hallserviceGetExtendCurrency(c *bm.Context) {
	p := new(GetExtendCurrencyReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := HallserviceSvc.GetExtendCurrency(c, p)
	c.JSON(resp, err)
}

func hallserviceUpdateExtendCurrency(c *bm.Context) {
	p := new(UpdateExtendCurrencyReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := HallserviceSvc.UpdateExtendCurrency(c, p)
	c.JSON(resp, err)
}

func hallserviceGetPetShareData(c *bm.Context) {
	p := new(GetPetShareDataReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := HallserviceSvc.GetPetShareData(c, p)
	c.JSON(resp, err)
}

func hallserviceUpdatePetShareData(c *bm.Context) {
	p := new(UpdatePetShareDataReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := HallserviceSvc.UpdatePetShareData(c, p)
	c.JSON(resp, err)
}

func hallserviceOnAwardUserTelephoneBill(c *bm.Context) {
	p := new(AwardUserTelephoneBillReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := HallserviceSvc.OnAwardUserTelephoneBill(c, p)
	c.JSON(resp, err)
}

func hallserviceGetJumpInfo(c *bm.Context) {
	p := new(GetJumpInfoReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := HallserviceSvc.GetJumpInfo(c, p)
	c.JSON(resp, err)
}

func hallserviceUpdateJumpInfo(c *bm.Context) {
	p := new(UpdateJumpInfoReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := HallserviceSvc.UpdateJumpInfo(c, p)
	c.JSON(resp, err)
}

func hallserviceGetWingState(c *bm.Context) {
	p := new(GetWingStateReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := HallserviceSvc.GetWingState(c, p)
	c.JSON(resp, err)
}

func hallserviceApplyGunSerialNumber(c *bm.Context) {
	p := new(ApplyGunSerialNumberReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := HallserviceSvc.ApplyGunSerialNumber(c, p)
	c.JSON(resp, err)
}

func hallserviceGetGunSerialNumInfo(c *bm.Context) {
	p := new(GetGunSerialNumInfoReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := HallserviceSvc.GetGunSerialNumInfo(c, p)
	c.JSON(resp, err)
}

// RegisterHallserviceBMServer Register the blademaster route
func RegisterHallserviceBMServer(e *bm.Engine, server HallserviceBMServer) {
	HallserviceSvc = server
	e.POST("/hallservice/OnGetPassportID", hallserviceOnGetPassportID)
	e.POST("/hallservice/OnGetModifyFlag", hallserviceOnGetModifyFlag)
	e.POST("/hallservice/OnGetUserTicket", hallserviceOnGetUserTicket)
	e.POST("/hallservice/OnGetUserMoney", hallserviceOnGetUserMoney)
	e.POST("/hallservice/OnGetByPartnerUserid", hallserviceOnGetByPartnerUserid)
	e.POST("/hallservice/OnPartnerUserLogin", hallserviceOnPartnerUserLogin)
	e.POST("/hallservice/OnUserInfoUserEnterGame", hallserviceOnUserInfoUserEnterGame)
	e.POST("/hallservice/OnWechatSign", hallserviceOnWechatSign)
	e.POST("/hallservice/OnUserAccountLogin", hallserviceOnUserAccountLogin)
	e.POST("/hallservice/OnGetGameProps", hallserviceOnGetGameProps)
	e.POST("/hallservice/OnExchangeGameProps", hallserviceOnExchangeGameProps)
	e.POST("/hallservice/OnGetGameQuest", hallserviceOnGetGameQuest)
	e.POST("/hallservice/OnGetGameReward", hallserviceOnGetGameReward)
	e.POST("/hallservice/OnGetActReward", hallserviceOnGetActReward)
	e.POST("/hallservice/OnGetPayMoney", hallserviceOnGetPayMoney)
	e.POST("/hallservice/OnTimeToCall", hallserviceOnTimeToCall)
	e.POST("/hallservice/OnUserInfoUserExitGame", hallserviceOnUserInfoUserExitGame)
	e.POST("/hallservice/OnGameTransferStatAdd", hallserviceOnGameTransferStatAdd)
	e.POST("/hallservice/OnGetGamePropsForServer", hallserviceOnGetGamePropsForServer)
	e.POST("/hallservice/OnConsumeGameProps", hallserviceOnConsumeGameProps)
	e.POST("/hallservice/OnGetCardGame", hallserviceOnGetCardGame)
	e.POST("/hallservice/OnGetCardType", hallserviceOnGetCardType)
	e.POST("/hallservice/OnUpdateGameSignConfig", hallserviceOnUpdateGameSignConfig)
	e.POST("/hallservice/OnGetGameSignConfig", hallserviceOnGetGameSignConfig)
	e.POST("/hallservice/OnCommonUserSignLogAdd", hallserviceOnCommonUserSignLogAdd)
	e.POST("/hallservice/OnGetVipInfos", hallserviceOnGetVipInfos)
	e.POST("/hallservice/OnGetVipInfo", hallserviceOnGetVipInfo)
	e.POST("/hallservice/OnGetIsBind", hallserviceOnGetIsBind)
	e.POST("/hallservice/OnGetTableUserRecord", hallserviceOnGetTableUserRecord)
	e.POST("/hallservice/GetExtendCurrency", hallserviceGetExtendCurrency)
	e.POST("/hallservice/UpdateExtendCurrency", hallserviceUpdateExtendCurrency)
	e.POST("/hallservice/GetPetShareData", hallserviceGetPetShareData)
	e.POST("/hallservice/UpdatePetShareData", hallserviceUpdatePetShareData)
	e.POST("/hallservice/OnAwardUserTelephoneBill", hallserviceOnAwardUserTelephoneBill)
	e.POST("/hallservice/GetJumpInfo", hallserviceGetJumpInfo)
	e.POST("/hallservice/UpdateJumpInfo", hallserviceUpdateJumpInfo)
	e.POST("/hallservice/GetWingState", hallserviceGetWingState)
	e.POST("/hallservice/ApplyGunSerialNumber", hallserviceApplyGunSerialNumber)
	e.POST("/hallservice/GetGunSerialNumInfo", hallserviceGetGunSerialNumInfo)
}
