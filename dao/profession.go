package dao

import (
	"centnet-fzmps/common/log"
	"centnet-fzmps/models"
)

func (db *Dao) QueryProfessionDict() (profession []models.FzmpsProfessionDict) {
	if err := db.db.Find(&profession).Error; err != nil {
		log.Error(err)
	}
	return
}
