package controllers

import (
	"genuine/app/settings/models"
	"genuine/core/controllers"
	"genuine/core/repositories"
)

func Users(repository repositories.Repository[models.User]) controllers.Controller {
	return Generic[models.User](repository, "/settings/users")
}
