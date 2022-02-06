package controllers

import (
	"homecloud/core/models"
	"homecloud/core/repositories"
)

type users struct {
	*Generic[models.User]
}

func Users(repository repositories.Repository[models.User]) Controller {
	return &users{Generics(repository, models.User{}, "/settings/users")}
}
