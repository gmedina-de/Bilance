package controller

import (
	"Bilance/model"
	"Bilance/repository"
	"net/http"
	"reflect"
	"strconv"
	"strings"
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
	if c.modelNamePlural() == "categories" || c.modelNamePlural() == "payments" {
		cookie, _ := request.Cookie(model.SelectedProjectIdCookie)
		list = c.repository.List("WHERE ProjectId = " + cookie.Value)
	} else {
		list = c.repository.List()
	}
	render(writer, request, &Parameters{Model: list, Toast: toast}, c.modelNamePlural(), "crud_table", c.modelNamePlural())
}

func (c *baseController) Edit(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		var data interface{}
		if c.dataProvider != nil {
			data = c.dataProvider(request)
		}
		if request.URL.Query().Has("Id") {
			idString := request.URL.Query().Get("Id")
			id, _ := strconv.ParseInt(idString, 10, 64)
			model := c.repository.Find(id)
			render(writer, request, &Parameters{Model: model, Data: data}, "edit", "crud_form", c.modelNamePlural())
		} else {
			render(writer, request, &Parameters{Model: c.repository.NewEmpty(), Data: data}, "new", "crud_form", c.modelNamePlural())
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

func (c *baseController) modelNamePlural() string {
	empty := c.repository.NewEmpty()
	of := reflect.TypeOf(empty).Elem()
	name := of.Name()
	lower := strings.ToLower(name)
	if strings.HasSuffix(lower, "y") {
		return lower[:len(lower)-1] + "ies"
	}
	return lower + "s"
}
