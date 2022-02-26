package repositories

import (
	"genuine/app/accounting/models"
	"genuine/core/database"
	"genuine/core/repositories"
)

func Payments(database database.Database) repositories.Repository[models.Payment] {
	return repositories.Generic(database, models.Payment{}, "Id DESC")
}
