package database

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/mattn/go-sqlite3"
)

const Name = "default"
const Path = "./database.db"

func Standard() Database {

	orm.Debug = web.BConfig.RunMode == web.DEV
	//log := injector.Inject[log.Log]((log.Log)(nil))
	//log.Debug("Opening %s database from %s", Name, Path)
	orm.RegisterDataBase(Name, "sqlite3", Path)
	return orm.NewOrm()
}
