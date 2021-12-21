package application

import (
	"Bilance/controller"
	"Bilance/server"
)

type application struct {
	controllers []controller.Controller
	router      server.Server
}

func Application(router server.Server, controllers []controller.Controller) *application {
	return &application{controllers, router}
}

func (b *application) Run() {
	for _, c := range b.controllers {
		c.Routing(b.router)
	}
	b.router.Start()
}
