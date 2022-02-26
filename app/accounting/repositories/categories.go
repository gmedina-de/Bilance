package repositories

import (
	"genuine/app/accounting/models"
	"genuine/core/database"
	"genuine/core/repositories"
)

func Categories(database database.Database) repositories.Repository[models.Category] {
	return repositories.Generic(database, models.Category{}, "Id DESC")
}
