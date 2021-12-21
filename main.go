package main

import (
	"Bilance/application"
	"Bilance/controller"
	"Bilance/repository"
	"Bilance/service"
)

func main() {
	application.Genuine(
		service.ConsoleLog,
		service.SqliteDatabase,
		service.BasicAuthenticator,
		service.HttpRouter,
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
