package services

import (
    "centnet-fzmps/common/log"
    "centnet-fzmps/models"
    "fmt"
    "github.com/gin-gonic/gin"
    "math/rand"
)

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
    var msg models.VictimAlarmMsg
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

    //a := models.FzmpsAlarm{
    //    Defender:    msg.Data.Defender,
    //    UnionId:     "",
    //    Phone:       msg.Data.Phone,
    //    Alarm:       formatMessage(&msg.Data),
    //    CreateTime:  time.Now().Format("2006-01-02 15:04:05"),
    //    IsPushed:    false,
    //    IsConfirm:   false,
    //    ConfirmTime: "",
    //}
    //
    //// 入库
    //controllers.DB(c)
}

func getOneUserId(ids []string) string {
    size := len(ids)
    if size == 0 {
        return ""
    } else if size == 1 {
        return ids[0]
    } else {
        // 随机选取一个
        return ids[rand.Intn(size)]
    }
}

func formatMessage(va *models.VictimAlarm) string {
    return fmt.Sprintf(""+
        "手机号码：%s\n"+
        "受害等级：%s\n"+
        "诈骗类型：%s\n"+
        "案发时间：%s\n"+
        "恶意账号：%s\n"+
        "诈骗信息：%s\n",
        va.Phone, va.VictimLevel.String(), va.EvilType.String(),
        va.OccurTime, va.InfoType.String(), va.EvilInfo)
}

// 定向预警劝阻
//func DirectedDissuasion(va *models.VictimAlarm) {
//
//    // 1. phone -> union_id
//    // redis中获取union_id
//    info, err := srv.dao.GetUserInfo(va.Phone)
//    if err != nil {
//        log.Errorf("dao.GetUserInfo(%s) failed: %v", va.Phone, err)
//        return
//    }
//    // redis未获取到union_id，从mysql中获取
//    if info == nil || len(info.UnionId) == 0 {
//        if info, err = srv.dao.QueryUserInfoByPhone(va.Phone); err != nil {
//            return
//        }
//        // 设置一个空的数据，1个小时超时
//        if len(info.UnionId) == 0 {
//            srv.dao.SetNullUserInfoWithExpired(va.Phone, info, 300)
//            return
//        }
//
//        // 缓存至redis
//        srv.dao.SetUserInfo(info.Phone, info)
//    }
//
//    // 2. union_id -> external_user_id + user_id[]
//    item, err := srv.dao.GetRelationship(info.UnionId)
//    if err != nil {
//        log.Errorf("dao.GetRelationship(%s) failed: %v", info.UnionId, err)
//        return
//    }
//    if item == nil || len(item.ExternalUserId) == 0 || len(item.DefenderIds) == 0 {
//        return
//    }
//    exUserId, userId := item.ExternalUserId, getOneUserId(item.DefenderIds)
//    if len(userId) == 0 {
//        return
//    }
//
//    // 3. push message
//    PushVictimAlarmMessage(&models.AlarmTask{
//        Sender:   userId,
//        Receiver: []string{exUserId},
//        Content:  formatMessage(va),
//    })
//}
//
//// 广播推送预警
//func BroadcastDissuasion(va *models.VictimAlarm) {
//    for userId, contacts := range weComClient.defenderGroup {
//        for exUserId, _ := range contacts {
//            PushVictimAlarmMessage(&models.AlarmTask{
//                Sender:   userId,
//                Receiver: []string{exUserId},
//                Content:  formatMessage(va),
//            })
//        }
//    }
//}
