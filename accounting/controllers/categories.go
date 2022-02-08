package controllers

import (
	"genuine/accounting/models"
	"genuine/core/controllers"
	"genuine/core/repositories"
)

func Categories(repository repositories.Repository[models.Category]) controllers.Controller {
	return controllers.Generic(repository, models.Category{}, "/accounting/categories")
}
