package controller

import (
	"Bilance/model"
	"Bilance/service/database"
	"Bilance/service/router"
	"net/http"
	"strconv"
)

type tagController struct {
	database database.Database
}

func TagController(database database.Database) Controller {
	return &tagController{database: database}
}

const tagBasePath = "/admin/tag"

func (this *tagController) Routing(router router.Router) {
	router.Get(tagBasePath, this.List)
	router.Post(tagBasePath, this.List)
	router.Get(tagBasePath+"/new", this.New)
	router.Post(tagBasePath+"/new", this.New)
	router.Get(tagBasePath+"/edit", this.Edit)
	router.Post(tagBasePath+"/edit", this.Edit)
	router.Get(tagBasePath+"/delete", this.Delete)
}

func (this *tagController) List(writer http.ResponseWriter, request *http.Request) {
	var tags []model.Tag
	if request.URL.Query().Has("Search") {
		search := request.URL.Query().Get("Search")
		this.database.Query(&tags, model.TagQuery, "WHERE Name LIKE '%"+search+"%'", "ORDER BY Id")
	} else {
		this.database.Query(&tags, model.TagQuery, "ORDER BY Id")
	}
	render(writer, request, "tag", struct{ Tags []model.Tag }{tags})
}

func (this *tagController) New(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		render(writer, request, "tagForm", &model.Tag{})
	} else if request.Method == "POST" {
		request.ParseForm()
		this.database.Insert(&model.Tag{
			0,
			request.Form.Get("Name"),
		})
		http.Redirect(writer, request, tagBasePath, http.StatusTemporaryRedirect)
	}
}

func (this *tagController) Edit(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" && request.URL.Query().Has("Id") {
		id := request.URL.Query().Get("Id")
		var tags []model.Tag
		this.database.Query(&tags, model.TagQuery, "WHERE Id = '"+id+"'")
		tag := tags[0]
		render(writer, request, "tagForm", tag)
	} else if request.Method == "POST" {
		request.ParseForm()
		id, _ := strconv.Atoi(request.Form.Get("Id"))
		this.database.Update(&model.Tag{
			id,
			request.Form.Get("Name"),
		})
		http.Redirect(writer, request, tagBasePath, http.StatusTemporaryRedirect)
	}
}

func (this *tagController) Delete(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" && request.URL.Query().Has("Id") {
		id := request.URL.Query().Get("Id")
		this.database.Delete("Tag", id)
		http.Redirect(writer, request, tagBasePath, http.StatusTemporaryRedirect)
	}
}
