package controller

import (
	"Bilance/repository"
)

type projectController struct {
	baseController
}

func ProjectController(repository repository.Repository) Controller {
	return &projectController{
		baseController{
			repository: repository,
			basePath:   "/admin/project",
		},
	}
}
