package repositories

import (
	"genuine/accounting/models"
	"genuine/core/inject"
	"genuine/core/repositories"
)

func Categories() repositories.Repository[models.Category] {
	return inject.Inject(&repositories.Generic[models.Category]{
		Model: models.Category{},
	})
}
