package controller

import (
	"Bilance/model"
	"Bilance/repository"
	"net/http"
	"strconv"
)

type baseController struct {
	repository   repository.Repository
	basePath     string
	dataProvider func(request *http.Request) interface{}
}

func (c *baseController) List(writer http.ResponseWriter, request *http.Request) {
	var toast string
	if request.URL.Query().Has("success") {
		toast = "record_saved_successfully"
	}
	var list interface{}
	modelName := c.repository.ModelNamePlural()
	switch modelName {
	case "payments":
		fallthrough
	case "categories":
		list = c.repository.List("WHERE ProjectId = " + model.GetSelectedProjectIdString(request))
	default:
		list = c.repository.List()
	}
	render(writer, request, &Parameters{Model: list, Toast: toast}, modelName, "crud_table", modelName)
}

func (c *baseController) Edit(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		var data interface{}
		if c.dataProvider != nil {
			data = c.dataProvider(request)
		}
		modelName := c.repository.ModelNamePlural()
		if request.URL.Query().Has("Id") {
			idString := request.URL.Query().Get("Id")
			id, _ := strconv.ParseInt(idString, 10, 64)
			model := c.repository.Find(id)
			render(writer, request, &Parameters{Model: model, Data: data}, "edit", "crud_form", modelName)
		} else {
			render(writer, request, &Parameters{Model: c.repository.NewEmpty(), Data: data}, "new", "crud_form", modelName)
		}
	} else if request.Method == "POST" {
		err := request.ParseForm()
		if err != nil {
			panic(err)
		}
		if request.Form.Has("Id") {
			id, _ := strconv.ParseInt(request.Form.Get("Id"), 10, 64)
			c.repository.Update(c.repository.NewFromRequest(request, id))
		} else {
			c.repository.Insert(c.repository.NewFromRequest(request, 0))
		}
		redirect(writer, request, c.basePath+"?success")
	}
}

func (c *baseController) Delete(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" && request.URL.Query().Has("Id") {
		id, _ := strconv.ParseInt(request.URL.Query().Get("Id"), 10, 64)
		item := c.repository.Find(id)
		c.repository.Delete(item)
		redirect(writer, request, c.basePath)
	}
}
