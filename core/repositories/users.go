package repositories

import (
	"homecloud/core/database"
	"homecloud/core/models"
)

func Users(database database.Database) Repository[models.User] {
	return Generic(database, models.User{})
}
