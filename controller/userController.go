package controller

import (
	. "Bilance/model"
	"Bilance/service/database"
	"Bilance/service/router"
	"html/template"
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
	newUser := User{Username: "Admin6", Password: "asdf"}

	this.database.Insert(&newUser)
	var users []User

	tmpl, err := template.New("user.html").ParseFiles("view/user.html")

	err = tmpl.Execute(writer,
		struct {
			Users   []User
			NewUser User
		}{users, newUser},
	)
	if err != nil {
		panic(err)
	}
}
