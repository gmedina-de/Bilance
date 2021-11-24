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
	typeRepository := repository.TypeRepository(database)
	paymentRepository := repository.PaymentRepository(database, userRepository, typeRepository)
	projectRepository := repository.ProjectRepository(database, paymentRepository, userRepository, typeRepository)

	indexController := controller.IndexController()
	userController := controller.UserController(userRepository)
	typeController := controller.TypeController(typeRepository)
	paymentController := controller.PaymentController(paymentRepository)
	projectController := controller.ProjectController(projectRepository)

	bilance := Bilance(
		server,
		router,
		indexController,
		paymentController,
		projectController,
		typeController,
		userController,
	)
	bilance.Run()
}
