package application

import (
	"Bilance/controller"
	"Bilance/service"
)

type bilanceApplication struct {
	controllers []controller.Controller
	server      service.Server
	router      service.Router
}

func BilanceApplication(server service.Server, router service.Router, controllers ...controller.Controller) *bilanceApplication {
	return &bilanceApplication{controllers, server, router}
}

func (this *bilanceApplication) Run() {
	for _, c := range this.controllers {
		c.Routing(this.router)
	}
	this.server.Start()
}
