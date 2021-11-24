package controller

import (
	"Bilance/model"
	"Bilance/repository"
	"Bilance/service"
	"net/http"
)

type tagController struct {
	baseController
}

func TagController(repository repository.Repository) Controller {
	return &tagController{
		baseController{
			repository: repository,
			basePath:   "/tags",
		},
	}
}

func (c *tagController) Routing(router service.Router) {
	router.Get(c.basePath, c.List)
	router.Post(c.basePath, c.List)
	router.Get(c.basePath+"/new", c.New)
	router.Post(c.basePath+"/new", c.New)
	router.Get(c.basePath+"/edit", c.Edit)
	router.Post(c.basePath+"/edit", c.Edit)
	router.Get(c.basePath+"/edit/delete", c.Delete)
}

func (c *tagController) List(writer http.ResponseWriter, request *http.Request) {
	cookie, _ := request.Cookie(model.SelectedProjectIdCookie)
	search := c.repository.List("WHERE ProjectId = " + cookie.Value)
	render(writer, request, c.modelName()+"s", search, "crud_table", c.modelName())
}
