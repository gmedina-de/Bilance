package assets

import (
	"homecloud/assets/model"
	"homecloud/core"
	"homecloud/core/controller"
	model2 "homecloud/core/model"
	"homecloud/core/repository"
	"homecloud/core/template"
)

func init() {
	models := model.Models
	core.Implementations(repository.Agnostic)
	for _, m := range models {
		core.Implementations(controllerProvider(m))
	}

	menuItem := template.AddNavigation("assets", "box")
	for i, m := range models {
		menuItem.WithChild(model2.Plural(m), model.Icons[i])
	}
	menuItem.Path = "/assets/" + model2.Plural(models[0])
}

func controllerProvider[T any](model T) func(repository repository.Repository[any]) controller.Controller {
	return func(repository repository.Repository[any]) controller.Controller {
		return &controller.Generic[any]{
			Model:        model,
			Repository:   repository,
			BaseTemplate: "assets/template/model.gohtml",
			BasePath:     "/assets/" + model2.Plural(model),
		}
	}
}
