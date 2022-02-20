package repositories

import (
	"genuine/apps/accounting/models"
	"genuine/core/repositories"
)

func Payments() repositories.Repository[models.Payment] {
	return &repositories.Generic[models.Payment]{
		T:        models.Payment{},
		Ordering: "Date, Id DESC",
	}
}
