package models

import (
	"genuine/core"
	"genuine/core/controllers"
	model2 "genuine/core/models"
	"genuine/core/repositories"
)

var Models []any
var Icons []string

func AddModel[T any](model T, icon string) {
	Models = append(Models, model)
	Icons = append(Icons, icon)

	core.Implementations(
		func() repositories.Repository[T] {
			return &repositories.Generic[T]{
				Model: model,
			}
		},
	)

	core.Implementations(
		func() controllers.Controller {
			return controllers.Generic(model, "/assets/"+model2.Plural(model))
		},
	)

}
