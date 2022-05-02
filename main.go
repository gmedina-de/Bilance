package main

import (
	"flag"
	controllers2 "genuine/controllers"
	"genuine/database"
	"genuine/decorators"
	"genuine/filters"
	"genuine/functions"
	"genuine/injector"
	"genuine/localizations"
	models2 "genuine/models"
	repositories2 "genuine/repositories"
	"genuine/server"
	"sync"
)

func init() {
	injector.Provide(
		controllers2.Balances,
		controllers2.Categories,
		controllers2.Expenses,
		controllers2.Files,
		controllers2.Index,
		controllers2.Payments,
		controllers2.Search,
		controllers2.Sites,
		controllers2.Users,
		database.Standard,
		decorators.Navigation,
		filters.Basic,
		functions.Form,
		functions.Paginate,
		functions.Table,
		localizations.All,
		repositories2.Categories,
		repositories2.Payments,
		repositories2.Sites,
		repositories2.Users,
		server.Webdav,
	)

	flag.Parse()

	registerAsset(models2.Person{})
	registerAsset(models2.Note{})
	registerAsset(models2.Book{})
}

func registerAsset[T models2.Asset](model T) {
	injector.Provide(
		func() models2.Asset {
			return model
		},
		func(database database.Database) repositories2.Repository[T] {
			return repositories2.Generic(database, model, "Id DESC")
		},
		func(repository repositories2.Repository[T]) controllers2.Controller {
			return controllers2.Generic(repository, "/assets/"+models2.Plural(model))
		},
	)
}

func main() {
	injector.Invoke(func(server []server.Server, translator functions.Translator) any {
		var wg sync.WaitGroup
		for _, s := range server {
			s := s
			wg.Add(1)
			go func() {
				s.Serve()
				wg.Done()
			}()
		}
		wg.Wait()
		return nil
	})
}
