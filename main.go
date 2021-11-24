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

	configuration := service.MapConfiguration()
	log := service.ConsoleLog()
	database := service.SqliteDatabase(log)
	authenticator := service.BasicAuthenticator(database)
	router := service.DefaultRouter(log, authenticator)
	server := service.DefaultServer(router, log, configuration)

	userRepository := repository.UserRepository(database)
	tagRepository := repository.TagRepository(database)
	projectRepository := repository.ProjectRepository(database, userRepository, tagRepository)

	indexController := controller.IndexController()
	userController := controller.UserController(userRepository)
	tagController := controller.TagController(tagRepository)
	projectController := controller.ProjectController(projectRepository)

	bilance := Bilance(
		server,
		router,
		indexController,
		userController,
		tagController,
		projectController,
	)
	bilance.Run()
}
