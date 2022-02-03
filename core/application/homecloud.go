package application

import (
	"homecloud/core/controller"
	"homecloud/core/server"
	"reflect"
	"strings"
)

type homecloud struct {
	controllers []controller.Controller
	server      server.Server
}

func Homecloud(server server.Server, controllers []controller.Controller) *homecloud {
	return &homecloud{controllers, server}
}

func (b *homecloud) Run() {
	for _, c := range b.controllers {
		controllerType := reflect.TypeOf(c).Elem()

		path := "/" + controllerType.PkgPath()
		path = path[:strings.LastIndex(path, "/controller")]
		path = path[strings.LastIndex(path, "/"):]
		path = path + "/" + controllerType.Name()
		path = strings.ReplaceAll(path, "/index", "")
		path = strings.ReplaceAll(path, "core", "")
		path = strings.ReplaceAll(path, "//", "/")
		b.server.SetBasePath(path)

		c.Routing(b.server)
	}
	b.server.Start()
}
