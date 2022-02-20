package repositories

import (
	"genuine/apps/accounting/models"
	"genuine/core/repositories"
)

func Categories() repositories.Repository[models.Category] {
	return &repositories.Generic[models.Category]{
		Model: models.Category{},
	}
}
