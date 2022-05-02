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

func (i *index) Index(controllers.Request) controllers.Response {
	// todo general functions for handlers
	return map[string]any{
		"Template": "index",
	}
}
