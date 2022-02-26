package controllers

import (
	"genuine/core/controllers"
	"genuine/core/http"
)

type index struct {
}

func Index() controllers.Controller {
	return &index{}
}

func (i *index) Routes() map[string]http.Handler {
	return map[string]http.Handler{
		"GET /": i.Index,
	}
}

func (i *index) Index(http.Request) http.Response {
	return map[string]any{"Template": "index"}
}
