package main

import (
	"Bilance/application"
	"Bilance/authenticator"
	"Bilance/controller"
	"Bilance/database"
	"Bilance/log"
	"Bilance/repository"
	"Bilance/server"
)

func main() {
	application.Genuine(
		log.Console,
		database.Gorm,
		authenticator.Basic,
		server.Authenticated,
		repository.Users,
		repository.Categories,
		repository.Payments,
		repository.Projects,
		controller.Index,
		controller.Payments,
		controller.Categories,
		controller.Balances,
		controller.Expenses,
		controller.Users,
		controller.Projects,
	)
}
