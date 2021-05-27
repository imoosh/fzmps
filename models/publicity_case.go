package models

import (
	xtime "centnet-fzmps/common/time"
	"gorm.io/gorm"
)

type FzmpsPublicityCase struct {
	gorm.Model
	Title         string     `json:"title"            gorm:"title;type:varchar(200)"`
	Picture       string     `json:"picture"          gorm:"picture;type:text"`
	Content       string     `json:"content"          gorm:"content;type:longtext"`
	CaseType      int        `json:"case_type"        gorm:"case_type;type:bigint(10)"`
	Source        string     `json:"source"           gorm:"source;type:varchar(64)"`
	PublisherId   int        `json:"publisher_id"     gorm:"publisher_id;type:bigint(20)"`
	PublisherName string     `json:"publisher_name"   gorm:"publisher_name;type:varchar(32)"`
	PublisherTime xtime.Time `json:"publisher_time"   gorm:"publisher_time;type:datetime(3)"`
}

func (FzmpsPublicityCase) TableName() string {
	return "fzmps_publicity_case"
}
