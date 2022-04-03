package controllers

import (
	"genuine/core/controllers"
)

type index struct {
}

func Index() controllers.Controller {
	return &index{}
}

func (i *index) Routes() controllers.Routes {
	return controllers.Routes{
		"GET /": i.Index,
	}
}

func (i *index) Index() controllers.Template {
	// todo general functions for handlers
	return "index"
}
