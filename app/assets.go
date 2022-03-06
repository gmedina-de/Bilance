package app

import (
	"genuine/app/controllers"
	"genuine/app/database"
	"genuine/app/models"
	"genuine/app/repositories"
	"genuine/core"
	controllers2 "genuine/core/controllers"
	repositories2 "genuine/core/repositories"
)

func init() {
	register(models.Person{})
	register(models.Note{})
	register(models.Book{})
}

func register[T models.Asset](model T) {
	core.Provide(func() models.Asset {
		return model
	})
	core.Provide(func(database database.Database) repositories2.Repository[T] {
		return repositories.Generic(database, model, "Id DESC")
	})
	core.Provide(func(repository repositories2.Repository[T]) controllers2.Controller {
		return controllers.Generic(repository, "/assets/"+models.Plural(model))
	})
}
