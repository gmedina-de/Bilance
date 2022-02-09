package controllers

import (
	"genuine/core/controllers"
	"genuine/core/router"
)

type files struct {
	controllers.BaseController
}

func Files() controllers.Controller {

	return &files{}
}

func (f *files) Routing() {
	router.Add(f, "/files", "get:Index()")
}

func (f *files) Index() {

}
