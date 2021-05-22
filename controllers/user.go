package controllers

import (
    "centnet-fzmps/common/log"
    "centnet-fzmps/utils"
    "github.com/astaxie/beego"
    "github.com/gin-gonic/gin"
    "net/http"
    "strings"
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
    Region       []string      `json:"region"`
}

// 更新用户信息
type UserUpdateReq struct {
    Age        string `json:"age"`
    Profession string `json:"profession"`
    RealName   string `json:"realName"`
}

// 请求用户信息
func GetUserInfo(c *gin.Context) {

    token := c.GetHeader("Authorization")
    if len(token) == 0 {
        c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(&UserInfoResp{}))
        return
    }

    user, err := DB(c).QueryUserByToken(token)
    if err != nil {
        log.Error(err)
        c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(nil))
    }
    resp := UserInfoResp{
        Age:          user.Age,
        Professional: user.Profession,
        RealName:     user.Username,
        AllNum:       "",
        Days:         "",
        RelativeNum:  "",
        Relations:    nil,
        OrgName:      "",
        Region:       strings.Split(user.Region, "|"),
    }

    c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(&resp))
}

// 更新用户信息
func UpdateUserInfo(c *gin.Context) {
    var req UserUpdateReq

    err := c.ShouldBindJSON(&req)
    if err != nil {
        log.Error(err)
        c.JSON(http.StatusBadRequest, utils.ReturnHTTPSuccess(nil))
        return
    }

    //DB(c).UpdateUserInfo()

    c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(nil))
}
