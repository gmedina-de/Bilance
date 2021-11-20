package main

import (
	"Bilance/application"
	"Bilance/controller"
	"Bilance/service/authenticator"
	"Bilance/service/configuration"
	"Bilance/service/database"
	"Bilance/service/log"
	"Bilance/service/router"
	server "Bilance/service/server"
)

func main() {

	configuration := configuration.MapConfiguration()
	log := log.ConsoleLog()
	database := database.SqliteDatabase(log)
	authenticator := authenticator.BasicAuthenticator(database)
	router := router.DefaultRouter(log, authenticator)
	server := server.DefaultServer(router, log, configuration)
	userController := controller.UserController(database)
	application := application.WebApplication(server, router, userController)
	application.Run()
}
