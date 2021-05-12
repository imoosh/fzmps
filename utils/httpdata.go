package utils

import (
	"encoding/json"
	"github.com/astaxie/beego"
)

type HTTPData struct {
	Code int         `json:"code"` // 1: 成功
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ReturnHTTPSuccess(c *beego.Controller, val interface{}) {
	var ret = HTTPData{
		Code: 0,
		Msg:  "",
		Data: val,
	}
	if data, err := json.Marshal(ret); err != nil {
		c.Data["json"] = err
	} else {
		c.Data["json"] = json.RawMessage(data)
	}
}

func GetHTTPRtnJsonData(errno int, errmsg string) interface{} {
	var ret = HTTPData{
		Code: errno,
		Msg:  errmsg,
		Data: nil,
	}
	data, _ := json.Marshal(ret)
	return json.RawMessage(data)
}
