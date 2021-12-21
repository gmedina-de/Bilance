package controller

import (
	"Bilance/model"
	"Bilance/repository"
)

type users struct {
	generic2[model.User]
}

func Users(repository repository.GRepository[model.User]) Controller {
	return &users{
		generic2[model.User]{
			repository: repository,
			basePath:   "/admin/users/",
		},
	}
}
