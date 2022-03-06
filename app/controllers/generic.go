package controllers

import (
	models2 "genuine/app/models"
	"genuine/core/controllers"
	"genuine/core/models"
	"genuine/core/repositories"
	"strconv"
	"strings"

	"github.com/gorilla/schema"
)

type generic[T any] struct {
	repository repositories.Repository[T]
	route      string
}

func Generic[T any](repository repositories.Repository[T], route string) *generic[T] {
	g := &generic[T]{repository, route}
	searchers = append(searchers, func(r controllers.Request) any {
		where, args := g.search(r)
		return repository.List(where, args)
	})
	return g
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
	title := models2.Plural(g.repository.Model())
	where, args := g.search(r)
	model, page, pages := g.paginate(r, where, args)
	return controllers.Response{
		"Title":    title,
		"Template": "generic_list",
		"Model":    model,
		"Pages":    pages,
		"Page":     page,
	}
}

func (g *generic[T]) search(r controllers.Request) (string, []any) {
	where := ""
	var args []any
	if r.URL.Query().Has("q") {
		var wheres []string
		search := r.URL.Query().Get("q")
		commas := strings.Split(search, ",")
		for _, comma := range commas {
			if strings.Contains(comma, ":") {
				colons := strings.Split(comma, ":")
				wheres = append(wheres, colons[0]+" LIKE ?")
				args = append(args, colons[1])
			} else {
				wheres = append(wheres, "NAME LIKE ?")
				args = append(args, "%"+comma+"%")
			}
		}
		where = strings.Join(wheres, " AND ")
	}
	return where, args
}

func (g *generic[T]) paginate(r controllers.Request, where string, args []any) (any, int64, int64) {
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
	return model, page, pages
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
		models2.RealValueOf(&model).FieldByName(models.ID).SetUint(id)
		g.repository.Update(&model)
	}
	return controllers.Redirect(g.route)(r)
}

func (g *generic[T]) Remove(r controllers.Request) controllers.Response {
	item := g.repository.Find(uint(g.getID(r)))
	g.repository.Delete(item)
	return controllers.Redirect(g.route)(r)
}

func (*generic[T]) getID(r controllers.Request) uint64 {
	id, _ := strconv.ParseUint(r.URL.Query().Get(models.ID), 10, 64)
	return id
}