package controllers

import (
	"homecloud/core/database"
)

type index struct {
	BaseController
	Database database.Database
}

func Index(database database.Database) Controller {
	return &index{Database: database}
}

func (this *index) Routing(router Router) {
	router.Add("/", "get:Index(id)")
}

func (this *index) Index(id string) {
	this.Data["Title"] = "Parameter id: " + id

}
