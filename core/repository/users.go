package repository

import (
	"homecloud/core/database"
	"homecloud/core/model"
)

type users struct {
	Repository[model.User]
}

func Users(database database.Database) Repository[model.User] {
	return &users{NewGeneric(database, model.User{})}
}
