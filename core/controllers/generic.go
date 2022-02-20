package controllers

import (
	"genuine/core/models"
	"genuine/core/repositories"
	"genuine/core/template"
	"net/http"
	"reflect"
)

type generic[T any] struct {
	*Base
	Repository repositories.Repository[T]
	Model      T
	Route      string
}

func Generic[T any](model T, route string) *generic[T] {
	return &generic[T]{Model: model, Route: route}
}

func (this *generic[T]) Routes() map[string]string {
	return map[string]string{
		this.Route + "/":       "get,post:List(p);get:Edit(id);post:Save(id)",
		this.Route + "/delete": "get:Remove(id)",
	}
}

const PageSize = 10
const PageSizeAll = 9223372036854775807

func (this *generic[T]) List(p string) {
	pageSize := PageSize
	if p == "all" {
		pageSize = PageSizeAll
	}

	paginator := template.NewPaginator(this.Request, pageSize, this.Repository.Count())
	this.Data["Model"] = this.Repository.Limit(pageSize, paginator.Offset())
	this.Data["Paginator"] = paginator
	this.Data["Title"] = models.Plural(this.Model)
	this.TemplateName = "generic_list.gohtml"
}

func (this *generic[T]) Edit(id int64) {
	if id == 0 {
		this.Data["Form"] = &this.Model
		this.Data["Title"] = "new"
	} else {
		this.Data["Form"] = this.Repository.Find(id)
		this.Data["Title"] = "edit"
	}
	this.TemplateName = "generic_edit.gohtml"
}

func (this *generic[T]) Save(id int64) {
	record := &this.Model
	this.ParseForm(record)
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
