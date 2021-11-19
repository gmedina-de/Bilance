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
	baseRoute := "/user"
	router.Get(baseRoute+"/", this.List)
}

func (this *userController) List(writer http.ResponseWriter, request *http.Request) {
	users := RetrieveUsers(this.database)
	render(writer, "user", struct{ Users []User }{users})
}
