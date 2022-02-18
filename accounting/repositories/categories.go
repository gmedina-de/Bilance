package repositories

import (
	"genuine/accounting/models"
	"genuine/core/injector"
	"genuine/core/repositories"
)

func Categories() repositories.Repository[models.Category] {
	return injector.Inject(&repositories.Generic[models.Category]{
		Model: models.Category{},
	})
}
