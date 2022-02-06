package repositories

import (
	"homecloud/core/database"
	"homecloud/core/models"
)

type users struct {
	Repository[models.User]
}

func Users(database database.Database) Repository[models.User] {
	return &users{Generic(database, models.User{})}
}
