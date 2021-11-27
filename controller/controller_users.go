package controller

import (
	"Bilance/repository"
	"Bilance/service"
)

type usersController struct {
	crudController
}

func UserController(repository repository.Repository) Controller {
	return &usersController{
		crudController{
			repository: repository,
			basePath:   "/admin/users/",
		},
	}
}

func (c *usersController) Routing(router service.Router) {
	router.Get(c.basePath, c.List)
	router.Post(c.basePath, c.List)
	router.Get(c.basePath+"edit", c.Edit)
	router.Post(c.basePath+"edit", c.Edit)
	router.Get(c.basePath+"edit/delete", c.Delete)
}
