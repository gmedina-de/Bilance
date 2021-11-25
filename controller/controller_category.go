package controller

import (
	"Bilance/model"
	"Bilance/repository"
	"Bilance/service"
	"net/http"
)

type typeController struct {
	baseController
}

func CategoryController(repository repository.Repository) Controller {
	return &typeController{
		baseController{
			repository: repository,
			basePath:   "/categories",
		},
	}
}

func (c *typeController) Routing(router service.Router) {
	router.Get(c.basePath, c.List)
	router.Post(c.basePath, c.List)
	router.Get(c.basePath+"/edit", c.Edit)
	router.Post(c.basePath+"/edit", c.Edit)
	router.Get(c.basePath+"/edit/delete", c.Delete)
}

func (c *typeController) List(writer http.ResponseWriter, request *http.Request) {
	cookie, _ := request.Cookie(model.SelectedProjectIdCookie)
	render(writer, request, &Parameters{Model: c.repository.List("WHERE ProjectId = " + cookie.Value)}, c.modelName(), "crud_table", c.modelName())
}
