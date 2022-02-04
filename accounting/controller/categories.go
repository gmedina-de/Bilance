package controller

import (
	"homecloud/accounting/model"
	"homecloud/core/controller"
	"homecloud/core/repository"
)

type categories struct {
	controller.Controller
}

func Categories(repository repository.Repository[model.Category]) controller.Controller {
	return &categories{
		&controller.Generic[model.Category]{
			Repository:   repository,
			BaseTemplate: "accounting/template/categories.gohtml",
			BasePath:     "/accounting/categories",
		},
	}
}
