package controller

import (
	"Bilance/repository"
	"Bilance/service"
)

type userController struct {
	baseController
}

func UserController(repository repository.Repository) Controller {
	return &userController{
		baseController{
			repository: repository,
			basePath:   "/admin/user",
		},
	}
}

func (c *userController) Routing(router service.Router) {
	router.Get(c.basePath, c.List)
	router.Post(c.basePath, c.List)
	router.Get(c.basePath+"/new", c.New)
	router.Post(c.basePath+"/new", c.New)
	router.Get(c.basePath+"/edit", c.Edit)
	router.Post(c.basePath+"/edit", c.Edit)
	router.Get(c.basePath+"/edit/delete", c.Delete)
}
