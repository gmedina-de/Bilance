package controller

import (
	"Bilance/model"
)

type userController struct {
	baseController
}

func UserController(repository model.Repository) Controller {
	return &userController{
		baseController{
			repository: repository,
			basePath:   "/admin/user",
		},
	}
}
