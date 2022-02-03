package controller

import (
	"homecloud/core/model"
	"homecloud/core/repository"
)

type users struct {
	Generic[model.User]
}

func Users(repository repository.Repository[model.User]) Controller {
	return &users{
		Generic[model.User]{
			Repository: repository,
		},
	}
}
