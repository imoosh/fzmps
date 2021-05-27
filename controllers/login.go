package controllers

import (
	"centnet-fzmps/common/log"
	"centnet-fzmps/conf"
	"centnet-fzmps/models"
	"centnet-fzmps/utils"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-uuid"
	"github.com/medivhzhan/weapp/v2"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

// 用户登录请求
type WeChatLoginReq struct {
	Code     string `json:"code"`
	IsPolice int    `json:"isPolice"`
}

// 用户登录响应
type WeChatLoginResp struct {
	OpenId  string `json:"openid"`
	Token   string `json:"token"`
	UnionId string `json:"unionId"`
	UserId  string `json:"userId"`
}

// 更新注册信息请求
type UpdateUserInfoReq struct {
	EncryptedData string  `json:"encryptedData"`
	IsPolice      int     `json:"isPolice"`
	IV            string  `json:"iv"`
	Lat           float32 `json:"lat"`
	Lng           float32 `json:"lng"`
	OpenId        string  `json:"openid"`
	UnionId       string  `json:"unionId"`
}

// 更新注册信息响应
type UpdateUserInfoResp struct {
	Region []string `json:"region"`
	Police struct {
		UserId   string `json:"userId"`
		RealName string `json:"realName"`
	}
	Organization struct {
		Id      string `json:"id"`
		OrgName string `json:"orgName"`
	}
	UserId string `json:"userId"`
}

func (req *UpdateUserInfoReq) String() string {
	return fmt.Sprintf("EncryptedData: %s | IsPolice: %d | IV: %s | OpenId: %s | UnionId: %s",
		req.EncryptedData, req.IsPolice, req.IV, req.OpenId, req.UnionId)
}

type PerfectUserInfoReq struct {
	Age           string   `json:"age"`
	Profession    string   `json:"profession"`
	OpenId        string   `json:"openid"`
	UnionId       string   `json:"unionId"`
	NickName      string   `json:"nickName"`
	AvatarUrl     string   `json:"avatarUrl"`
	Sex           int      `json:"sex"`
	RegionCode    []string `json:"regionCode"`
	Region        []string `json:"region"`
	IsPolice      int      `json:"isPolice"`
	EncryptedData string   `json:"encryptedData"`
	IV            string   `json:"iv"`
}

//https://developers.weixin.qq.com/miniprogram/dev/api/wx.getUserInfo.html
type Watermark struct {
	AppID     string `json:"appid"`
	TimeStamp int64  `json:"timestamp"`
}

//https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html
type WXUserInfo struct {
	OpenID    string    `json:"openId,omitempty"`
	NickName  string    `json:"nickName"`
	AvatarUrl string    `json:"avatarUrl"`
	Gender    int       `json:"gender"`
	Country   string    `json:"country"`
	Province  string    `json:"province"`
	City      string    `json:"city"`
	UnionID   string    `json:"unionId,omitempty"`
	Language  string    `json:"language"`
	Watermark Watermark `json:"watermark,omitempty"`
}

func (info *WXUserInfo) String() string {
	return fmt.Sprintf("OpenID: %s | NickName: %s | AvatarUrl: %s | Gender: %d | Country: %s | Province: %s | City: %s | UnionID: %s | Language: %s | %s %d",
		info.OpenID, info.NickName, info.AvatarUrl, info.Gender, info.Country, info.Province, info.City, info.UnionID, info.Language, info.Watermark.AppID, info.Watermark.TimeStamp)
}

//https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/getPhoneNumber.html
type WXPhoneNumber struct {
	PhoneNumber     string    `json:"phoneNumber"`
	PurePhoneNumber string    `json:"purePhoneNumber"`
	CountryCode     string    `json:"countryCode"`
	Watermark       Watermark `json:"watermark,omitempty"`
}

func (num *WXPhoneNumber) String() string {
	return fmt.Sprintf("PhoneNUmber: %s | PurePhoneNumber: %s | CountryCode: %s | %s %d", num.PhoneNumber,
		num.PurePhoneNumber, num.CountryCode, num.Watermark.AppID, num.Watermark.TimeStamp)
}

// 通过微信登陆
func WeChatLogin(c *gin.Context) {
	var (
		err       error
		req       WeChatLoginReq
		appId     = conf.Conf.Services.MiniPro.AppId
		appSecret = conf.Conf.Services.MiniPro.AppSecret
	)

	// 读取请求参数
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, utils.ReturnHTTPSuccess(nil))
		return
	}

	log.Debug(req)

	// 登陆微信平台
	loginResp, err := weapp.Login(appId, appSecret, req.Code)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusUnauthorized, utils.ReturnHTTPSuccess(nil))
		return
	}

	// 查表之前是否登录过
	var user *models.FzmpsUser
	user, err = DB(c).QueryUserByOpenId(loginResp.OpenID)
	if err == gorm.ErrRecordNotFound {
		var newUser = models.FzmpsUser{
			Avatar:        "",
			Birthday:      0,
			Gender:        0,
			Id:            0,
			LastLoginIp:   "",
			LastLoginTime: 0,
			Mobile:        "",
			Nickname:      "",
			Password:      "",
			RegisterIp:    c.Request.RemoteAddr,
			RegisterTime:  0,
			UserLevelId:   0,
			Username:      "",
			OpenId:        loginResp.OpenID,
			UnionId:       loginResp.UnionID,
			Token:         loginResp.SessionKey,
		}
		DB(c).InsertUser(&newUser)
		user, err = DB(c).QueryUserByOpenId(loginResp.OpenID)
		if err != nil {
			log.Error(err)
		}
	}

	if err != nil || user == nil {
		log.Debug("wechat login error")
		c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(nil))
		return
	}

	if user.Token != loginResp.SessionKey {
		DB(c).UpdateToken(loginResp.OpenID, loginResp.SessionKey)
	}

	var resp = WeChatLoginResp{
		OpenId:  user.OpenId,
		UnionId: user.UnionId,
		Token:   loginResp.SessionKey,
		UserId:  user.UserId,
	}

	log.Debugf("openid: %s | unionid: %s | token: %s", resp.OpenId, resp.UnionId, resp.Token)

	c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(resp))
}

