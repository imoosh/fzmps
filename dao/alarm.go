package dao

import (
	"centnet-fzmps/common/log"
	"centnet-fzmps/models"
	"gorm.io/gorm"
	"time"
)

func (db *Dao) InsertWeComAlarm(a *models.FzmpsAlarm) {
	if err := db.db.Model(&models.FzmpsAlarm{}).Create(a).Error; err != nil {
		log.Error(err)
	}
}

func (db *Dao) QueryAlarmByPhone(phone string) (a []models.FzmpsAlarm) {
	// 查询没查看的预警数据
	if err := db.db.Where("phone = ? and is_confirm = 0", phone).Find(&a).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			log.Error(err)
		}
	}
	return
}

func (db *Dao) UpdateAlarmConfirmStatus(phone string) {
	alarm := models.FzmpsAlarm{
		IsPushed:    true,
		IsConfirm:   true,
		ConfirmTime: time.Now().Format("2006-01-02 15:04:05"),
	}

	if err := db.db.Model(&models.FzmpsAlarm{}).Where("phone = ?", phone).Updates(alarm).Error; err != nil {
		log.Error(err)
	}
}

func (db *Dao) UpdateAlarmPushedStatus(id uint) {
	if err := db.db.Where("id = ?", id).Update("is_pushed", 1).Error; err != nil {
		log.Error(err)
	}
}
