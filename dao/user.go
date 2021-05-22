package dao

import (
    "centnet-fzmps/common/log"
    "centnet-fzmps/models"
    "gorm.io/gorm"
)

func (d *Dao) InsertUser(u *models.FzmpsUser) {
    err := d.db.Create(&u).Error
    if err != nil {
        log.Error(err)
    }
}

func (d *Dao) QueryUserByOpenId(id string) (*models.FzmpsUser, error) {
    var (
        u   models.FzmpsUser
        err error
    )
    err = d.db.Where("openid = ?", id).First(&u).Error
    if err != nil {
        if err != gorm.ErrRecordNotFound {
            log.Error(err)
        }
        return &u, err
    }

    return &u, nil
}

func (d *Dao) QueryUserByToken(token string) (*models.FzmpsUser, error) {
    var (
        u   models.FzmpsUser
        err error
    )
    err = d.db.Where("token = ?", token).First(&u).Error
    if err != nil {
        if err != gorm.ErrRecordNotFound {
            log.Error(err)
        }
        return &u, err
    }

    return &u, nil
}

func (d *Dao) UpdateUserPhone(openId, phone string) {
    err := d.db.Model(&models.FzmpsUser{}).Where("openid = ?", openId).Update("mobile", phone).Error
    if err != nil {
        log.Error(err)
    }
}

func (d *Dao) UpdateUserInfo(openId, age, profession, realName string) {
    info := &models.FzmpsUser{Username: realName, Age: age, Profession: profession}
    err := d.db.Model(&models.FzmpsUser{}).Where("openid = ?", openId).Updates(info).Error
    if err != nil {
        log.Error(err)
    }
}

func (d *Dao) UpdateToken(openId, token string) {
    err := d.db.Model(&models.FzmpsUser{}).Where("openid = ?", openId).Update("token", token).Error
    if err != nil {
        log.Error(err)
    }
}

func (d *Dao) UpdateRegisterInfo(openId string, user *models.FzmpsUser) {
    err := d.db.Model(&models.FzmpsUser{}).Where("openid = ?", openId).Updates(user)
    if err != nil {
        log.Error(err)
    }
}
