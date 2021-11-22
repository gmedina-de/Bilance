package controller

import (
	"Bilance/model"
)

type projectController struct {
	baseController
}

func ProjectController(repository model.Repository) Controller {
	return &projectController{
		baseController{
			repository: repository,
			basePath:   "/admin/project",
		},
	}
}
