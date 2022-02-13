package controllers

import (
	"genuine/core/models"
	"genuine/core/repositories"
	"genuine/core/router"
	"github.com/beego/beego/v2/server/web/pagination"
	"net/http"
	"reflect"
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
	router.Add(this, this.Route+"/:id/delete", "get:Remove(id path)")
}

const PageSize = 10
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

func (this *generic[T]) Edit(id int64) {
	if id == 0 {
		this.Data["Form"] = &this.Model
		this.Data["Title"] = "new"
	} else {
		this.Data["Form"] = this.Repository.Find(id)
		this.Data["Title"] = "edit"
	}
	this.TplName = "generic_edit.gohtml"
}

func (this *generic[T]) Save(id int64) {
	record := &this.Model
	if err := this.ParseForm(record); err != nil {
		panic(err)
	}
	if id == 0 {
		this.Repository.Insert(record)
	} else {
		reflect.ValueOf(record).Elem().FieldByName("Id").SetInt(id)
		this.Repository.Update(record)
	}
	this.Redirect(this.Route, http.StatusTemporaryRedirect)
}

func (this *generic[T]) Remove(id int64) {
	item := this.Repository.Find(id)
	this.Repository.Delete(item)
	this.Redirect(this.Route, http.StatusTemporaryRedirect)
}
