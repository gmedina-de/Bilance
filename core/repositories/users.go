package repositories

import (
	"genuine/apps/users/models"
)

func Users() Repository[models.User] {
	return &Generic[models.User]{T: models.User{}}
}
