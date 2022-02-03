package core

import (
	"homecloud/core/application"
	"homecloud/core/authenticator"
	"homecloud/core/controller"
	"homecloud/core/database"
	"homecloud/core/injector"
	"homecloud/core/log"
	"homecloud/core/repository"
	"homecloud/core/server"
	"homecloud/core/template"
)

func init() {
	template.AddNavigation(
		template.MenuItem("home", "home", "/"),
	)
	template.AddNavigation(
		template.MenuItem("users", "user", "/users"),
	)
	Register(
		log.Console,
		database.Gorm,
		authenticator.Basic,
		server.Authenticated,
		repository.Users,
		controller.Index,
		controller.Users,
	)
}

var inj injector.Injector = injector.Recursive()

func Register(constructors ...interface{}) {
	inj.Add(constructors...)
}

func Init() {
	inj.Inject(application.Homecloud).Interface().(application.Application).Run()
}
