package accounting

import (
	"gorm.io/gorm"
	"homecloud/core"
	"homecloud/core/controller"
	"homecloud/core/database"
	model2 "homecloud/core/model"
	"homecloud/core/repository"
	"homecloud/core/template"
)

func init() {

	models := []interface{}{Person{}, Note{}}

	for _, model := range models {
		core.Implementations(controllerProvider(model))
		core.Implementations(repositoryProvider(model))
	}

	menuItem := template.AddNavigation("assets", "box", "/assets/persons")
	for _, model := range models {
		menuItem.WithChild(model2.NamePlural(model), "box", "/assets/"+model2.NamePlural(model))
	}
}

type Person struct {
	gorm.Model
	Name string
	Age  int
}

type Note struct {
	gorm.Model
	Name        string
	Description string
}

func controllerProvider[T any](model T) func(repository repository.Repository[T]) controller.Controller {
	return func(repository repository.Repository[T]) controller.Controller {
		return &controller.Generic[T]{
			Repository:   repository,
			BaseTemplate: "assets/template/model.gohtml",
			BasePath:     "/assets/" + model2.NamePlural(model),
		}
	}
}

func repositoryProvider[T any](model T) func(database database.Database) repository.Repository[T] {
	return func(database database.Database) repository.Repository[T] {
		return repository.NewGeneric(database, model)
	}
}
