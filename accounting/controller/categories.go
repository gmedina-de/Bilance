package controller

import (
	"homecloud/accounting/model"
	"homecloud/core/controller"
	"homecloud/core/repository"
)

type categories struct {
	controller.Generic[model.Category]
}

func Categories(repository repository.Repository[model.Category]) controller.Controller {
	return &categories{
		controller.Generic[model.Category]{
			BasePath:     "categories",
			BaseTemplate: "accounting/template/categories.gohtml",
			Repository:   repository,
		},
	}
}
