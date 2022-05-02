package repositories

import (
	"genuine/app/database"
	"genuine/app/models"
	"genuine/core/repositories"
)

func Payments(database database.Database) repositories.Repository[models.Payment] {
	return Generic(database, models.Payment{}, "Date DESC")
}
