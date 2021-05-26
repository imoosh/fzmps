package dao

import (
    "centnet-fzmps/common/log"
    "centnet-fzmps/models"
)

func (db *Dao) QueryProvinces() (p []models.FzmpsProvince) {
    err := db.db.Select("province_name", "id").Find(&p).Error
    if err != nil {
        log.Error(err)
        return nil
    }
    return p
}

func (db *Dao) QueryCities(provId int) (c []models.FzmpsCity) {

    err := db.db.Select("city_id").Where("province_id = ?", provId).Find(&c).Error
    if err != nil {
        log.Error(err)
        return nil
    }
    return c
}

func (db *Dao) QueryCounties(provId, cityId int) (c []models.FzmpsCity) {
    err := db.db.Where("province_id = ? and city_id = ?", provId, cityId).Find(&c).Error
    if err != nil {
        log.Error(err)
        return nil
    }
    return c
}
