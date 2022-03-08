package controllers

import (
	"genuine/app/models"
	"genuine/core/controllers"
	"genuine/core/repositories"
)

func Books(repository repositories.Repository[models.Book]) controllers.Controller {
	return Generic[models.Book](repository, "/sites/books")
}
