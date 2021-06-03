package wecom

import (
	"centnet-fzmps/common/log"
	"centnet-fzmps/models"
	"github.com/gin-gonic/gin"
	"time"
)

// 对接企业微信后台预警数据

var respStatus = map[int]gin.H{
	100: {"RetCode": 100, "RetMsg": "repeat request"},
	200: {"RetCode": 200, "RetMsg": "ok"},
	201: {"RetCode": 201, "RetMsg": "request volume exceeded"},
	301: {"RetCode": 301, "RetMsg": "invalid appid"},
	304: {"RetCode": 304, "RetMsg": "invalid signature"},
	500: {"RetCode": 500, "RetMsg": "unknown error"},
	501: {"RetCode": 501, "RetMsg": "internal error"},
}

// 腾讯反诈平台推送预警数据
func WeComVictimAlarm(c *gin.Context) {
	var msg models.AlarmMsg
	if err := c.BindJSON(&msg); err != nil {
		log.Error(err)
		c.JSON(200, respStatus[500])
		return
	}

	// 验证appid
	if !msg.ValidateAppId() {
		log.Errorf("invalid appid")
		c.JSON(200, respStatus[301])
		return
	}

	// 验证签名
	if !msg.VerifySignature() {
		log.Errorf("signature error (%s)", msg.SignId)
		c.JSON(200, respStatus[304])
		return
	}

	log.Debugf("From [%v]: %v", c.ClientIP(), msg.Data)

	msg.Data.CreateTime = time.Now().Format("2006-01-02 15:04:05")

	// 入库
	DB(c).InsertWeComAlarm(&msg.Data)
}
