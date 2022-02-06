package database

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/mattn/go-sqlite3"
	"homecloud/core/log"
)

const Name = "default"
const Path = "./database.db"

func Orm(log log.Log) Database {
	orm.Debug = web.BConfig.RunMode == web.DEV
	log.Debug("Opening %s database from %s", Name, Path)
	orm.RegisterDataBase(Name, "sqlite3", Path)
	return orm.NewOrm()
}

func SyncDB(force bool, verbose bool) {
	err := orm.RunSyncdb(Name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
}
