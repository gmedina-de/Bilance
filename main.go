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

	userController := controller.UserController(userRepository)
	categoryController := controller.CategoryController(categoryRepository)
	expensesController := controller.ExpensesController(paymentRepository, categoryRepository)
	paymentController := controller.PaymentController(paymentRepository, categoryRepository, userRepository)
	projectController := controller.ProjectController(projectRepository, userRepository)

	bilance := Bilance(
		server,
		router,
		userController,
		categoryController,
		expensesController,
		paymentController,
		projectController,
	)
	bilance.Run()
}
