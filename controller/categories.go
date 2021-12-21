package controller

import (
	"Bilance/model"
	"Bilance/repository"
)

type categories struct {
	generic2[model.Category]
}

func Categories(repository repository.GRepository[model.Category]) Controller {
	return &categories{
		generic2[model.Category]{
			repository: repository,
			basePath:   "/categories/",
		},
	}
}
