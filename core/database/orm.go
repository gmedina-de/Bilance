package database

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/mattn/go-sqlite3"
)

const Name = "default"
const Path = "./database.db"

func Orm() orm.Ormer {
	orm.Debug = web.BConfig.RunMode == web.DEV
	orm.RegisterDataBase(Name, "sqlite3", Path)
	return orm.NewOrm()
}

func SyncDB(force bool, verbose bool) {
	_ = Orm()
	err := orm.RunSyncdb(Name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
}
