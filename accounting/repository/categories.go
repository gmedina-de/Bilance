package repository

import (
	"homecloud/accounting/model"
	"homecloud/core/database"
	"homecloud/core/repository"
)

type categories struct {
	repository.Repository[model.Category]
}

func Categories(database database.Database) repository.Repository[model.Category] {
	return &categories{repository.Generic(database, model.Category{})}
}
