package dao

import (
    "centnet-fzmps/common/log"
    "centnet-fzmps/models"
)

func (d *Dao) InsertWeComAlarm(a *models.FzmpsAlarm) {
    if err := d.db.Model(&models.FzmpsAlarm{}).Create(a).Error; err != nil {
        log.Error(err)
    }
}
