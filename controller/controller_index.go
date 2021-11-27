package controller

import (
	"Bilance/model"
	"Bilance/service"
	"net/http"
)

type indexController struct {
}

func IndexController() Controller {
	return &indexController{}
}

func (c *indexController) Routing(router service.Router) {
	router.Get("/", c.Index)
	router.Get("/changeProject", c.ChangeProject)
}

func (c *indexController) ChangeProject(writer http.ResponseWriter, request *http.Request) {
	selectedProjectId, _ := request.URL.Query()[model.SelectedProjectIdCookie]
	model.SetSelectedProjectId(writer, selectedProjectId[0])
	redirect(writer, request, "/")
}

func (c *indexController) Index(writer http.ResponseWriter, request *http.Request) {
	render(writer, request, nil, "", "index")
}
