package controllers

import (
	"genuine/apps/accounting/models"
	"genuine/core/controllers"
)

func Categories() controllers.Controller {
	return controllers.Generic(models.Category{}, "/accounting/categories")
}
