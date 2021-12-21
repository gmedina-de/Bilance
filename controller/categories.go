package controller

import (
	"Bilance/model"
	"Bilance/repository"
)

type categories struct {
	generic[model.Category]
}

func Categories(repository repository.Repository[model.Category]) Controller {
	return &categories{
		generic[model.Category]{
			repository: repository,
			basePath:   "/categories/",
		},
	}
}
