package controllers

import (
	"genuine/database"
	"github.com/jinzhu/inflection"
	"reflect"
)

type search struct {
}

func Search(database database.Database) Controller {

	return &search{}
}

func (s *search) Routes() map[string]Handler {
	return map[string]Handler{
		"GET /search": s.Index,
	}
}

func (s *search) Index(r Request) Response {
	response := Response{"Template": "search"}
	var results = make(map[string]any)
	for _, search := range searchers {
		result := search(r)
		results[inflection.Plural(reflect.TypeOf(result).Elem().Name())] = result
	}
	response["Results"] = results
	return response
}

type searcher = func(r Request) any

var searchers []searcher
