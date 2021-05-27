package dao

import (
	"centnet-fzmps/common/log"
	"centnet-fzmps/models"
	"strconv"
)

// 创建家庭成员
func (db *Dao) CreateFamilyMember(m *models.FzmpsFamily) {
	if err := db.db.Create(m).Error; err != nil {
		log.Error(err)
	}
}

// 删除家庭成员
func (db *Dao) DeleteFamilyMember(id string) {
	fid, err := strconv.Atoi(id)
	if err != nil {
		log.Error(err)
		return
	}

	if err := db.db.Where("id = ?", fid).Update("is_deleted", 1).Error; err != nil {
		log.Error(err)
	}
}

// 更新家庭成员
func (db *Dao) UpdateFamilyMember(m *models.FzmpsFamily) {
	tmp := &models.FzmpsFamily{Relation: m.Relation, Name: m.Name, Phone: m.Phone, Sex: m.Sex, Age: m.Age, QQ: m.QQ, WX: m.WX}
	if err := db.db.Where("id = ?", m.ID).Updates(tmp).Error; err != nil {
		log.Error(err)
	}
}

// 获取所有家庭成员
func (db *Dao) QueryFamilyMembers(userId uint) (m []models.FzmpsFamily) {
	if err := db.db.Where("user_id = ? and is_deleted = 0", userId).Find(&m); err != nil {
		log.Error(err)
	}
	return
}

// 获取家庭成员
func (db *Dao) QueryFamilyMember(id string) (m *models.FzmpsFamily) {
	fid, err := strconv.Atoi(id)
	if err != nil {
		log.Error(err)
		return
	}

	if err := db.db.Where("id = ?", fid).Find(&m).Error; err != nil {
		log.Error(err)
	}
	return
}
