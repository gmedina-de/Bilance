package main

import (
	"Bilance/application"
	"Bilance/controller"
	"Bilance/repository"
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
	userRepository := repository.UserRepository(database)
	tagRepository := repository.TagRepository(database)

	authenticator := authenticator.BasicAuthenticator(userRepository)
	router := router.DefaultRouter(log, authenticator)
	server := server.DefaultServer(router, log, configuration)

	indexController := controller.IndexController(database)
	userController := controller.UserController(userRepository)
	tagController := controller.TagController(tagRepository)

	application := application.WebApplication(
		server,
		router,
		indexController,
		userController,
		tagController,
	)
	application.Run()
}
