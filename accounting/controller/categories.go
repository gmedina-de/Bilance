package controller

import (
	"homecloud/accounting/model"
	"homecloud/core/controllers"
	"homecloud/core/repositories"
)

func Categories(repository repositories.Repository[model.Category]) controllers.Controller {
	return controllers.Generic(repository, model.Category{})
}
