package repositories

import (
	"genuine/core/models"
)

func Users() Repository[models.User] {
	return &Generic[models.User]{Model: models.User{}}
}
