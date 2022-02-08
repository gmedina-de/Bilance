package controllers

import (
	"genuine/framework/models"
	"genuine/framework/repositories"
	"genuine/framework/router"
	"github.com/beego/beego/v2/server/web/pagination"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type generic[T any] struct {
	BaseController
	Repository repositories.Repository[T]
	Model      T
	Route      string
}

func Generic[T any](repository repositories.Repository[T], model T, route string) *generic[T] {
	return &generic[T]{Repository: repository, Model: model, Route: route}
}

func (this *generic[T]) Routing() {
	router.Add(this, this.Route, "get,post:List(p)")
	router.Add(this, this.Route+"/:id", "get:Edit(id path);post:Save(id path)")
}

const PageSize = 5
const PageSizeAll = 9223372036854775807

func (this *generic[T]) List(p string) {
	pageSize := PageSize
	if p == "all" {
		pageSize = PageSizeAll
	}
	paginator := pagination.SetPaginator(this.Ctx, pageSize, this.Repository.Count())
	this.Data["Model"] = this.Repository.Limit(pageSize, paginator.Offset())
	this.Data["Paginator"] = paginator
	this.Data["Title"] = models.Plural(this.Model)
	this.TplName = "generic_list.gohtml"
}

func (this *generic[T]) Edit(id string) {
	if id == "new" {
		this.Data["Form"] = &this.Model
		this.Data["Title"] = "new"
	} else {
		id, _ := strconv.ParseInt(id, 10, 64)
		this.Data["Form"] = this.Repository.Find(id)
		this.Data["Title"] = "edit"
	}
	this.TplName = "generic_edit.gohtml"
}

func (this *generic[T]) Save(id string) {
	record := &this.Model
	if err := this.ParseForm(record); err != nil {
		panic(err)
	}
	if id == "new" {
		this.Repository.Insert(record)
	} else {
		id, _ := strconv.ParseInt(id, 10, 64)
		reflect.ValueOf(record).Elem().FieldByName("Id").SetInt(id)
		this.Repository.Update(record)
	}
	this.Redirect(this.Route, http.StatusTemporaryRedirect)
}

func (this *generic[T]) Remove(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" && request.URL.Query().Get("Id") != "" {
		id, _ := strconv.ParseInt(request.URL.Query().Get("Id"), 10, 64)
		item := this.Repository.Find(id)
		this.Repository.Delete(item)
		http.Redirect(writer, request, strings.Replace(request.URL.Path, "/edit/delete", "", 1), http.StatusTemporaryRedirect)
	}
}
