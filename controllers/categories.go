package controllers

import (
	"genuine/models"
	"genuine/repositories"
)

func Categories(repository repositories.Repository[models.Category]) Controller {
	return Generic[models.Category](repository, "/accounting/categories")
}
