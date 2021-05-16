package controllers

import (
	"centnet-fzmps/common/log"
	"centnet-fzmps/utils"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DictController struct {
	beego.Controller
}

type DictListReq struct {
	Type string `json:"type"`
}

// api/dict/list：获取字典信列表（年龄、职业、关联家人列表）
func (c *DictController) DictList() {
	var req DictListReq
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	log.Debug(req)

	switch req.Type {
	case "age_group":
	case "professional":
	case "relation":
	}
}

func GetDictList(c *gin.Context) {
	var req DictListReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, utils.ReturnHTTPSuccess1(nil))
		return
	}
	log.Debug(req)

	switch req.Type {
	case "age_group":
	case "professional":
	case "relation":
	}

	c.JSON(http.StatusOK, utils.ReturnHTTPSuccess1(nil))
}
