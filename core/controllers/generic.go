package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"homecloud/core/models"
	"homecloud/core/repositories"
	"homecloud/core/template"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type Generic[T any] struct {
	BaseController
	Repository repositories.Repository[T]
	Model      T
	Route      string
}

func Generics[T any](repository repositories.Repository[T], model T, route string) *Generic[T] {
	return &Generic[T]{Repository: repository, Model: model, Route: route}
}

func (this *Generic[T]) Routing() {
	web.Router(this.Route, this, "get:List;post:List")
	web.Router(this.Route+"/:id", this, "get:Edit;post:Save")
}

func (this *Generic[T]) List() {
	var toast string
	if this.Ctx.Request.URL.Query().Get("success") != "" {
		toast = "record_saved_successfully"
	}
	//todo use framework pagination
	limit, offset, pagination := template.HandlePagination(this.Ctx.Request, this.Repository.Count())
	i := this.Repository.Limit(limit, offset)
	this.Data["Model"] = i
	this.Data["Pagination"] = pagination
	this.Data["Toast"] = toast
	this.Data["Title"] = models.Plural(this.Model)
	this.TplName = "generic.gohtml"
}

func (this *Generic[T]) Edit() {
	param := this.Ctx.Input.Param(":id")
	if param == "new" {
		this.Data["Form"] = this.Model
		this.Data["Title"] = "new"
	} else {
		id, _ := strconv.ParseInt(param, 10, 64)
		this.Data["Form"] = this.Repository.Find(id)
		this.Data["Title"] = "edit"
	}
	this.TplName = "form.gohtml"
}

func (this *Generic[T]) Save() {
	param := this.Ctx.Input.Param(":id")
	record := &this.Model
	if err := this.ParseForm(record); err != nil {
		panic(err)
	}
	if param == "new" {
		this.Repository.Insert(record)
	} else {
		id, _ := strconv.ParseInt(param, 10, 64)
		reflect.ValueOf(record).Elem().FieldByName("Id").SetInt(id)
		this.Repository.Update(record)
	}
	this.Redirect(this.Route, http.StatusTemporaryRedirect)
}

func (this *Generic[T]) Remove(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" && request.URL.Query().Get("Id") != "" {
		id, _ := strconv.ParseInt(request.URL.Query().Get("Id"), 10, 64)
		item := this.Repository.Find(id)
		this.Repository.Delete(item)
		http.Redirect(writer, request, strings.Replace(request.URL.Path, "/edit/delete", "", 1), http.StatusTemporaryRedirect)
	}
}
