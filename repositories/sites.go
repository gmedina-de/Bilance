package repositories

import (
	"genuine/database"
	"genuine/models"
)

func Sites(database database.Database) Repository[models.Site] {
	return Generic(database, models.Site{}, "Name ASC")
}
