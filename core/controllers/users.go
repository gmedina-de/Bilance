package controllers

import (
	"homecloud/core/model"
	"homecloud/core/repositories"
)

type users struct {
	*Generic[model.User]
}

func Users(repository repositories.Repository[model.User]) Controller {
	return &users{Generics(repository, model.User{}, "/settings/users")}
}
