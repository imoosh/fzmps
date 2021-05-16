package controllers

import (
	"centnet-fzmps/common/log"
	"centnet-fzmps/utils"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	beego.Controller
}

// 响应用户信息
type UserInfoResp struct {
	Age          string        `json:"age"`
	Professional string        `json:"professional"`
	RealName     string        `json:"realName"`
	AllNum       string        `json:"allNum"`
	Days         string        `json:"days"`
	RelativeNum  string        `json:"relativeNum"`
	Relations    []interface{} `json:"relations"`
	OrgName      string        `json:"orgName"`
}

// 更新用户信息
type UserUpdateReq struct {
	Age          string `json:"age"`
	Professional string `json:"professional"`
	RealName     string `json:"realName"`
}

// api/user：请求用户信息
func (c *UserController) UserInfo() {
	var u UserInfoResp

	resp := make(map[string]interface{})
	resp["age"] = u.Age
	resp["professional"] = u.Professional
	resp["realName"] = u.RealName
	resp["allNum"] = u.AllNum
	resp["days"] = u.Days
	resp["relativeNum"] = u.RelativeNum
	resp["relations"] = nil
	resp["orgName"] = u.OrgName

	utils.ReturnHTTPSuccess(&c.Controller, resp)
	c.ServeJSON()
}

// api/user/update：更新用户信息
func (c *UserController) UpdateUserInfo() {
	var req UserUpdateReq
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	log.Debug(req)

	utils.ReturnHTTPSuccess(&c.Controller, nil)
	c.ServeJSON()
}

// 请求用户信息
func GetUserInfo(c *gin.Context) {
	var u UserInfoResp

	c.JSON(http.StatusOK, utils.ReturnHTTPSuccess1(&u))
}

// 更新用户信息
func UpdateUserInfo(c *gin.Context) {
	var req UserUpdateReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, utils.ReturnHTTPSuccess1(nil))
		return
	}

	c.JSON(http.StatusOK, utils.ReturnHTTPSuccess1(nil))
}
