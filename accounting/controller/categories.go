package controller

import (
	"homecloud/accounting/model"
	"homecloud/core/controllers"
	"homecloud/core/repository"
)

type categories struct {
	controllers.ControllerOld
}

func Categories(repository repository.Repository[model.Category]) controllers.ControllerOld {
	return &categories{
		&controllers.Generic[model.Category]{
			Repository: repository,
			BasePath:   "/accounting/categories",
		},
	}
}
