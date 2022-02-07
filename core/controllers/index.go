package controllers

import (
	"homecloud/core/database"
	"homecloud/core/router"
)

type index struct {
	BaseController
	Database database.Database
}

func Index(database database.Database) Controller {
	return &index{Database: database}
}

func (this *index) Routing() {
	router.Add(this, "/", "get:Index(id)")
}

func (this *index) Index(id string) {
	this.Data["Title"] = "Parameter id: " + id

}
