package controller

import (
	. "Bilance/model"
	"Bilance/service/database"
	"Bilance/service/router"
	"net/http"
)

type userController struct {
	database database.Database
}

func UserController(database database.Database) Controller {
	return &userController{database: database}
}

func (this *userController) Routing(router router.Router) {
	router.Get("/", this.SayHello)
}

func (this *userController) SayHello(writer http.ResponseWriter, request *http.Request) {
	user := User{Username: "Admin6", Password: "asdf"}
	this.database.Insert(&user)
	render(writer, "user", struct{ user User }{user})
}
