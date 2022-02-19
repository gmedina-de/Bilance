package repositories

import (
	"genuine/accounting/models"
	"genuine/core/inject"
	"genuine/core/repositories"
)

func Payments() repositories.Repository[models.Payment] {
	return inject.Inject(&repositories.Generic[models.Payment]{
		Model:    models.Payment{},
		Ordering: "Date, Id DESC",
	})
}
