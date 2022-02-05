package controllers

import (
	"homecloud/core/model"
	"homecloud/core/repository"
)

type users struct {
	Controller
}

func Users(repository repository.Repository[model.User]) Controller {
	return &users{
		&Generic[model.User]{
			Repository: repository,
			BasePath:   "/settings/users",
		},
	}
}