// 更新信息
func UpdateRegisterInfo(c *gin.Context) {
	var (
		req UpdateUserInfoReq
	)
	token := c.GetHeader("Authorization")
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, utils.ReturnHTTPSuccess(nil))
		return
	}
	log.Debugf("EncryptedData: ... | IsPolice: %d | IV: %s | OpenId: %s | UnionId: %s | Token: %s",
		req.IsPolice, req.IV, req.OpenId, req.UnionId, token)
	info := decryptPhoneNumberData(token, req.EncryptedData, req.IV)
	if info != nil {
		// 更新手机号码
		DB(c).UpdateUserPhone(req.OpenId, info.PhoneNumber)
		log.Debug(info)
	}

	c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(&UpdateUserInfoResp{}))
}

// 完善个人信息
func PerfectRegisterInfo(c *gin.Context) {
	var (
		req PerfectUserInfoReq
	)
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, utils.ReturnHTTPSuccess(nil))
		return
	}
	log.Debug(req)

	//token := c.GetHeader("Authorization")
	//info := decryptUserInfoData(token, req.EncryptedData, req.IV)
	//if info != nil {
	//    log.Debug(info)
	//}

	userId, _ := uuid.GenerateUUID()
	user := models.FzmpsUser{
		Avatar:     req.AvatarUrl,
		Gender:     req.Sex,
		Nickname:   req.NickName,
		Profession: req.Profession,
		Age:        req.Age,
		Region:     strings.Join(req.Region, "|"),
		UserId:     userId,
	}

	//DB(c).UpdateRegisterInfo(req.OpenId, &user)
	DB(c).UpdateRegisterInfo(req.OpenId, &user)
	c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(&UpdateUserInfoResp{UserId: userId}))
}

func decryptUserInfoData(sessionKey string, encryptedData string, iv string) *WXUserInfo {

	sk, _ := base64.StdEncoding.DecodeString(sessionKey)
	ed, _ := base64.StdEncoding.DecodeString(encryptedData)
	i, _ := base64.StdEncoding.DecodeString(iv)

	decryptedData, err := utils.AesCBCDecrypt(ed, sk, i)
	if err != nil {
		log.Error("utils.AesCBCDecrypt error.")
		return nil
	}

	var wxUserInfo WXUserInfo
	//fmt.Println(string(decryptedData))
	err = json.Unmarshal(decryptedData, &wxUserInfo)
	if err != nil {
		log.Debug(err)
		return nil
	}
	return &wxUserInfo
}

func decryptPhoneNumberData(sessionKey string, encryptedData string, iv string) *WXPhoneNumber {

	sk, _ := base64.StdEncoding.DecodeString(sessionKey)
	ed, _ := base64.StdEncoding.DecodeString(encryptedData)
	i, _ := base64.StdEncoding.DecodeString(iv)

	if len(sk) == 0 {
		log.Error("sessionKey == 0")
		return nil
	}

	decryptedData, err := utils.AesCBCDecrypt(ed, sk, i)
	if err != nil {
		log.Error("utils.AesCBCDecrypt error.")
		return nil
	}

	var wxPhoneNum WXPhoneNumber
	err = json.Unmarshal(decryptedData, &wxPhoneNum)
	if err != nil {
		log.Debug(err)
		return nil
	}
	return &wxPhoneNum
}
