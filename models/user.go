package models

type FzmpsUser struct {
    Avatar        string `json:"avatar"`
    Birthday      int    `json:"birthday"`
    Gender        int    `json:"gender"`
    Id            uint   `json:"id"`
    LastLoginIp   string `json:"last_login_ip"`
    LastLoginTime int64  `json:"last_login_time"`
    Mobile        string `json:"mobile"`
    Nickname      string `json:"nickname"       gorm:"column:nickname"`
    Password      string `json:"password"`
    RegisterIp    string `json:"register_ip"`
    RegisterTime  int64  `json:"register_time"`
    UserLevelId   int    `json:"user_level_id"`
    Username      string `json:"username"       gorm:"column:username"`
    Profession    string `json:"profession"`
    Age           string `json:"age"`

    Token   string `json:"token"`
    OpenId  string `json:"openid"         gorm:"column:openid"`
    UnionId string `json:"unionid"        gorm:"column:unionid"`
    UserId  string `json:"userId"         grom:"column:user_id"`
    Region  string `json:"region"`
}

func (FzmpsUser) TableName() string {
    return "fzmps_user"
}
