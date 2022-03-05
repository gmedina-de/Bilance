package repositories

import (
	"genuine/app/database"
	"genuine/app/models"
	"genuine/core/repositories"
)

func Categories(database database.Database) repositories.Repository[models.Category] {
	return Generic(database, models.Category{}, "Id DESC")
}
