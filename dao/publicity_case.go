package dao

import (
	"centnet-fzmps/common/log"
	"centnet-fzmps/models"
)

func (db *Dao) CreatePublicityCase(pc *models.FzmpsPublicityCase) {
	if err := db.db.Model(&models.FzmpsPublicityCase{}).Create(pc).Error; err != nil {
		log.Error(err)
	}
}

func (db *Dao) QueryPublicityCase() (pc []models.FzmpsPublicityCase) {
	// 不包含content部分
	words := "id, created_at, updated_at, deleted_at, title, picture, case_type, source, publisher_id, publisher_name, publisher_time"
	if err := db.db.Model(&models.FzmpsPublicityCase{}).Select(words).Find(&pc).Error; err != nil {
		log.Error(err)
	}
	return pc
}

func (db *Dao) QueryPublicityCaseContent(id int) *models.FzmpsPublicityCase {
	var fpc models.FzmpsPublicityCase
	if err := db.db.Model(&models.FzmpsPublicityCase{}).Select("content").Where("id = ?", id).Find(&fpc).Error; err != nil {
		log.Error(err)
	}
	return &fpc
}
