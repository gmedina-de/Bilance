package controllers

import (
	"genuine/core/controllers"
	"genuine/core/models"
	"genuine/core/repositories"
	"genuine/core/router"
	"strconv"
	"strings"

	"github.com/gorilla/schema"
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
		"GET " + g.route:                  g.List,
		"POST " + g.route:                 g.List,
		"GET " + g.route + "/new":         g.New,
		"POST " + g.route + "/new":        g.Save,
		"GET " + g.route + "/edit":        g.Edit,
		"POST " + g.route + "/edit":       g.Save,
		"GET " + g.route + "/edit/delete": g.Remove,
	}
}

func (g *generic[T]) List(r controllers.Request) controllers.Response {
	title := models.Plural(g.repository.Model())

	// HANDLE SEARCH
	where := ""
	var args []any
	if r.URL.Query().Has("q") {
		var wheres []string
		search := r.URL.Query().Get("q")
		commas := strings.Split(search, ",")
		for _, comma := range commas {
			colons := strings.Split(comma, ":")
			wheres = append(wheres, colons[0]+" LIKE ?")
			args = append(args, colons[1])
		}
		where = strings.Join(wheres, " AND ")
		title = "search_results"
	}

	// HANDLE PAGINATION
	page, err := strconv.ParseInt(r.URL.Query().Get("p"), 10, 64)
	if err != nil {
		page = 1
	}
	var pageSize int64 = 10
	var model []T
	var pages = g.repository.Count(where, args...) / pageSize
	var offset = pageSize * (page - 1)
	if page == 0 {
		model = g.repository.List(where, args...)
	} else {
		pages++
		model = g.repository.Limit(int(pageSize), int(offset), where, args...)
	}

	// RENDER
	return controllers.Response{
		"Model":    model,
		"Title":    title,
		"Template": "generic_list",
		"Pages":    pages,
		"Page":     page,
	}
}

func (g *generic[T]) New(controllers.Request) controllers.Response {
	return controllers.Response{
		"Model":    g.repository.Model(),
		"Title":    "new",
		"Template": "generic_edit",
	}
}

func (g *generic[T]) Edit(r controllers.Request) controllers.Response {
	id := uint(g.getID(r))
	return controllers.Response{
		"Model":    g.repository.Find(id),
		"Title":    "edit",
		"Template": "generic_edit",
	}
}

var decoder = schema.NewDecoder()

func (g *generic[T]) Save(r controllers.Request) controllers.Response {
	id := g.getID(r)
	model := g.repository.Model()

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	err = decoder.Decode(&model, r.PostForm)
	if err != nil {
		panic(err)
	}

	if id == 0 {
		g.repository.Insert(&model)
	} else {
		models.RealValueOf(&model).FieldByName(models.ID).SetUint(id)
		g.repository.Update(&model)
	}
	return router.Redirect(g.route)(r)
}

func (g *generic[T]) Remove(r controllers.Request) controllers.Response {
	item := g.repository.Find(uint(g.getID(r)))
	g.repository.Delete(item)
	return router.Redirect(g.route)(r)
}

func (*generic[T]) getID(r controllers.Request) uint64 {
	id, _ := strconv.ParseUint(r.URL.Query().Get(models.ID), 10, 64)
	return id
}
