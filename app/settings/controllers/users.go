package controllers

import (
	controllers2 "genuine/app/common/controllers"
	"genuine/app/common/models"
	"genuine/core/controllers"
	"genuine/core/repositories"
)

func Users(repository repositories.Repository[models.User]) controllers.Controller {
	return controllers2.Generic[models.User](repository, "/settings/users")
}
