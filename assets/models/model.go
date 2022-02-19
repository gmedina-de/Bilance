package models

import (
	"genuine/core/controllers"
	"genuine/core/inject"
	model2 "genuine/core/models"
	"genuine/core/repositories"
)

var Models []any
var Icons []string

func AddModel[T any](model T, icon string) {
	Models = append(Models, model)
	Icons = append(Icons, icon)

	inject.Implementations(
		func() repositories.Repository[T] {
			return inject.Inject(&repositories.Generic[T]{
				Model: model,
			})
		},
	)

	inject.Implementations(
		func() controllers.Controller {
			return controllers.Generic(model, "/assets/"+model2.Plural(model))
		},
	)

}
