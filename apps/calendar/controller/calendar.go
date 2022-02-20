package controllers

import (
	"genuine/core/controllers"
	"genuine/core/router"
)

type calendar struct {
	controllers.Base
}

func Calendar() controllers.Controller {

	return &calendar{}
}

func (f *calendar) Routing() {
	router.Add(f, "/calendar", "get:Index()")
}

func (f *calendar) Index() {

}
