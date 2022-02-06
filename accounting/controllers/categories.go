package controllers

import (
	"homecloud/accounting/models"
	"homecloud/core/controllers"
	"homecloud/core/repositories"
)

func Categories(repository repositories.Repository[models.Category]) controllers.Controller {
	return controllers.Generics(repository, models.Category{}, "/accounting/categories")
}
