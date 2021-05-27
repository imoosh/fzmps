package dao

import (
	"centnet-fzmps/common/log"
	"centnet-fzmps/models"
	"gorm.io/gorm"
)

func (db *Dao) InsertUser(u *models.FzmpsUser) {
	err := db.db.Create(&u).Error
	if err != nil {
		log.Error(err)
	}
}

func (db *Dao) QueryUserByOpenId(id string) (*models.FzmpsUser, error) {
	var (
		u   models.FzmpsUser
		err error
	)
	err = db.db.Where("openid = ?", id).First(&u).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			log.Error(err)
		}
		return &u, err
	}

	return &u, nil
}

func (db *Dao) QueryUserByToken(token string) (*models.FzmpsUser, error) {
	var (
		u   models.FzmpsUser
		err error
	)
	err = db.db.Where("token = ?", token).First(&u).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			log.Error(err)
		}
		return &u, err
	}

	return &u, nil
}

func (db *Dao) UpdateUserPhone(openId, phone string) {
	err := db.db.Model(&models.FzmpsUser{}).Where("openid = ?", openId).Update("mobile", phone).Error
	if err != nil {
		log.Error(err)
	}
}

func (db *Dao) UpdateUserInfo(openId, age, profession, realName string) {
	info := &models.FzmpsUser{Username: realName, Age: age, Profession: profession}
	err := db.db.Model(&models.FzmpsUser{}).Where("openid = ?", openId).Updates(info).Error
	if err != nil {
		log.Error(err)
	}
}

func (db *Dao) UpdateToken(openId, token string) {
	err := db.db.Model(&models.FzmpsUser{}).Where("openid = ?", openId).Update("token", token).Error
	if err != nil {
		log.Error(err)
	}
}

func (db *Dao) UpdateRegisterInfo(openId string, user *models.FzmpsUser) {
	err := db.db.Model(&models.FzmpsUser{}).Where("openid = ?", openId).Updates(user)
	if err != nil {
		log.Error(err)
	}
}
