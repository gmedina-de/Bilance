package main

import (
	"Bilance/application"
	"Bilance/controller"
	"Bilance/repository"
	"Bilance/service"
)

func main() {

	configuration := service.MapConfiguration()
	log := service.ConsoleLog()
	database := service.SqliteDatabase(log)
	authenticator := service.BasicAuthenticator(database)
	router := service.DefaultRouter(log, authenticator)
	server := service.DefaultServer(router, log, configuration)

	userRepository := repository.UserRepository(database)
	tagRepository := repository.TagRepository(database)
	projectRepository := repository.ProjectRepository(database, userRepository)

	controller.MyUserRepository = userRepository
	indexController := controller.IndexController(database)
	userController := controller.UserController(userRepository)
	tagController := controller.TagController(tagRepository)
	projectController := controller.ProjectController(projectRepository)

	application := application.BilanceApplication(
		server,
		router,
		indexController,
		userController,
		tagController,
		projectController,
	)
	application.Run()
}
