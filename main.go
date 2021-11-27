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

	userRepository := repository.UserRepository(database)
	categoryRepository := repository.CategoryRepository(database)
	paymentRepository := repository.PaymentRepository(database, userRepository, categoryRepository)
	projectRepository := repository.ProjectRepository(database, paymentRepository, userRepository, categoryRepository)

	bilance := Bilance(
		server,
		router,
		controller.IndexController(),
		controller.PaymentController(paymentRepository, categoryRepository, userRepository),
		controller.CategoriesController(categoryRepository),
		controller.BalancesController(projectRepository, paymentRepository),
		controller.ExpensesController(paymentRepository, categoryRepository),
		controller.UsersController(userRepository),
		controller.ProjectsController(projectRepository, userRepository),
	)
	bilance.Run()
}
