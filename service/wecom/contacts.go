package wecom

import (
    "bytes"
    "centnet-fzmps/common/log"
    "centnet-fzmps/conf"
    "crypto/md5"
    "encoding/hex"
    "encoding/json"
    "io/ioutil"
    "math/rand"
    "net/http"
    "strconv"
    "time"
)

type WeChatContactInfo struct {
    Phone   string `json:"Phone"     gorm:"column:phone; type:varchar(32);"`
    UnionId string `json:"UnionId"   gorm:"column:union_id; type:varchar(128);"`
}

func (WeChatContactInfo) TableName() string {
    return "user_info"
}

//func (u *UserInfo) BeforeCreate(tx *gorm.DB) (err error) {
//    return
//}

type WeChatContactInfoMsg struct {
    AppId     string              `json:"AppId"`
    Seq       int                 `json:"Seq"`
    Timestamp int64               `json:"Timestamp"`
    SignId    string              `json:"SignId"`
    Data      []WeChatContactInfo `json:"Data"`
}

func (m WeChatContactInfoMsg) signature(appId, appSecret string) string {
    sum1 := md5.Sum([]byte(appId + appSecret + strconv.FormatInt(int64(m.Timestamp), 10)))
    sign := hex.EncodeToString(sum1[:])
    sum2 := md5.Sum([]byte(sign + strconv.FormatInt(int64(m.Seq), 10)))
    signId := hex.EncodeToString(sum2[:])

    //log.Debug(c.AppId + c.AppSecret + strconv.FormatInt(int64(m.Timestamp), 10))
    //log.Debug(sign)
    //
    //log.Debug(sign + strconv.FormatInt(int64(m.Seq), 10))
    //log.Debug(signId)

    return signId
}

func PushContactInfoToWeCom(contacts []WeChatContactInfo) {
    var (
        appId     = conf.Conf.Services.WeCom.AppId
        appSecret = conf.Conf.Services.WeCom.AppSecret
        pushUrl   = conf.Conf.Services.WeCom.PushUrl
    )

    msg := &WeChatContactInfoMsg{
        AppId:     appId,
        Seq:       rand.Intn(10000),
        Timestamp: time.Now().Unix(),
        SignId:    "",
        Data:      contacts,
    }
    msg.SignId = msg.signature(appId, appSecret)

    js, err := json.Marshal(&msg)
    if err != nil {
        log.Error(err)
        return
    }

    log.Debug(string(js))
    resp, err := http.Post(pushUrl, "application/json", bytes.NewReader(js))
    if err != nil {
        log.Error(err)
        return
    }

    _, err = ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Error(err)
        return
    }
}
