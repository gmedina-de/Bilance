package controllers

import (
	"genuine/app/accounting/models"
	controllers2 "genuine/app/common/controllers"
	"genuine/core/controllers"
	"genuine/core/repositories"
)

func Categories(repository repositories.Repository[models.Category]) controllers.Controller {
	return controllers2.Generic[models.Category](repository, "/accounting/categories")
}
