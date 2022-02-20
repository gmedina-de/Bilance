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
	repository repositories.Repository[T]
	route      string
}

func Generic[T any](route string) *generic[T] {
	return &generic[T]{route: route}
}

func (this *generic[T]) Routes() map[string]string {
	return map[string]string{
		this.route + "/":       "get,post:List(p);get:Edit(id);post:Save(id)",
		this.route + "/delete": "get:Remove(id)",
	}
}

const PageSize = 10
const PageSizeAll = 9223372036854775807

func (this *generic[T]) List(p string) {
	pageSize := PageSize
	if p == "all" {
		pageSize = PageSizeAll
	}

	paginator := template.NewPaginator(this.Request, pageSize, this.repository.Count())
	this.Data["Model"] = this.repository.Limit(pageSize, paginator.Offset())
	this.Data["Paginator"] = paginator
	this.Data["Title"] = models.Plural(this.repository.Model())
	this.TemplateName = "generic_list.gohtml"
}

func (this *generic[T]) Edit(id int64) {
	if id == 0 {
		this.Data["Form"] = this.repository.Model()
		this.Data["Title"] = "new"
	} else {
		this.Data["Form"] = this.repository.Find(id)
		this.Data["Title"] = "edit"
	}
	this.TemplateName = "generic_edit.gohtml"
}

func (this *generic[T]) Save(id int64) {
	record := this.repository.Model()
	this.ParseForm(record)
	if id == 0 {
		this.repository.Insert(record)
	} else {
		reflect.ValueOf(record).Elem().FieldByName("Id").SetInt(id)
		this.repository.Update(record)
	}
	this.Redirect(this.route, http.StatusTemporaryRedirect)
}

func (this *generic[T]) Remove(id int64) {
	item := this.repository.Find(id)
	this.repository.Delete(item)
	this.Redirect(this.route, http.StatusTemporaryRedirect)
}
