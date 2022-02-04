package application

import (
	"homecloud/core/controller"
	"homecloud/core/server"
	"homecloud/core/template"
)

type homecloud struct {
	controllers []controller.Controller
	server      server.Server
}

func Homecloud(server server.Server, controllers []controller.Controller) *homecloud {
	return &homecloud{controllers, server}
}

func (b *homecloud) Run() {
	// adds settings to the end of navigation
	template.AddNavigation("settings", "settings").
		WithChild("users", "users").
		Path = "/settings/users"

	for _, c := range b.controllers {
		c.Routing(b.server)
	}
	b.server.Start()
}
