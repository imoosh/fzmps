package controllers

import (
	"centnet-fzmps/common/log"
	"centnet-fzmps/models"
	"centnet-fzmps/utils"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/gin-gonic/gin"
	"github.com/medivhzhan/weapp/v2"
	"net/http"
)

type LoginController struct {
	beego.Controller
}

type LoginReq struct {
	Code     string `json:"code"`
	IsPolice string `json:"isPolice"`
}

type UpdateInfoReq struct {
	EncryptedData string `json:"encryptedData"`
	IV            string `json:"iv"`
	OpenId        string `json:"openid"`
	UnionId       string `json:"unionId"`
	IsPolice      string `json:"isPolice"`
}

type UpdateInfoResp struct {
	Region string
}

func (c *LoginController) Login() {
	var (
		err       error
		req       LoginReq
		appId     = "wx3a6389b153c6df95"
		appSecret = "6ca15b48f10d7359f0d0e28330ccd36a"
	)

	json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	log.Debug(req)

	loginResp, err := weapp.Login(appId, appSecret, req.Code)
	if err != nil {
		log.Error(err)
		return
	}

	var (
		user    models.FzmpsUser
		userTbl = new(models.FzmpsUser)
		o       = orm.NewOrm()
	)

	// 查表之前是否登录过
	err = o.QueryTable(userTbl).Filter("openid", loginResp.OpenID).One(&user)
	if err == orm.ErrNoRows {
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
			RegisterIp:    c.Ctx.Input.IP(),
			RegisterTime:  0,
			UserLevelId:   0,
			Username:      "",
			OpenId:        loginResp.OpenID,
			UnionId:       loginResp.UnionID,
		}
		_, err := o.Insert(&newUser)
		if err != nil {
			log.Error(err)
		}
		o.QueryTable(userTbl).Filter("openid", loginResp.OpenID).One(&user)
	} else if err != nil {
		log.Error(err)
	}

	resp := make(map[string]interface{})
	resp["openid"] = user.OpenId
	resp["unionId"] = user.UnionId
	resp["token"] = loginResp.SessionKey

	utils.ReturnHTTPSuccess(&c.Controller, resp)
	c.ServeJSON()
}

// 通过微信登陆
func WeChatLogin(c *gin.Context) {
	var (
		err       error
		req       LoginReq
		appId     = "wx3a6389b153c6df95"
		appSecret = "6ca15b48f10d7359f0d0e28330ccd36a"
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, utils.ReturnHTTPSuccess1(nil))
		return
	}

	log.Debug(req)

	loginResp, err := weapp.Login(appId, appSecret, req.Code)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusUnauthorized, utils.ReturnHTTPSuccess1(nil))
		return
	}

	var (
		user    models.FzmpsUser
		userTbl = new(models.FzmpsUser)
		o       = orm.NewOrm()
	)

	// 查表之前是否登录过
	err = o.QueryTable(userTbl).Filter("openid", loginResp.OpenID).One(&user)
	if err == orm.ErrNoRows {
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
		}
		_, err := o.Insert(&newUser)
		if err != nil {
			log.Error(err)
		}
		o.QueryTable(userTbl).Filter("openid", loginResp.OpenID).One(&user)
	} else if err != nil {
		log.Error(err)
	}

	resp := make(map[string]interface{})
	resp["openid"] = user.OpenId
	resp["unionId"] = user.UnionId
	resp["token"] = loginResp.SessionKey

	c.JSON(http.StatusOK, utils.ReturnHTTPSuccess1(resp))
}

// 完善注册信息
func UpdateRegisterInfo(c *gin.Context) {
	var (
		req UpdateInfoReq
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, utils.ReturnHTTPSuccess1(nil))
		return
	}

	c.JSON(http.StatusOK, utils.ReturnHTTPSuccess1(nil))
}
