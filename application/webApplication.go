package application

import (
	"Bilance/controller"
	"Bilance/service/router"
	"Bilance/service/server"
)

type webApplication struct {
	controllers []controller.Controller
	server      server.Server
	router      router.Router
}

func WebApplication(server server.Server, router router.Router, controllers ...controller.Controller) *webApplication {
	return &webApplication{controllers, server, router}
}

func (this *webApplication) Run() {
	for _, c := range this.controllers {
		c.Routing(this.router)
	}
	this.server.Start()
}
