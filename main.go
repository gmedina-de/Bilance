package main

import (
	"Bilance/controller"
	"Bilance/repository"
	"Bilance/service"
)

type bilance struct {
	controllers []controller.Controller
	server      service.Server
	router      service.Router
}

func Bilance(server service.Server, router service.Router, controllers ...controller.Controller) *bilance {
	return &bilance{controllers, server, router}
}

func (this *bilance) Run() {
	for _, c := range this.controllers {
		c.Routing(this.router)
	}
	this.server.Start()
}

func main() {

	log := service.ConsoleLog()
	database := service.SqliteDatabase(log)
	authenticator := service.BasicAuthenticator(database)
	router := service.DefaultRouter(log, authenticator)
	server := service.DefaultServer(router, log)

	users := repository.Users(database)
	categories := repository.Categories(database)
	payments := repository.Payments(database)
	projects := repository.Projects(database, payments, users, categories)

	bilance := Bilance(
		server,
		router,
		controller.Index(),
		controller.Payments(payments, categories, users),
		controller.Categories(categories),
		controller.Balances(projects, payments),
		controller.Expenses(payments, categories),
		controller.Users(users),
		controller.Projects(projects, users),
	)
	bilance.Run()
}
