package controllers

import (
	"genuine/core/controllers"
	"genuine/core/http"
)

type files struct {
}

func Files() controllers.Controller {
	return &files{}
}

func (f *files) Routes() map[string]http.Handler {
	return map[string]http.Handler{
		"GET /files": f.Index,
	}
}

func (f *files) Index(http.Request) http.Response {
	return nil
}
