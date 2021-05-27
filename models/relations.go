package models

type FzmpsRelationDict struct {
	ID    uint   `json:"-"          gorm:"primarykey"`
	Label string `json:"label"      gorm:"label;type:varchar(64)"`
	Value string `json:"value"      gorm:"value;type:varchar(64)"`
}

func (FzmpsRelationDict) TableName() string {
	return "fzmps_relation_dict"
}
