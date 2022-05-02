package repositories

import (
	"genuine/database"
	"genuine/models"
)

func Payments(database database.Database) Repository[models.Payment] {
	return Generic(database, models.Payment{}, "Date DESC")
}
