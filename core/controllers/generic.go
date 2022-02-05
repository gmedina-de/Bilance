package controllers

import (
	"github.com/gorilla/schema"
	"homecloud/core/repository"
	"homecloud/core/server"
	"homecloud/core/template"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type Generic[T any] struct {
	BaseController
	Model        T
	Repository   repository.Repository[T]
	BasePath     string
	DataProvider func(request *http.Request) interface{}
}

func (this *Generic[T]) Route() string {
	this.ViewPath = "core/views"
	return this.BasePath
}

func (this *Generic[T]) Routing(server server.Server) {

	server.Get(this.BasePath+"/edit", this.Edit)
	server.Post(this.BasePath+"/edit", this.Edit)
	server.Get(this.BasePath+"/edit/delete", this.Remove)
}

func (this *Generic[T]) Get() {
	var data interface{}
	//if this.DataProvider != nil {
	//	data = this.DataProvider(request)
	//}

	var toast string
	if this.Ctx.Request.URL.Query().Get("success") != "" {
		toast = "record_saved_successfully"
	}

	// todo use framework pagination
	limit, offset, pagination := template.HandlePagination(this.Ctx.Request, this.Repository.Count())
	this.Data["Model"] = this.Repository.Limit(limit, offset)
	this.Data["Data"] = data
	this.Data["Pagination"] = pagination
	this.Data["Toast"] = toast
	this.TplName = "generic.gohtml"
}

func (this *Generic[T]) Edit(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		var data interface{}
		if this.DataProvider != nil {
			data = this.DataProvider(request)
		}
		if request.URL.Query().Get("Id") != "" {
			idString := request.URL.Query().Get("Id")
			id, _ := strconv.ParseInt(idString, 10, 64)
			model := this.Repository.Find(id)
			template.Render(writer, request, "edit", &template.Parameters{Model: model, Data: data}, "core/template/form.gohtml")
		} else {
			template.Render(writer, request, "new", &template.Parameters{Model: this.Model, Data: data}, "core/template/form.gohtml")
		}
	} else if request.Method == "POST" {
		err := request.ParseForm()
		if err != nil {
			panic(err)
		}
		if request.Form.Get("Id") != "" {
			this.Repository.Update(this.FromRequest(request))
		} else {
			this.Repository.Insert(this.FromRequest(request))
		}
		http.Redirect(writer, request, strings.Replace(request.URL.Path, "/edit", "", 1)+"?success", http.StatusTemporaryRedirect)
	}
}

func (this *Generic[T]) Remove(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" && request.URL.Query().Get("Id") != "" {
		id, _ := strconv.ParseInt(request.URL.Query().Get("Id"), 10, 64)
		item := this.Repository.Find(id)
		this.Repository.Delete(item)
		http.Redirect(writer, request, strings.Replace(request.URL.Path, "/edit/delete", "", 1), http.StatusTemporaryRedirect)
	}
}

var decoder = schema.NewDecoder()

func (this *Generic[T]) FromRequest(request *http.Request) any {
	request.ParseForm()

	result := reflect.New(reflect.TypeOf(this.Model)).Interface()
	err := decoder.Decode(result, request.PostForm)
	if err != nil {
		err.Error()
		println(err.Error())
		// Handle error
	}
	return result
}
