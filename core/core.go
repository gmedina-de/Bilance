package core

import (
	"genuine/core/controllers"
	"genuine/core/database"
	"genuine/core/injector"
	"genuine/core/log"
	"genuine/core/router"
	"genuine/core/server"
	"genuine/core/template"
	"genuine/core/translator"
)

func init() {
	Implementations(func() *controllers.Base { return &controllers.Base{} })
	Implementations(log.Standard)
	Implementations(database.Standard)
	Implementations(server.Standard)
	Implementations(router.Standard)
	Implementations(translator.Standard)
	Implementations(template.Standard)
}

var inj = injector.Standard()

func Implementations[T any](constructors ...func() T) {
	for _, constructor := range constructors {
		inj.Implementation(constructor)
	}
}

func Init[T any](constructor func() T) T {
	return inj.Inject(constructor).Interface().(T)
}
