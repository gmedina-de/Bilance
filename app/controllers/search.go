package controllers

import (
	"genuine/app/database"
	"genuine/core/controllers"
)

type search struct {
}

func Search(database database.Database) controllers.Controller {
	return &search{}
}

func (s *search) Routes() map[string]controllers.Handler {
	return map[string]controllers.Handler{
		"GET /search": s.Index,
	}
}

func (s *search) Index(controllers.Request) controllers.Response {

	return controllers.Response{
		"Template": "search",
	}
}
