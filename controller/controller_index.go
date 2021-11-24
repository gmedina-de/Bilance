package controller

import (
	"Bilance/service"
	"net/http"
)

type indexController struct {
	database service.Database
}

func IndexController(database service.Database) Controller {
	return &indexController{database: database}
}

func (this *indexController) Routing(router service.Router) {
	router.Get("/", this.Index)
}

func (this *indexController) Index(writer http.ResponseWriter, request *http.Request) {
	render(writer, request, "dashboard", nil, "index")
}
