package controllers

import (
	"genuine/core/models"
	"genuine/core/repositories"
	"genuine/core/template"
	"reflect"
	"strconv"
)

type generic[T any] struct {
	repository repositories.Repository[T]
	route      string
	listView   string
	editView   string
}

func Generic[T any](repository repositories.Repository[T], route string) *generic[T] {
	return &generic[T]{repository, route}
}

func (g *generic[T]) Routes() map[string]Handler {
	return map[string]Handler{
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

func (g *generic[T]) List(r Request) Response {
	p := r.URL.Query().Get("p")

	pageSize := PageSize
	if p == "all" {
		pageSize = PageSizeAll
	}

	paginator := template.NewPaginator(r, pageSize, g.repository.Count())

	return Response{
		"Model":     g.repository.Limit(pageSize, paginator.Offset()),
		"Paginator": paginator,
		"Title":     models.Name(g.repository.T()),
		"Template":  g.listView,
	}
}

func (g *generic[T]) New(Request) Response {
	return Response{
		"Form":     g.repository.T(),
		"Title":    "new",
		"Template": g.editView,
	}
}

func (g *generic[T]) Edit(r Request) Response {
	id, _ := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	return Response{
		"Form":     g.repository.Find(id),
		"Title":    "edit",
		"Template": g.editView,
	}
}

func (g *generic[T]) Save(r Request) Response {
	id, _ := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	record := g.repository.T()
	//g.ParseForm(record)
	if id == 0 {
		g.repository.Insert(record)
	} else {
		reflect.ValueOf(record).Elem().FieldByName("Id").SetInt(id)
		g.repository.Update(record)
	}
	return Response{"Redirect": g.route}
}

func (g *generic[T]) Remove(r Request) Response {
	id, _ := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	item := g.repository.Find(id)
	g.repository.Delete(item)
	return Response{"Redirect": g.route}
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
