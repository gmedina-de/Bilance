package controller

import (
	"Bilance/model"
)

type tagController struct {
	baseController
}

func TagController(repository model.Repository) Controller {
	return &tagController{
		baseController{
			repository: repository,
			basePath:   "/admin/tag",
		},
	}
}
