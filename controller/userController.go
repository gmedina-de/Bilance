package controller

import (
	"Bilance/model"
	"Bilance/service/database"
	"Bilance/service/router"
	"net/http"
	"strconv"
)

type userController struct {
	database database.Database
}

func UserController(database database.Database) Controller {
	return &userController{database: database}
}

const userBasePath = "/admin/user"

func (this *userController) Routing(router router.Router) {
	router.Get(userBasePath, this.List)
	router.Post(userBasePath, this.List)
	router.Get(userBasePath+"/new", this.New)
	router.Post(userBasePath+"/new", this.New)
	router.Get(userBasePath+"/edit", this.Edit)
	router.Post(userBasePath+"/edit", this.Edit)
	router.Get(userBasePath+"/delete", this.Delete)
}

func (this *userController) List(writer http.ResponseWriter, request *http.Request) {
	var users []model.User
	if request.URL.Query().Has("Search") {
		search := request.URL.Query().Get("Search")
		this.database.Query(&users, model.TagQuery, "WHERE Username LIKE '%"+search+"%'", "ORDER BY Id")
	} else {
		this.database.Query(&users, model.UserQuery, "ORDER BY Id")
	}
	render(writer, request, "user", struct{ Users []model.User }{users})
}

func (this *userController) New(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		render(writer, request, "userForm", &model.User{})
	} else if request.Method == "POST" {
		request.ParseForm()
		admin, _ := strconv.Atoi(request.Form.Get("UserRole"))
		this.database.Insert(&model.User{
			0,
			request.Form.Get("Username"),
			request.Form.Get("Password"),
			model.UserRole(admin),
		})
		http.Redirect(writer, request, userBasePath, http.StatusTemporaryRedirect)
	}
}

func (this *userController) Edit(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" && request.URL.Query().Has("Id") {
		id := request.URL.Query().Get("Id")
		var users []model.User
		this.database.Query(&users, model.UserQuery, "WHERE Id = '"+id+"'")
		user := users[0]
		render(writer, request, "userForm", user)
	} else if request.Method == "POST" {
		request.ParseForm()
		id, _ := strconv.Atoi(request.Form.Get("Id"))
		admin, _ := strconv.Atoi(request.Form.Get("UserRole"))
		this.database.Update(&model.User{
			id,
			request.Form.Get("Username"),
			request.Form.Get("Password"),
			model.UserRole(admin),
		})
		http.Redirect(writer, request, userBasePath, http.StatusTemporaryRedirect)
	}
}

func (this *userController) Delete(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" && request.URL.Query().Has("Id") {
		id := request.URL.Query().Get("Id")
		this.database.Delete("User", id)
		http.Redirect(writer, request, userBasePath, http.StatusTemporaryRedirect)
	}
}
