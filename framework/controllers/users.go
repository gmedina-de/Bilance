package controllers

import (
	"genuine/framework/models"
	"genuine/framework/repositories"
)

func Users(repository repositories.Repository[models.User]) Controller {
	return Generic(repository, models.User{}, "/settings/users")
}
