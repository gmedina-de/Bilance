package controller

import (
	"Bilance/repository"
	"Bilance/service/router"
	"net/http"
	"strconv"
)

type baseController struct {
	repository repository.Repository
	basePath   string
	template   string
}

func (c *baseController) Routing(router router.Router) {
	router.Get(c.basePath, c.List)
	router.Post(c.basePath, c.List)
	router.Get(c.basePath+"/new", c.New)
	router.Post(c.basePath+"/new", c.New)
	router.Get(c.basePath+"/edit", c.Edit)
	router.Post(c.basePath+"/edit", c.Edit)
	router.Get(c.basePath+"/delete", c.Delete)
}

func (c *baseController) List(writer http.ResponseWriter, request *http.Request) {
	var search interface{}
	if request.URL.Query().Has("Search") {
		search = c.repository.List("WHERE Name LIKE '%" + request.URL.Query().Get("Search") + "%'")
	} else {
		search = c.repository.List()
	}
	render(writer, request, c.template, search)
}

func (c *baseController) New(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		render(writer, request, c.template+"Form", c.repository.NewEmpty())
	} else if request.Method == "POST" {
		request.ParseForm()
		c.repository.Insert(c.repository.NewFromRequest(request, 0))
		http.Redirect(writer, request, c.basePath, http.StatusTemporaryRedirect)
	}
}

func (c *baseController) Edit(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" && request.URL.Query().Has("Id") {
		id := request.URL.Query().Get("Id")
		item := c.repository.List("WHERE Id = " + id)
		render(writer, request, c.template+"Form", item)
	} else if request.Method == "POST" {
		request.ParseForm()
		id, _ := strconv.Atoi(request.Form.Get("Id"))
		c.repository.Update(c.repository.NewFromRequest(request, id))
		http.Redirect(writer, request, c.basePath, http.StatusTemporaryRedirect)
	}
}

func (c *baseController) Delete(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" && request.URL.Query().Has("Id") {
		id := request.URL.Query().Get("Id")
		item := c.repository.List("WHERE Id = " + id)
		c.repository.Delete(item)
		http.Redirect(writer, request, c.basePath, http.StatusTemporaryRedirect)
	}
}
