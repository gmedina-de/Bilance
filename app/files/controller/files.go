package controllers

import (
	"genuine/core/controllers"
)

type files struct {
}

func Files() controllers.Controller {
	return &files{}
}

func (f *files) Routes() map[string]controllers.Handler {
	return map[string]controllers.Handler{
		"GET /files": f.Index,
	}
}

func (f *files) Index(controllers.Request) controllers.Response {
	return nil
}
