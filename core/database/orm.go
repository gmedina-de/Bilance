package database

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/mattn/go-sqlite3"
)

func Orm() orm.Ormer {
	orm.Debug = true
	orm.RegisterDataBase("default", "sqlite3", "./database.db")
	return orm.NewOrm()
}
