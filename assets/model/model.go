package model

import (
	"homecloud/core"
	"homecloud/core/controller"
	"homecloud/core/database"
	model2 "homecloud/core/model"
	"homecloud/core/repository"
)

var Models []any
var Icons []string

func AddModel[T any](model T, icon string) {
	Models = append(Models, model)
	Icons = append(Icons, icon)

	core.Implementations(
		func(database database.Database) repository.Repository[T] {
			return repository.Generic(database, model)
		},
	)

	core.Implementations(
		func(repository repository.Repository[T]) controller.Controller {
			return &controller.Generic[T]{
				Model:      model,
				Repository: repository,
				BasePath:   "/assets/" + model2.Plural(model),
			}
		},
	)

}
