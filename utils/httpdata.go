package utils

import (
	"centnet-fzmps/common/log"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/gin-gonic/gin"
)

type HTTPData struct {
	ErrCode int         `json:"code"` // 1: 成功
	ErrMsg  string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func ReturnHTTPSuccess(c *beego.Controller, val interface{}) {
	var ret = HTTPData{
		ErrCode: 0,
		ErrMsg:  "",
		Data:    val,
	}
	if data, err := json.Marshal(ret); err != nil {
		c.Data["json"] = err
	} else {
		c.Data["json"] = json.RawMessage(data)
	}

	log.Debug(ret)
}

func GetHTTPRtnJsonData(errno int, errmsg string) interface{} {
	var ret = HTTPData{
		ErrCode: errno,
		ErrMsg:  errmsg,
		Data:    nil,
	}
	data, _ := json.Marshal(ret)
	return json.RawMessage(data)
}

var CommonH = gin.H{
	"ErrCode": 0,
	"ErrMsg":  "",
	"Data":    nil,
}

func ReturnHTTPSuccess1(val interface{}) gin.H {

	data, err := json.Marshal(val)
	if err != nil {
		log.Error(err)
	}

	var h = gin.H{
		"ErrCode": 0,
		"ErrMsg":  "ok",
		"Data":    data,
	}

	log.Debug(h)
	return h
}
