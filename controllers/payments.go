package controllers

import (
	"genuine/models"
	"genuine/repositories"
)

func Payments(repository repositories.Repository[models.Payment]) Controller {
	return Generic(repository, "/accounting/payments")
}
