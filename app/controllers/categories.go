package controllers

import (
	"genuine/app/models"
	"genuine/core/controllers"
	"genuine/core/repositories"
)

func Categories(repository repositories.Repository[models.Category]) controllers.Controller {
	return Generic[models.Category](repository, "/accounting/categories")
}
