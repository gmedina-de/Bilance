package controllers

import (
	"genuine/core/http"
	"genuine/core/models"
	"genuine/core/repositories"
	"genuine/core/template"
	"reflect"
	"strconv"
)

type generic[T any] struct {
	repository repositories.Repository[T]
	route      string
}

func Generic[T any](repository repositories.Repository[T], route string) *generic[T] {
	return &generic[T]{repository, route}
}

func (g *generic[T]) Routes() map[string]http.Handler {
	return map[string]http.Handler{
		"GET " + g.route:             g.List,
		"POST " + g.route:            g.List,
		"GET " + g.route + "/new":    g.New,
		"GET " + g.route + "/edit":   g.Edit,
		"POST " + g.route + "/edit":  g.Save,
		"GET " + g.route + "/delete": g.Remove,
	}
}

const PageSize = 10
const PageSizeAll = 9223372036854775807

func (g *generic[T]) List(r http.Request) http.Response {
	p := r.URL.Query().Get("p")

	pageSize := PageSize
	if p == "all" {
		pageSize = PageSizeAll
	}

	paginator := template.NewPaginator(r, pageSize, g.repository.Count())

	return http.Response{
		"Model":     g.repository.Limit(pageSize, paginator.Offset()),
		"Paginator": paginator,
		"Title":     models.Name(g.repository.T()),
		"Template":  "generic_list",
	}
}

func (g *generic[T]) New(http.Request) http.Response {
	return http.Response{
		"Form":     g.repository.T(),
		"Title":    "new",
		"Template": "generic_edit",
	}
}

func (g *generic[T]) Edit(r http.Request) http.Response {
	id, _ := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	return http.Response{
		"Form":     g.repository.Find(id),
		"Title":    "edit",
		"Template": "generic_edit",
	}
}

func (g *generic[T]) Save(r http.Request) http.Response {
	id, _ := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	record := g.repository.T()
	//g.ParseForm(record)
	if id == 0 {
		g.repository.Insert(record)
	} else {
		reflect.ValueOf(record).Elem().FieldByName("Id").SetInt(id)
		g.repository.Update(record)
	}
	return http.Response{"Redirect": g.route}
}

func (g *generic[T]) Remove(r http.Request) http.Response {
	id, _ := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	item := g.repository.Find(id)
	g.repository.Delete(item)
	return http.Response{"Redirect": g.route}
}

func (g *generic[T]) ParseForm(model any) {
	//b.Request.ParseForm()
	//err := decoder.Decode(model, b.Request.PostForm)
	//if err != nil {
	//	err.Error()
	//	println(err.Error())
	//	// Handle error
	//}
}
