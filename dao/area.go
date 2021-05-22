package dao

import (
    "centnet-fzmps/common/log"
    "centnet-fzmps/models"
)

func (d *Dao) QueryProvinces() (p []models.Province) {
    err := d.db.Select("province_name", "id").Find(&p).Error
    if err != nil {
        log.Error(err)
        return nil
    }
    return p
}

func (d *Dao) QueryCities(provId int) (c []models.City) {

    err := d.db.Select("city_id").Where("province_id = ?", provId).Find(&c).Error
    if err != nil {
        log.Error(err)
        return nil
    }
    return c
}

func (d *Dao) QueryCounties(provId, cityId int) (c []models.City) {
    err := d.db.Where("province_id = ? and city_id = ?", provId, cityId).Find(&c).Error
    if err != nil {
        log.Error(err)
        return nil
    }
    return c
}
