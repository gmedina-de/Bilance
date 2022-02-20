package core

import (
	"genuine/core/database"
	"genuine/core/injector"
	"genuine/core/log"
	"genuine/core/router"
	"genuine/core/server"
)

func init() {
	Implementations(log.Console)
	Implementations(database.Orm)
	Implementations(server.Standard)
	Implementations(router.Standard)
}

var inj = injector.Standard()

func Implementations[T any](constructors ...func() T) {
	for _, constructor := range constructors {
		inj.Implementation(constructor)
	}
}

func Init() {
	inj.Inject(App)
}

type app struct{ Server server.Server }

func App() injector.Initiable {
	return &app{}
}

func (a *app) Init() {
	a.Server.Start()
}
