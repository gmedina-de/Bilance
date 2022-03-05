package controllers

import (
	"genuine/core/controllers"
	"genuine/core/database"
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
