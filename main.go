package main

import (
	"Bilance/application"
	"Bilance/controller"
	"Bilance/service"
)

func main() {

	configuration := service.MapConfiguration()
	log := service.ConsoleLog(configuration)
	repository := service.DbRepository()
	router := service.DefaultRouter(log)
	server := service.HttpServer(router, log, configuration)
	userController := controller.UserController(repository)
	webApplication := application.WebApplication(server, router, userController)
	webApplication.Run()
}
