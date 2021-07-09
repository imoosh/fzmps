package wxmp

import (
    "centnet-fzmps/common/log"
    "centnet-fzmps/utils"
    "github.com/gin-gonic/gin"
    "net/http"
)

type alarmResp struct {
    Id        string `json:"id"`
    Sender    string `json:"sender"`
    Message   string `json:"message"`
    Time      string `json:"time"`
    IsConfirm bool   `json:"is_confirm"`
}

func RequestAlarm(c *gin.Context) {

    openid := c.Query("openid")
    if len(openid) == 0 {
        c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(nil))
        return
    }

    log.Debug("openid: ", openid)
    // 查询用户信息
    user, err := DB(c).QueryUserByOpenId(openid)
    if err != nil {
        c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(nil))
        return
    }
    log.Debug(user)

    // 查询预警数据
    alarms := DB(c).QueryAlarmByPhone(user.Mobile)
    if alarms == nil || len(alarms) == 0 {
        c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(nil))
        return
    }

    var resp []alarmResp
    for _, a := range alarms {
        resp = append(resp, alarmResp{
            Id:        a.AlarmId,
            Sender:    a.Defender,
            Message:   a.AlarmMessage,
            Time:      a.CreateTime,
            IsConfirm: a.IsConfirm,
        })
        DB(c).UpdateAlarmPushedStatus(a.ID)
    }

    c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(resp))

}

func ConfirmAlarm(c *gin.Context) {
    openid := c.Query("openid")
    alarmId := c.Query("alarmId")

    if len(openid) != 0 && len(alarmId) != 0 {
        u, err := DB(c).QueryUserByOpenId(openid)
        if err == nil && len(u.Mobile) != 0 {
            DB(c).UpdateAlarmConfirmStatus(alarmId, u.Mobile)
        }
    }

    c.JSON(http.StatusOK, utils.ReturnHTTPSuccess(nil))
}
