package repositories

import (
	"genuine/apps/accounting/models"
	"genuine/core/repositories"
)

func Payments() repositories.Repository[models.Payment] {
	return repositories.NewGeneric(models.Payment{})

}
