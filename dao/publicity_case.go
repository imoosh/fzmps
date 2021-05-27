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
	if err := db.db.Model(&models.FzmpsPublicityCase{}).Find(&pc).Error; err != nil {
		log.Error(err)
	}
	return pc
}
