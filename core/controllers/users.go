package controllers

import (
	"homecloud/core/models"
	"homecloud/core/repositories"
)

type users struct {
	Controller
}

func Users(repository repositories.Repository[models.User]) Controller {
	return &users{Generic(repository, models.User{}, "/settings/users")}
}
