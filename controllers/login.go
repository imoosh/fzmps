package controllers

import (
	"centnet-fzmp/models"
	"centnet-fzmp/utils"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/medivhzhan/weapp/v2"
)

type LoginController struct {
	beego.Controller
}

type LoginReq struct {
	Code     string `json:"code"`
	IsPolice string `json:"isPolice"`
}

func (c *LoginController) Login() {
	var (
		err              error
		req              LoginReq
		appId, appSecret string
	)

	err = json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	loginResp, err := weapp.Login(appId, appSecret, req.Code)
	if err != nil {
		return
	}

	var (
		user    models.FzmpUser
		userTbl = new(models.FzmpUser)
		o       = orm.NewOrm()
	)

	// 查表之前是否登录过
	err = o.QueryTable(userTbl).Filter("openid", loginResp.OpenID).One(&user)
	if err == orm.ErrNoRows {
		var newUser = models.FzmpUser{
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
		_, _ = o.Insert(&newUser)
		o.QueryTable(userTbl).Filter("openid", loginResp.OpenID).One(&user)
	}

	resp := make(map[string]interface{})
	resp["openId"] = user.OpenId
	resp["unionId"] = user.UnionId
	resp["token"] = loginResp.SessionKey

	utils.ReturnHTTPSuccess(&c.Controller, resp)
	c.ServeJSON()
}
