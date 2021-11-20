package controller

import (
	"Bilance/service/database"
	"Bilance/service/router"
	"net/http"
)

type indexController struct {
	database database.Database
}

func IndexController(database database.Database) Controller {
	return &indexController{database: database}
}

func (this *indexController) Routing(router router.Router) {
	router.Get("/", this.Index)
}

func (this *indexController) Index(writer http.ResponseWriter, request *http.Request) {
	render(writer, "index", nil)
}
