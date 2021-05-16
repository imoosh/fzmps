package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/nideshop?charset=utf8mb4", 30)

	// register model
	orm.RegisterModel(new(FzmpsUser))
}
