package dao

import (
    "centnet-fzmps/common/database/orm"
    "centnet-fzmps/common/log"
    "centnet-fzmps/conf"
    "centnet-fzmps/models"
    "errors"
    "gorm.io/gorm"
)

type Dao struct {
    c  *conf.Config
    db *gorm.DB
}

func New(c *conf.Config) (dao *Dao, err error) {
    dao = &Dao{c: c}
    if c.ORM != nil {
        dao.db = orm.NewMySQL(c.ORM)
        if dao.db == nil {
            return nil, errors.New("orm.NewMySQL failed")
        }
    }

    if err := dao.init(); err != nil {
        return nil, err
    }

    return dao, nil
}

func (dao *Dao) init() error {
    var options = "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;"
    err := dao.db.Set("gorm:table_options", options).AutoMigrate(&models.FzmpsUser{}, &models.VictimAlarm{}, &models.FzmpsAlarm{})
    if err != nil {
        log.Error(err)
        return err
    }

    return nil
}
