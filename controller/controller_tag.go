package controller

import (
	"Bilance/repository"
)

type tagController struct {
	baseController
}

func TagController(repository repository.Repository) Controller {
	return &tagController{
		baseController{
			repository: repository,
			basePath:   "/admin/tag",
		},
	}
}
