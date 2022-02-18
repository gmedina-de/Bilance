package repositories

import (
	"genuine/accounting/models"
	"genuine/core/injector"
	"genuine/core/repositories"
)

func Payments() repositories.Repository[models.Payment] {
	return injector.Inject(&repositories.Generic[models.Payment]{
		Model:    models.Payment{},
		Ordering: "Date, Id DESC",
	})
}
