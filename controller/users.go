package controller

import (
	"Bilance/model"
	"Bilance/repository"
)

type users struct {
	generic[model.User]
}

func Users(repository repository.Repository[model.User]) Controller {
	return &users{
		generic[model.User]{
			repository: repository,
			basePath:   "/admin/users/",
		},
	}
}
