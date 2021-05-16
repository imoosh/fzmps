package models

type FzmpsUser struct {
	Avatar        string `json:"avatar"`
	Birthday      int    `json:"birthday"`
	Gender        int    `json:"gender"`
	Id            int    `json:"id"`
	LastLoginIp   string `json:"last_login_ip"`
	LastLoginTime int64  `json:"last_login_time"`
	Mobile        string `json:"mobile"`
	Nickname      string `json:"nickname"`
	Password      string `json:"password"`
	RegisterIp    string `json:"register_ip"`
	RegisterTime  int64  `json:"register_time"`
	UserLevelId   int    `json:"user_level_id"`
	Username      string `json:"username"`
	OpenId        string `json:"openid"  orm:"column(openid)"`
	UnionId       string `json:"unionid" orm:"column(unionid)"`
}
