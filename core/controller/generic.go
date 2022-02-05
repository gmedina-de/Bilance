package controller

import (
	"github.com/gorilla/schema"
	"homecloud/core/model"
	"homecloud/core/repository"
	"homecloud/core/server"
	"homecloud/core/template"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type Generic[T any] struct {
	Model        T
	Repository   repository.Repository[T]
	BasePath     string
	DataProvider func(request *http.Request) interface{}
}

func (g *Generic[T]) Routing(server server.Server) {
	server.Get(g.BasePath+"", g.Index)
	server.Post(g.BasePath+"", g.Index)
	server.Get(g.BasePath+"/edit", g.Edit)
	server.Post(g.BasePath+"/edit", g.Edit)
	server.Get(g.BasePath+"/edit/delete", g.Delete)
}

func (g *Generic[T]) Index(writer http.ResponseWriter, request *http.Request) {

	var data interface{}
	if g.DataProvider != nil {
		data = g.DataProvider(request)
	}

	var toast string
	if request.URL.Query().Get("success") != "" {
		toast = "record_saved_successfully"
	}

	modelName := model.Plural(g.Model)
	pagination, _ := g.handlePagination(request)
	all := g.Repository.All()
	parameters := &template.Parameters{
		Model: all,
		//Model:      g.repository.Raw(strings.Join(conditions, " AND ") + " " + limitCondition),
		Data:       data,
		Pagination: pagination,
		Toast:      toast,
	}
	template.Render(
		writer,
		request,
		parameters,
		modelName,
		"core/template/table.gohtml",
	)
}

func (g *Generic[T]) handlePagination(request *http.Request) (*template.Pagination, string) {
	if strings.HasPrefix(request.URL.Path, "/admin/") {
		return nil, ""
	}
	var limit int64 = 10
	var page int64 = 1
	if request.URL.Query().Get("page") != "" {
		page, _ = strconv.ParseInt(request.URL.Query().Get("page"), 10, 64)
	}
	var offset = limit * (page - 1)
	//var pages = g.Repository.Count("") / limit
	var pages int64 = 10
	pages++
	return &template.Pagination{pages, page},
		"LIMIT " + strconv.FormatInt(limit, 10) + " OFFSET " + strconv.FormatInt(offset, 10)
}

func (g *Generic[T]) Edit(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		var data interface{}
		if g.DataProvider != nil {
			data = g.DataProvider(request)
		}
		if request.URL.Query().Get("Id") != "" {
			idString := request.URL.Query().Get("Id")
			id, _ := strconv.ParseInt(idString, 10, 64)
			model := g.Repository.Find(id)
			template.Render(writer, request, &template.Parameters{Model: model, Data: data}, "edit", "core/template/form.gohtml")
		} else {
			template.Render(writer, request, &template.Parameters{Model: g.Model, Data: data}, "new", "core/template/form.gohtml")
		}
	} else if request.Method == "POST" {
		err := request.ParseForm()
		if err != nil {
			panic(err)
		}
		if request.Form.Get("Id") != "" {
			g.Repository.Update(g.FromRequest(request))
		} else {
			g.Repository.Insert(g.FromRequest(request))
		}
		http.Redirect(writer, request, strings.Replace(request.URL.Path, "/edit", "", 1)+"?success", http.StatusTemporaryRedirect)
	}
}

func (g *Generic[T]) Delete(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" && request.URL.Query().Get("Id") != "" {
		id, _ := strconv.ParseInt(request.URL.Query().Get("Id"), 10, 64)
		item := g.Repository.Find(id)
		g.Repository.Delete(item)
		http.Redirect(writer, request, strings.Replace(request.URL.Path, "/edit/delete", "", 1), http.StatusTemporaryRedirect)
	}
}

var decoder = schema.NewDecoder()

func (g *Generic[T]) FromRequest(request *http.Request) any {
	request.ParseForm()

	result := reflect.New(reflect.TypeOf(g.Model)).Interface()
	err := decoder.Decode(result, request.PostForm)
	if err != nil {
		err.Error()
		println(err.Error())
		// Handle error
	}
	return result
}
