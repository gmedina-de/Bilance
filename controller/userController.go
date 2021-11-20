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

func (this *userController) Routing(router router.Router) {
	router.Get("/admin/user", this.List)
	router.Post("/admin/user", this.List)
	router.Get("/admin/user/new", this.New)
	router.Post("/admin/user/new", this.New)
	router.Get("/admin/user/edit", this.Edit)
	router.Post("/admin/user/edit", this.Edit)
	router.Get("/admin/user/delete", this.Delete)
}

func (this *userController) List(writer http.ResponseWriter, request *http.Request) {
	var users []model.User
	if request.URL.Query().Has("Search") {
		search := request.URL.Query().Get("Search")
		users = model.RetrieveUsers(this.database, "WHERE Username LIKE '"+search+"'", "ORDER BY Id")
	} else {
		users = model.RetrieveUsers(this.database, "ORDER BY Id")
	}
	render(writer, request, "user", struct{ Users []model.User }{users})
}

func (this *userController) New(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		render(writer, request, "userForm", &model.User{})
	} else if request.Method == "POST" {
		request.ParseForm()
		admin, _ := strconv.Atoi(request.Form.Get("Admin"))
		this.database.Insert(&model.User{
			0,
			request.Form.Get("Username"),
			request.Form.Get("Password"),
			admin,
		})
		http.Redirect(writer, request, "/user", http.StatusTemporaryRedirect)
	}
}

func (this *userController) Edit(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" && request.URL.Query().Has("Id") {
		id := request.URL.Query().Get("Id")
		render(writer, request, "userForm", model.RetrieveUsers(this.database, "WHERE Id = '"+id+"'")[0])
	} else if request.Method == "POST" {
		request.ParseForm()
		id, _ := strconv.Atoi(request.Form.Get("Id"))
		admin, _ := strconv.Atoi(request.Form.Get("Admin"))
		this.database.Update(&model.User{
			id,
			request.Form.Get("Username"),
			request.Form.Get("Password"),
			admin,
		})
		http.Redirect(writer, request, "/user", http.StatusTemporaryRedirect)
	}
}

func (this *userController) Delete(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" && request.URL.Query().Has("Id") {
		id := request.URL.Query().Get("Id")
		this.database.Delete("User", id)
		http.Redirect(writer, request, "/user", http.StatusTemporaryRedirect)
	}
}
