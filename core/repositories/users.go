package repositories

import (
	"genuine/core/database"
	"genuine/core/models"
)

func Users(database database.Database) Repository[models.User] {
	return &Generic[models.User]{Database: database, Model: models.User{}}
}
