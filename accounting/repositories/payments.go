package repositories

import (
	"genuine/accounting/models"
	"genuine/core/repositories"
)

func Payments() repositories.Repository[models.Payment] {
	return &repositories.Generic[models.Payment]{
		Model:    models.Payment{},
		Ordering: "Date, Id DESC",
	}
}
