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
	tagRepository := repository.TagRepository(database)
	paymentRepository := repository.PaymentRepository(database, userRepository, tagRepository)
	projectRepository := repository.ProjectRepository(database, paymentRepository, userRepository, tagRepository)

	indexController := controller.IndexController()
	userController := controller.UserController(userRepository)
	tagController := controller.TagController(tagRepository)
	paymentController := controller.PaymentController(paymentRepository)
	projectController := controller.ProjectController(projectRepository)

	bilance := Bilance(
		server,
		router,
		indexController,
		paymentController,
		projectController,
		tagController,
		userController,
	)
	bilance.Run()
}
