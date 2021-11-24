package controller

import (
	"Bilance/repository"
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
