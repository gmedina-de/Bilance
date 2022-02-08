package controllers

import (
	"genuine/core/router"
)

type index struct {
	BaseController
}

func Index() Controller {
	return &index{}
}

func (this *index) Routing() {
	router.Add(this, "/", "get:Index(id)")
}

func (this *index) Index(id string) {
	this.Data["Title"] = "Parameter id: " + id
}
