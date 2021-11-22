package main

import (
	"Bilance/application"
	"Bilance/controller"
	"Bilance/model"
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
	userRepository := model.UserRepository(database)
	tagRepository := model.TagRepository(database)
	projectRepository := model.ProjectRepository(database)

	model.MyUserRepository = userRepository
	model.MyProjectRepository = projectRepository

	authenticator := authenticator.BasicAuthenticator(userRepository)
	router := router.DefaultRouter(log, authenticator)
	server := server.DefaultServer(router, log, configuration)

	indexController := controller.IndexController(database)
	userController := controller.UserController(userRepository)
	tagController := controller.TagController(tagRepository)
	projectController := controller.ProjectController(projectRepository)

	application := application.WebApplication(
		server,
		router,
		indexController,
		userController,
		tagController,
		projectController,
	)
	application.Run()
}
