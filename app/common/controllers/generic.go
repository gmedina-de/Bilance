package controllers

import (
	"genuine/core/controllers"
	"genuine/core/models"
	"genuine/core/repositories"
	"genuine/core/router"
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

func (g *generic[T]) Routes() map[string]controllers.Handler {
	return map[string]controllers.Handler{
		"GET " + g.route:             g.List,
		"POST " + g.route:            g.List,
		"GET " + g.route + "/new":    g.New,
		"GET " + g.route + "/edit":   g.Edit,
		"POST " + g.route + "/edit":  g.Save,
		"GET " + g.route + "/delete": g.Remove,
	}
}

func (g *generic[T]) List(r controllers.Request) controllers.Response {
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
	return controllers.Response{
		"Model":    model,
		"Title":    models.Plural(g.repository.Model()),
		"Pages":    pages,
		"Page":     page,
		"Template": "generic_list",
	}
}

func (g *generic[T]) New(controllers.Request) controllers.Response {
	return controllers.Response{
		"Form":     g.repository.Model(),
		"Title":    "new",
		"Template": "generic_edit",
	}
}

func (g *generic[T]) Edit(r controllers.Request) controllers.Response {
	id, _ := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	return controllers.Response{
		"Form":     g.repository.Find(id),
		"Title":    "edit",
		"Template": "generic_edit",
	}
}

func (g *generic[T]) Save(r controllers.Request) controllers.Response {
	id, _ := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	record := g.repository.Model()
	//g.parseForm(record)
	if id == 0 {
		g.repository.Insert(record)
	} else {
		reflect.ValueOf(record).Elem().FieldByName("Id").SetInt(id)
		g.repository.Update(record)
	}
	return router.Redirect(g.route)(r)
}

func (g *generic[T]) Remove(r controllers.Request) controllers.Response {
	id, _ := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	item := g.repository.Find(id)
	g.repository.Delete(item)
	return router.Redirect(g.route)(r)
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
