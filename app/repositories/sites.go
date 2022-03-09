package repositories

import (
	"genuine/app/database"
	"genuine/app/models"
	"genuine/core/repositories"
)

func Sites(database database.Database) repositories.Repository[models.Site] {
	return Generic(database, models.Site{}, "Name ASC")
}
