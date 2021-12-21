package application

import (
	"Bilance/controller"
	"Bilance/service"
)

type application struct {
	controllers []controller.Controller
	router      service.Router
}

func Application(router service.Router, controllers []controller.Controller) *application {
	return &application{controllers, router}
}

func (b *application) Run() {
	for _, c := range b.controllers {
		c.Routing(b.router)
	}
	b.router.Start()
}
