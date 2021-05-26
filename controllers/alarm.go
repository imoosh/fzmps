package controllers

import (
    "centnet-fzmps/utils"
    "github.com/gin-gonic/gin"
    "net/http"
)

type alarmItem struct {
    id      string `json:"id"`
    sender  string `json:"sender"`
    message string `json:"message"`
    time    string `json:"time"`
}

type alarmResp struct {
    alarms []alarmItem
}

func RequestAlarm(c *gin.Context) {
    var resp alarmResp

    openid := c.Query("openid")
    if len(openid) == 0 {
        c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(&resp))
        return
    }

    // 查询用户信息
    user, err := DB(c).QueryUserByOpenId(openid)
    if err != nil {
        c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(&resp))
        return
    }

    // 查询预警数据
    alarms := DB(c).QueryAlarmByPhone(user.Mobile)
    if alarms == nil || len(alarms) == 0 {
        c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(&resp))
        return
    }

    for _, a := range alarms {
        item := alarmItem{
            id:      a.AlarmId,
            sender:  a.Defender,
            message: a.AlarmMessage,
            time:    a.CreateTime,
        }
        resp.alarms = append(resp.alarms, item)
        DB(c).UpdateAlarmPushedStatus(a.ID)
    }
    c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(&resp))

}

func ConfirmAlarm(c *gin.Context) {
    openid := c.Query("openid")
    alarmId := c.Query("alarmId")

    if len(openid) != 0 && len(alarmId) != 0 {
        u, err := DB(c).QueryUserByOpenId(openid)
        if err == nil && len(u.Mobile) != 0 {
            DB(c).UpdateAlarmConfirmStatus(u.Mobile)
        }
    }

    c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(nil))
}
