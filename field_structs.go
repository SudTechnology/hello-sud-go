package main

// BaseResp 结构体
type BaseResp struct {
	RetCode      int32       `json:"ret_code"`
	RetMsg       string      `json:"ret_msg"`
	SdkErrorCode int32       `json:"sdk_error_code"`
	Data         interface{} `json:"data"`
}

// LoginReq 请求结构体
type LoginReq struct {
	UserID string `json:"user_id"` //code id
}

// LoginResp 结构体
type LoginResp struct {
	Code       string `json:"code"`
	ExpireDate int64  `json:"expire_date"`
}

// GetSSTokenReq 请求结构体
type GetSSTokenReq struct {
	Code string `json:"code"` //code id
}

// GetSSTokenResp 结构体
type GetSSTokenResp struct {
	Token      string `json:"ss_token"`
	ExpireDate int64  `json:"expire_date"`
}

// UpdateSSTokenReq 请求结构体
type UpdateSSTokenReq struct {
	Token string `json:"ss_token"`
}

// UpdateSSTokenResp 结构体
type UpdateSSTokenResp struct {
	Token      string `json:"ss_token"`
	ExpireDate int64  `json:"expire_date"`
}

// GetUserInfoReq 请求结构体
type GetUserInfoReq struct {
	Token string `json:"ss_token"`
}

// GetUserInfoResp 结构体
type GetUserInfoResp struct {
	Uid      string `json:"uid"`
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar_url"`
	Gender   string `json:"gender"`
}
