package controllers

import (
	"genuine/framework/controllers"
	"genuine/framework/repositories"
	"genuine/prototype/accounting/models"
)

func Categories(repository repositories.Repository[models.Category]) controllers.Controller {
	return controllers.Generic(repository, models.Category{}, "/accounting/categories")
}
