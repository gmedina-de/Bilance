package controllers

import (
	"genuine/core/controllers"
)

type files struct {
	controllers.Base
}

func Files() controllers.Controller {
	return &files{}
}

func (f *files) Routes() map[string]string {
	return map[string]string{
		"/files": "get:Index()",
	}
}

func (f *files) Index() {

}
