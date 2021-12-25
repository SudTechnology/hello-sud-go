package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"

	"github.com/gin-gonic/gin"

	"log"

	sdkAPI "github.com/SudTechnology/sud-mgp-auth-go/api"
)

const (
	appID     = "1461564080052506636"
	appSecret = "xJL0HU9ailVSGInqPyNK3Ev3qNHReRbR"
)
const (
	BaseRespSuccess    = 0
	BaseRespParamError = 1
)
const (
	BaseRespSuccessTips    = "成功"
	BaseRespParamErrorTips = "参数错误"
)

func NewBaseRespSuccess() *BaseResp {
	reps := &BaseResp{
		RetCode: BaseRespSuccess,
		RetMsg:  BaseRespSuccessTips,
	}
	return reps
}

func NewBaseRespParamError() *BaseResp {
	reps := &BaseResp{
		RetCode: BaseRespParamError,
		RetMsg:  BaseRespParamErrorTips,
	}
	return reps
}

func NewBaseRespError(sdkErrorCode int32) *BaseResp {
	reps := &BaseResp{
		RetCode:      BaseRespParamError,
		RetMsg:       BaseRespParamErrorTips,
		SdkErrorCode: sdkErrorCode,
	}
	return reps
}

/**
 * 登录接口，获取针对当前用户(UID)的短期令牌Code
 * 调用方：接入端APP
 */
func Login(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	req := &LoginReq{}
	json.Unmarshal(body, req)
	client := sdkAPI.NewSudMGPAuth(appID, appSecret)
	// 生成code和有效期（有效期默认30分钟）
	codeResp := client.GetCode(req.UserID, 0)
	data := &LoginResp{
		Code:       codeResp.Code,
		ExpireDate: codeResp.ExpireDate,
	}
	resp := NewBaseRespSuccess()
	resp.Data = data
	c.JSON(200, resp)
}

/**
 * 短期令牌Code更换长期令牌SSToken
 * 调用方：游戏服务
 */
func GetSsToken(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	req := &GetSSTokenReq{}
	json.Unmarshal(body, req)
	client := sdkAPI.NewSudMGPAuth(appID, appSecret)
	uidResp := client.GetUidByCode(req.Code)
	if !uidResp.IsSuccess {
		log.Printf("code err.code:%+v,uidResp:%+v \n", req.Code, uidResp)
		c.JSON(200, NewBaseRespError(uidResp.SdkErrorCode))
		return
	}

	// 生成token和有效期（有效期默认2小时）
	ssTokenResp := client.GetSSToken(uidResp.Uid, 0)
	data := &GetSSTokenResp{
		Token:      ssTokenResp.Token,
		ExpireDate: ssTokenResp.ExpireDate,
	}
	resp := NewBaseRespSuccess()
	resp.Data = data
	c.JSON(200, resp)
}

/**
 * 刷新长期令牌
 * 调用方：游戏服务
 */
func UpdateSSToken(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	req := &UpdateSSTokenReq{}
	json.Unmarshal(body, req)

	client := sdkAPI.NewSudMGPAuth(appID, appSecret)
	uidResp := client.GetUidBySSToken(req.Token)
	if !uidResp.IsSuccess {
		log.Printf("code err.code:%+v,uidResp:%+v \n", req.Token, uidResp)
		c.JSON(200, NewBaseRespError(uidResp.SdkErrorCode))
		return
	}

	// 生成token和有效期（有效期默认2小时）
	ssTokenResp := client.GetSSToken(uidResp.Uid, 0)
	data := &GetSSTokenResp{
		Token:      ssTokenResp.Token,
		ExpireDate: ssTokenResp.ExpireDate,
	}
	resp := NewBaseRespSuccess()
	resp.Data = data
	c.JSON(200, resp)
}

/**
 * 获取用户信息
 * 调用方：游戏服务
 */
func GetUserInfo(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	req := &GetUserInfoReq{}
	json.Unmarshal(body, req)

	client := sdkAPI.NewSudMGPAuth(appID, appSecret)
	uidResp := client.GetUidBySSToken(req.Token)
	if !uidResp.IsSuccess {
		log.Printf("code err.code:%+v,uidResp:%+v \n", req.Token, uidResp)
		c.JSON(200, NewBaseRespError(uidResp.SdkErrorCode))
		return
	}
	log.Printf("uidResp:%+v \n", uidResp)

	index := rand.Intn(100)
	uid := fmt.Sprintf("uid%d", index)
	nickName := fmt.Sprintf("name%d", index)
	gender := "male"
	avatar := "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fi2.hdslb.com%2Fbfs%2Fface%2F3d4f30235a8f3bb9914fe59ff58e1009e5498ba6.jpg%4068w_68h.jpg&refer=http%3A%2F%2Fi2.hdslb.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=jpeg?sec=1637226890&t=63d2ca921e1fde4abe30329f67e3612d"

	data := &GetUserInfoResp{
		Uid:      uid,
		NickName: nickName,
		Gender:   gender,
		Avatar:   avatar,
	}
	resp := NewBaseRespSuccess()
	resp.Data = data
	c.JSON(200, resp)
}

func ReportGameInfo(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Printf("ReportGameInfo body:%+v", string(body))

	data := &GetUserInfoResp{}
	resp := NewBaseRespSuccess()
	resp.Data = data
	c.JSON(200, resp)
}
