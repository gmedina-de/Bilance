package controllers

import (
	"genuine/models"
	"genuine/repositories"
)

func Users(repository repositories.Repository[models.User]) Controller {
	return Generic[models.User](repository, "/settings/users")
}
