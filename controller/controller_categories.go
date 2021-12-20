package controller

import (
	"Bilance/model"
	"Bilance/repository"
	"Bilance/service"
)

type categoriesController struct {
	crudController2[model.Category]
}

func CategoriesController(repository repository.GRepository[model.Category]) Controller {
	return &categoriesController{
		crudController2[model.Category]{
			repository: repository,
			basePath:   "/categories/",
		},
	}
}

func (c *categoriesController) Routing(router service.Router) {
	router.Get(c.basePath, c.List)
	router.Post(c.basePath, c.List)
	router.Get(c.basePath+"edit", c.Edit)
	router.Post(c.basePath+"edit", c.Edit)
	router.Get(c.basePath+"edit/delete", c.Delete)
}
