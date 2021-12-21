package controller

import (
	"Bilance/model"
	"Bilance/repository"
	"Bilance/service"
)

type usersController struct {
	crudController2[model.User]
}

func UsersController(repository repository.GRepository[model.User]) Controller {
	return &usersController{
		crudController2[model.User]{
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
