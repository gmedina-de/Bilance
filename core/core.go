package core

import (
	"genuine/core/database"
	"genuine/core/injector"
	"genuine/core/log"
	"genuine/core/router"
	"genuine/core/server"
	"genuine/core/template"
	"genuine/core/translator"
)

func init() {
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
		inj.Constructor(constructor)
	}
}

func Init[T injector.Initiable](constructor func() T) {
	inj.Inject(constructor)
}
