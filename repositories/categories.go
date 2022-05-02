package repositories

import (
	"genuine/database"
	"genuine/models"
)

func Categories(database database.Database) Repository[models.Category] {
	return Generic(database, models.Category{}, "Id DESC")
}
