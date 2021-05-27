package dao

import (
	"centnet-fzmps/common/log"
	"centnet-fzmps/models"
)

func (db *Dao) QueryRelationDict() (relation []models.FzmpsRelationDict) {
	if err := db.db.Find(&relation).Error; err != nil {
		log.Error(err)
	}
	return
}
