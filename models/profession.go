package models

type FzmpsProfessionDict struct {
    ID    uint   `json:"-"          gorm:"primarykey"`
    Label string `json:"label"      gorm:"label;type:varchar(64)"`
    Value string `json:"value"      gorm:"value;type:varchar(64)"`
}

func (FzmpsProfessionDict) TableName() string {
    return "fzmps_profession_dict"
}
