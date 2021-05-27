package models

// 家人表
type FzmpsFamily struct {
	ID        uint   `json:"-"              gorm:"primarykey"`
	UserId    uint   `json:"user_id"        gorm:"user_id;type:int(20)"`      // 关联的小程序用户id
	Relation  string `json:"relation"       gorm:"relation;type:varchar(32)"` // 与关联用户的家庭关系id
	Name      string `json:"name"           gorm:"name;type:varchar(32)"`
	Phone     string `json:"phone"          gorm:"phone;type:varchar(20)"`
	Sex       string `json:"sex"            gorm:"sex;type:varchar(4)"`
	Age       string `json:"age"            gorm:"age;type:varchar(64)"`
	QQ        string `json:"qq"             gorm:"qq;type:varchar(32)"`
	WX        string `json:"wx"             gorm:"wx;type:varchar(32)"`
	IsDeleted bool   `json:"is_deleted"     gorm:"is_deleted;type:int(2)"`
}

func (FzmpsFamily) TableName() string {
	return "fzmps_family"
}
