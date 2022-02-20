package core

import (
	"genuine/core/database"
	"genuine/core/injector"
	"genuine/core/log"
	"genuine/core/router"
	"genuine/core/server"
	"genuine/core/template"
	"reflect"
)

func init() {

	template.AddNavigation("settings", "settings").
		WithChild("users", "users").
		Path = "/settings/users"

	Implementations(log.Console)
	Implementations(database.Orm)
	Implementations(server.Standard)
	Implementations(router.Standard)
}

func Implementations[T any](constructors ...func() T) {
	for _, constructor := range constructors {
		returnType := reflect.ValueOf(constructor).Type().Out(0)
		implementations[returnType] = append(implementations[returnType], constructor)
	}
}

func Init() {
	injector.Inject(App)
}

type app struct{ Server server.Server }

func App() injector.Initiable {
	return &app{}
}

func (a *app) Init() {
	a.Server.Start()
}
