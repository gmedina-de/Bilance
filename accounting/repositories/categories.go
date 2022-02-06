package repositories

import (
	"homecloud/accounting/models"
	"homecloud/core/database"
	"homecloud/core/repositories"
)

type categories struct {
	repositories.Repository[models.Category]
}

func Categories(database database.Database) repositories.Repository[models.Category] {
	return &categories{repositories.Generic(database, models.Category{})}
}
