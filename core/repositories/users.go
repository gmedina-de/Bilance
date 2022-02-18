package repositories

import (
	"genuine/core/injector"
	"genuine/core/models"
)

func Users() Repository[models.User] {
	return injector.Inject(&Generic[models.User]{Model: models.User{}})
}
