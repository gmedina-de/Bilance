package controller

import (
	"Bilance/repository"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type baseController struct {
	repository repository.Repository
	basePath   string
}

func (c *baseController) List(writer http.ResponseWriter, request *http.Request) {
	var toast string
	if request.URL.Query().Has("success") {
		toast = "record_saved_successfully"
	}
	render(writer, request, &Parameters{Model: c.repository.List(), Toast: toast}, c.modelName(), "crud_table", c.modelName())
}

func (c *baseController) Edit(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		if request.URL.Query().Has("Id") {
			idString := request.URL.Query().Get("Id")
			id, _ := strconv.ParseInt(idString, 10, 64)
			item := c.repository.Find(id)
			render(writer, request, &Parameters{Model: item}, "edit", "crud_form", c.modelName())
		} else {
			render(writer, request, &Parameters{Model: c.repository.NewEmpty()}, "new", "crud_form", c.modelName())
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

func (c *baseController) modelName() string {
	empty := c.repository.NewEmpty()
	of := reflect.TypeOf(empty).Elem()
	name := of.Name()
	lower := strings.ToLower(name)
	if strings.HasSuffix(lower, "y") {
		return lower[:len(lower)-1] + "ies"
	}
	return lower + "s"
}
