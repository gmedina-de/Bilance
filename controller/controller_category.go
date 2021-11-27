package controller

import (
	"Bilance/repository"
	"Bilance/service"
)

type typeController struct {
	crudController
}

func CategoryController(repository repository.Repository) Controller {
	return &typeController{
		crudController{
			repository: repository,
			basePath:   "/categories/",
		},
	}
}

func (c *typeController) Routing(router service.Router) {
	router.Get(c.basePath, c.List)
	router.Post(c.basePath, c.List)
	router.Get(c.basePath+"edit", c.Edit)
	router.Post(c.basePath+"edit", c.Edit)
	router.Get(c.basePath+"edit/delete", c.Delete)
}
