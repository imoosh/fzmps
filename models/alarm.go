package models

import (
	"centnet-fzmps/conf"
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"sync"
)

const (
	VictimLevelUnknown = 0
	VictimLevelLow     = 1
	VictimLevelMedium  = 2
	VictimLevelHigh    = 3
)

const (
	InfoTypeUnknown  = 0
	InfoTypePhone    = 1
	InfoTypeURL      = 2
	InfoTypeWeChat   = 3
	InfoTypeQQ       = 4
	InfoTypeApp      = 5
	InfoTypePlatForm = 6
	InfoTypeOther    = 7
)

// 受害等级
type VictimLevel int

func (level VictimLevel) String() string {
	switch level {
	case VictimLevelLow:
		return "低"
	case VictimLevelMedium:
		return "中"
	case VictimLevelHigh:
		return "高"
	default:
		return "未知"
	}
}

// 恶意账号/恶意信息
type InfoType int

func (t InfoType) String() string {
	switch t {
	case InfoTypePhone:
		return "电话号码"
	case InfoTypeURL:
		return "链接"
	case InfoTypeWeChat:
		return "微信"
	case InfoTypeQQ:
		return "QQ"
	case InfoTypeApp:
		return "APP"
	case InfoTypePlatForm:
		return "平台"
	case InfoTypeOther:
		return "其它"
	default:
		return "未知"
	}
}

// 诈骗类型
type EvilType string

func (t EvilType) String() string {
	switch t {
	case "sd":
		return "刷单诈骗"
	case "dk":
		return "虚假贷款诈骗"
	case "gjf":
		return "仿冒公检法诈骗"
	case "shopping_cheat":
		return "网购订单诈骗"
	case "recharge":
		return "游戏充值"
	case "gamble/invest/dating":
		return "赌博/投资/交友"
	case "fake_other":
		return "仿冒他人诈骗"
	case "other":
		return "其它类型诈骗"
	case "unknown":
		fallthrough
	default:
		return "未知诈骗"
	}
}

type AlarmMsg struct {
	AppId     string     `json:"AppId"`
	Seq       int        `json:"Seq"`
	Timestamp int        `json:"Timestamp"`
	SignId    string     `json:"SignId"`
	Data      FzmpsAlarm `json:"Data"`
}

type AlarmMsgRet struct {
	RetCode int    `json:"RetCode"`
	RetMsg  string `json:"RetMsg"`
}

func (m AlarmMsg) ValidateAppId() bool {
	return m.AppId == conf.Conf.API.WeCom.AppId
}

func (m AlarmMsg) VerifySignature() bool {
	c := conf.Conf.API.WeCom

	sum1 := md5.Sum([]byte(c.AppId + c.AppSecret + strconv.FormatInt(int64(m.Timestamp), 10)))
	sign := hex.EncodeToString(sum1[:])
	sum2 := md5.Sum([]byte(sign + strconv.FormatInt(int64(m.Seq), 10)))
	signId := hex.EncodeToString(sum2[:])

	//log.Debug(c.AppId + c.AppSecret + strconv.FormatInt(int64(m.Timestamp), 10))
	//log.Debug(sign)
	//
	//log.Debug(sign + strconv.FormatInt(int64(m.Seq), 10))
	//log.Debug(signId)

	return signId == m.SignId
}

type AlarmTask struct {
	Sender   string
	Receiver []string
	Content  string
}

var alarmPool = sync.Pool{
	New: func() interface{} { return new(FzmpsAlarm) },
}

func NewAlarm() *FzmpsAlarm {
	return alarmPool.Get().(*FzmpsAlarm)
}

func (va *FzmpsAlarm) Free() {
	alarmPool.Put(va)
}

type FzmpsAlarm struct {
	ID           uint   `json:"-"            gorm:"primarykey"`
	AlarmId      string `json:"alarm_id"     gorm:"alarm_id;type:varchar(64)"`
	Defender     string `json:"defender"     gorm:"defender;type:varchar(64)"`       // 劝阻人信息
	UnionId      string `json:"unionid"      gorm:"unionid;type:varchar(64)"`        // 被劝阻人微信union_id
	Phone        string `json:"phone"        gorm:"phone;type:varchar(32)"`          // 被劝阻人手机号
	AlarmMessage string `json:"alarm_message"gorm:"alarm_message;type:varchar(512)"` // 预警数据
	CreateTime   string `json:"time"         gorm:"create_time;type:varchar(32)"`    // 创建时间
	IsPushed     bool   `json:"is_pushed"    gorm:"is_pushed;type:int(4)"`           // 是否已推送过
	IsConfirm    bool   `json:"is_confirm"   gorm:"is_confirm;type:int(4)"`          // 是否已确认
	ConfirmTime  string `json:"confirm_time" gorm:"confirm_time;type:varchar(32)"`   // 确认时间
}

func (FzmpsAlarm) TableName() string {
	return "fzmps_alarm"
}
