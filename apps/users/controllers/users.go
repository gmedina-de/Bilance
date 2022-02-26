package controllers

import (
	"genuine/apps/users/models"
	"genuine/core/controllers"
	"genuine/core/repositories"
)

func Users(repository repositories.Repository[models.User]) controllers.Controller {
	return controllers.Generic[models.User](repository, "/settings/users")
}
