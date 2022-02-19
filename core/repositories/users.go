package repositories

import (
	"genuine/core/inject"
	"genuine/core/models"
)

func Users() Repository[models.User] {
	return inject.Inject(&Generic[models.User]{Model: models.User{}})
}
