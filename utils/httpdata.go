package utils

import (
	"centnet-fzmps/common/log"
	"encoding/json"
)

type HTTPData struct {
	Code int         `json:"code"` // 1: 成功
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// code: 0      toast
// code: 1      silence
// code: 2/401  redirect to page/register/register
func ReturnHTTPSuccess(val interface{}) interface{} {
	var d = HTTPData{
		Code: 1,
		Msg:  "OK",
		Data: val,
	}

	js, _ := json.Marshal(&d)

	if len(js) < 1024 {
		log.Debug(string(js))
	}
	return d
}
