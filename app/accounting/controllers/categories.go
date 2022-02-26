package controllers

import (
	"genuine/app/accounting/models"
	"genuine/core/controllers"
	"genuine/core/repositories"
)

func Categories(repository repositories.Repository[models.Category]) controllers.Controller {
	return controllers.Generic[models.Category](repository, "/accounting/categories")
}
