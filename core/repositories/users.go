package repositories

import (
	"genuine/core/database"
	"genuine/core/models"
)

func Users(database database.Database) Repository[models.User] {
	return Generic(database, models.User{})
}
