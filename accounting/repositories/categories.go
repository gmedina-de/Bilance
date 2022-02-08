package repositories

import (
	"genuine/accounting/models"
	"genuine/core/database"
	"genuine/core/repositories"
)

type categories struct {
	repositories.Repository[models.Category]
}

func Categories(database database.Database) repositories.Repository[models.Category] {
	return &categories{repositories.Generic(database, models.Category{})}
}
