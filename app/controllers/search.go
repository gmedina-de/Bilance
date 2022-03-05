package controllers

import (
	"genuine/app/database"
	"genuine/core/controllers"
	"github.com/jinzhu/inflection"
	"reflect"
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

func (s *search) Index(r controllers.Request) controllers.Response {
	response := controllers.Response{"Template": "search"}
	var results = make(map[string]any)
	for _, search := range searchers {
		result := search(r)
		results[inflection.Plural(reflect.TypeOf(result).Elem().Name())] = result
	}
	response["Results"] = results
	return response
}

type searcher = func(r controllers.Request) any

var searchers []searcher
