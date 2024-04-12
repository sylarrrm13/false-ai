package req

import (
	"utils"
)

type CreateAdminReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateAdminReq struct {
	Username string `json:"username" binding:"required"`
	OriPass  string `json:"oriPass" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GetUserInfoList struct {
	PageNum  int `json:"pageNum" binding:"required"`
	PageSize int `json:"pageSize" binding:"required"`
}

type GetUserInfoListResp struct {
	Id           int              `json:"id" validate:"required"`
	Username     string           `json:"username" validate:"required"`
	Status       int              `json:"status"`
	Email        string           `json:"email" validate:"required"`
	Phone        string           `json:"phone" validate:"number"`
	Coins        int              `json:"coins"`
	UsedCoins    int              `json:"used_coins"`
	CoinsExpired *utils.LocalTime `json:"coins_expired"`
	Expired      *utils.LocalTime `json:"expired"`
}

type GetModelCateList struct {
	Id         int    `json:"id"`
	Sort       int    `json:"sort"`
	Name       string `json:"name"`
	Visible    int    `json:"visible"`
	Model      string `json:"model"`
	MaxToken   int    `json:"max_token"`
	Upload     int    `json:"upload"`
	UploadType string `json:"upload_type"`
	Net        int    `json:"net"`
	Tool       int    `json:"tool"`
}
type CodelkupWithoutRef struct {
	Listname string `json:"listname"`
	Value    string `json:"value"`
}

type CodelkupWithRef struct {
	Listname string `json:"listname"`
	Value    string `json:"value"`
	Ref      string `json:"ref"`
	Ref2     string `json:"ref2"`
}

type UpdateEmailConfigReq struct {
	EmailServer string `json:"emailServer" validate:"required"`
	EmailPort   string `json:"emailPort" validate:"number"`
	EmailUser   string `json:"emailUser" validate:"required"`
	EmailPass   string `json:"emailPass" validate:"required"`
	EmailTLS    string `json:"emailTLS" validate:"required"`
}
type UpdateEmailRegReq struct {
	EmailValidation string `json:"emailValidation" validate:"required"`
	EmailFormat     string `json:"emailFormat" validate:"required"`
	EmailSubject    string `json:"emailSubject" validate:"required"`
	EmailFrom       string `json:"emailFrom" validate:"required"`
	Duration        string `json:"duration" validate:"required,number"`
}

type wxLogin struct {
	Login     string `json:"login" validate:"required,oneof=0 1"` // 0:关闭 1:开启
	Auto      string `json:"auto" validate:"required,oneof=0 1"`
	Name      string `json:"name" validate:"required"`
	AppID     string `json:"appID" validate:"required"`
	Token     string `json:"token" validate:"required"`
	AppSecret string `json:"appSecret" validate:"required"`
	Ref1      string `json:"ref1" validate:"required"`
	Ref2      string `json:"ref2" validate:"required"`
	Ref3      string `json:"ref3" validate:"required"`
	Ref4      string `json:"ref4" validate:"required"`
}
type SendEmailCode struct {
	Email string `json:"email" validate:"required,email"`
}
type Notice struct {
	Notice string `json:"notice" validate:"required"`
}

type ModelKeyCateList struct {
	ModelCateID []int    `json:"model_cate_id"`
	Key         []string `json:"key"`
	ApiAddr     string   `json:"api_addr"`
	Weight      int      `json:"weight"`
	Enable      int      `json:"enable"`
}

type CardGen struct {
	CardTypeID int  `json:"card_type_id" validate:"required"`
	Number     int  `json:"number" validate:"required,number"`
	Export     bool `json:"export" validate:"required,boolean"`
}
type RegUser struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Code     string `json:"code" validate:"required"`
}
type LoginUser struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type RefreshToken struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type ReCharge struct {
	CardNo string `json:"card_no" validate:"required"`
}
type ReChargeUp struct {
	Duration int `json:"duration" validate:"required,number"`
	Coin     int `json:"coin" validate:"required,number"`
}
type ChatMessage struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Model   string `json:"model"`
	Created int    `json:"created"`
	MsgID   string `json:"msgID,omitempty"`
	Choices []struct {
		Index int `json:"index"`
		Delta struct {
			Content string `json:"content,omitempty"`
		} `json:"delta"`
		FinishReason interface{} `json:"finish_reason"`
	} `json:"choices"`
}
type ChatChain struct {
	Content     []map[string]string `json:"content"`
	ParentMsgID string              `json:"msgID"`
	Token       int                 `json:"token"`
}
type ChatReq struct {
	Frequency_penalty int                 `json:"frequency_penalty" validate:"required,number"`
	Fax_tokens        int                 `json:"max_tokens" validate:"required,number"`
	Messages          []map[string]string `json:"messages" validate:"required"`
	Model             any                 `json:"model" validate:"required"`
	ParentMsgID       string              `json:"parentMsgID"`
	OriMsgID          string              `json:"oriMsgID"`
	Presence_penalty  int                 `json:"presence_penalty" validate:"required,number"`
	Stream            bool                `json:"stream"`
	Temperature       float64             `json:"temperature" validate:"required,number"`
	TopP              int                 `json:"top_p" validate:"required,number"`
}

type R2Bucket struct {
	Enable    string `json:"enable"`
	AccountID string `json:"accountID"`
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
	Bucket    string `json:"bucket"`
	EndPoint  string `json:"endPoint"`
}

type AliBucket struct {
	Enable    string `json:"enable"`
	EndPoint  string `json:"endPoint"`
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
	Bucket    string `json:"bucket"`
}

type TencentBucket struct {
	Enable    string `json:"enable"`
	EndPoint  string `json:"endPoint"`
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
}

type PreSign struct {
	Filename string `json:"filename" validate:"file_name"`
}

var ReqMap = map[string]func() interface{}{
	"CreateAdminReq":        func() interface{} { return &CreateAdminReq{} },
	"UpdateAdminReq":        func() interface{} { return &UpdateAdminReq{} },
	"GetUserInfoList":       func() interface{} { return &GetUserInfoList{} },
	"GetUserInfoListResp":   func() interface{} { return &GetUserInfoListResp{} },
	"CodelkupWithoutRef":    func() interface{} { return &CodelkupWithoutRef{} },
	"CodelkupWithRef":       func() interface{} { return &CodelkupWithRef{} },
	"update_email":          func() interface{} { return &UpdateEmailConfigReq{} },
	"update_email_reg":      func() interface{} { return &UpdateEmailRegReq{} },
	"update_wx_login":       func() interface{} { return &wxLogin{} },
	"ModelKeyCateList":      func() interface{} { return &ModelKeyCateList{} },
	"update_notice":         func() interface{} { return &Notice{} },
	"update_R2_bucket":      func() interface{} { return &R2Bucket{} },
	"update_Ali_bucket":     func() interface{} { return &AliBucket{} },
	"update_Tencent_bucket": func() interface{} { return &TencentBucket{} },
}

func ReqMapFunc(key string) interface{} {
	return ReqMap[key]()
}
