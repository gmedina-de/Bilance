package controllers

import (
	"genuine/core/http"
	"genuine/core/models"
	"genuine/core/repositories"
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

func (g *generic[T]) List(r http.Request) http.Response {
	page, err := strconv.ParseInt(r.URL.Query().Get("page"), 10, 64)
	if err != nil {
		page = 1
	}
	var pageSize int64 = 10
	var model []T
	var pages = g.repository.Count() / pageSize
	var offset = pageSize * (page - 1)
	if page == 0 {
		model = g.repository.All()
	} else {
		pages++
		model = g.repository.Limit(pageSize, offset)
	}
	return http.Response{
		"Model":    model,
		"Title":    models.Plural(g.repository.T()),
		"Pages":    pages,
		"Page":     page,
		"Template": "generic_list",
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
	//g.parseForm(record)
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

func (g *generic[T]) parseForm(model any) {
	//b.Request.parseForm()
	//err := decoder.Decode(model, b.Request.PostForm)
	//if err != nil {
	//	err.Error()
	//	println(err.Error())
	//	// Handle error
	//}
}
