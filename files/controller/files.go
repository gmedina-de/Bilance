package controllers

import (
	"genuine/core/controllers"
)

type files struct {
	controllers.BaseController
}

func Files() controllers.Controller {

	return &files{}
}

func (f *files) Routing() {

}
