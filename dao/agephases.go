package dao

import (
	"centnet-fzmps/common/log"
	"centnet-fzmps/models"
)

func (db *Dao) QueryAgePhasesDict() (agePhases []models.FzmpsAgePhasesDict) {
	if err := db.db.Find(&agePhases).Error; err != nil {
		log.Error(err)
	}
	return
}
