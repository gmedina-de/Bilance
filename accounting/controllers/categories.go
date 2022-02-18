package controllers

import (
	"genuine/accounting/models"
	"genuine/core/controllers"
)

func Categories() controllers.Controller {
	return controllers.Generic(models.Category{}, "/accounting/categories")
}
