package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"homecloud/core/database"
)

type index struct {
	BaseController
	Database database.Database
}

func Index(database database.Database) Controller {
	return &index{Database: database}
}

func (this *index) Routing() {
	web.Router("/", this, "get:Index")
}

func (this *index) Index() {
	if this.Database == nil {
		panic("AAAAAAAAAAAAAAAAAAA")
	}
}
