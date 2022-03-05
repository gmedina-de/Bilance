package controllers

import (
	"genuine/core/controllers"
)

type index struct {
}

func Index() controllers.Controller {
	return &index{}
}

func (i *index) Routes() map[string]controllers.Handler {
	return map[string]controllers.Handler{
		"GET /": i.Index,
	}
}

func (i *index) Index(controllers.Request) controllers.Response {
	return controllers.Response{"Template": "index"}
}
