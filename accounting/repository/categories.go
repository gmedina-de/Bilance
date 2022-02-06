package repository

import (
	"homecloud/accounting/model"
	"homecloud/core/database"
	"homecloud/core/repositories"
)

type categories struct {
	repositories.Repository[model.Category]
}

func Categories(database database.Database) repositories.Repository[model.Category] {
	return &categories{repositories.Generic(database, model.Category{})}
}
