package application

import (
	"Bilance/controller"
	"Bilance/service"
)

type Application interface {
	Run()
}

type webApplication struct {
	controllers []controller.Controller
	server      service.Server
	router      service.Router
}

func WebApplication(server service.Server, router service.Router, controllers ...controller.Controller) *webApplication {
	return &webApplication{controllers, server, router}
}

func (this *webApplication) Run() {
	for _, c := range this.controllers {
		c.Routing(this.router)
	}
	this.server.Start()
}
