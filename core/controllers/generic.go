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

func (g *generic[T]) Routes() map[string]string {
	return map[string]string{
		g.route:             "get,post:List(p)",
		g.route + "/new":    "get:New()",
		g.route + "/edit":   "get:Edit(id);post:Save(id)",
		g.route + "/delete": "get:Remove(id)",
	}
}

const PageSize = 10
const PageSizeAll = 9223372036854775807

func (g *generic[T]) List(p string) {
	pageSize := PageSize
	if p == "all" {
		pageSize = PageSizeAll
	}

	paginator := template.NewPaginator(g.Request, pageSize, g.repository.Count())
	g.Data["Model"] = g.repository.Limit(pageSize, paginator.Offset())
	g.Data["Paginator"] = paginator
	g.Data["Title"] = models.Name(g.repository.T())
	g.TemplateName = "generic_list.gohtml"
}

func (g *generic[T]) New() {
	g.Data["Form"] = g.repository.T()
	g.Data["Title"] = "new"
	g.TemplateName = "generic_edit.gohtml"
}

func (g *generic[T]) Edit(id int64) {
	g.Data["Form"] = g.repository.Find(id)
	g.Data["Title"] = "edit"
	g.TemplateName = "generic_edit.gohtml"
}

func (g *generic[T]) Save(id int64) {
	record := g.repository.T()
	g.ParseForm(record)
	if id == 0 {
		g.repository.Insert(record)
	} else {
		reflect.ValueOf(record).Elem().FieldByName("Id").SetInt(id)
		g.repository.Update(record)
	}
	g.Redirect(g.route, http.StatusTemporaryRedirect)
}

func (g *generic[T]) Remove(id int64) {
	item := g.repository.Find(id)
	g.repository.Delete(item)
	g.Redirect(g.route, http.StatusTemporaryRedirect)
}
