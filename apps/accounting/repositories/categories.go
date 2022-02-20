package repositories

import (
	"genuine/apps/accounting/models"
	"genuine/core/repositories"
)

func Categories() repositories.Repository[models.Category] {
	return repositories.NewGeneric(models.Category{})
}
