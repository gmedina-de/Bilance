package repositories

import (
	"genuine/framework/database"
	"genuine/framework/repositories"
	"genuine/prototype/accounting/models"
)

type categories struct {
	repositories.Repository[models.Category]
}

func Categories(database database.Database) repositories.Repository[models.Category] {
	return &categories{repositories.Generic(database, models.Category{})}
}
