package controllers

import (
	"genuine/core/models"
	"genuine/core/repositories"
)

func Users(repository repositories.Repository[models.User]) Controller {
	return Generic(repository, models.User{}, "/settings/users")
}
