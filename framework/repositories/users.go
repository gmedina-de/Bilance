package repositories

import (
	"genuine/framework/database"
	"genuine/framework/models"
)

func Users(database database.Database) Repository[models.User] {
	return Generic(database, models.User{})
}
