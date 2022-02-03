package controller

import (
	"homecloud/core/model"
	"homecloud/core/repository"
	"homecloud/core/server"
	"net/http"
	"strconv"
	"strings"
)

type Generic[T model.Model] struct {
	BaseTemplate string
	Repository   repository.Repository[T]
	DataProvider func(request *http.Request) interface{}
}

func (g *Generic[T]) Routing(server server.Server) {
	server.Get("", g.Index)
	server.Post("", g.Index)
	server.Get("/edit", g.Edit)
	server.Post("/edit", g.Edit)
	server.Get("/edit/delete", g.Delete)
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

	//projectIdString := GetSelectedProjectIdString(request)
	projectIdString := "0"
	modelName := g.Repository.ModelNamePlural()
	var conditions []string
	switch modelName {
	case "categories":
		conditions = append(conditions, "ProjectId = "+projectIdString)
	case "payments":
		conditions = append(conditions, "ProjectId = "+projectIdString)
		conditions = append(conditions, "ORDER BY Date DESC")
	}

	pagination, _ := g.handlePagination(request, projectIdString)
	parameters := &Parameters{
		Model: g.Repository.All(),
		//Model:      g.repository.Raw(strings.Join(conditions, " AND ") + " " + limitCondition),
		Data:       data,
		Pagination: pagination,
		Toast:      toast,
	}
	Render(
		writer,
		request,
		parameters,
		modelName,
		"core/template/crud_table.gohtml",
		g.BaseTemplate,
	)
}

func (g *Generic[T]) handlePagination(request *http.Request, projectIdString string) (*Pagination, string) {
	if strings.HasPrefix(request.URL.Path, "/admin/") {
		return nil, ""
	}
	var limit int64 = 10
	var page int64 = 1
	if request.URL.Query().Get("page") != "" {
		page, _ = strconv.ParseInt(request.URL.Query().Get("page"), 10, 64)
	}
	var offset = limit * (page - 1)
	var pages = g.Repository.Count("ProjectId = "+projectIdString) / limit
	pages++
	return &Pagination{pages, page},
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
			Render(writer, request, &Parameters{Model: model, Data: data}, "edit", "core/template/crud_table.gohtml", g.BaseTemplate)
		} else {
			Render(writer, request, &Parameters{Model: g.Repository.NewEmpty(), Data: data}, "new", "core/template/crud_form.gohtml", g.BaseTemplate)
		}
	} else if request.Method == "POST" {
		err := request.ParseForm()
		if err != nil {
			panic(err)
		}
		if request.Form.Get("Id") != "" {
			id, _ := strconv.ParseInt(request.Form.Get("Id"), 10, 64)
			g.Repository.Update(g.Repository.FromRequest(request, id))
		} else {
			g.Repository.Insert(g.Repository.FromRequest(request, 0))
		}
		Redirect(writer, request, strings.Replace(request.URL.Path, "/edit", "", 1)+"?success")
	}
}

func (g *Generic[T]) Delete(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" && request.URL.Query().Get("Id") != "" {
		id, _ := strconv.ParseInt(request.URL.Query().Get("Id"), 10, 64)
		item := g.Repository.Find(id)
		g.Repository.Delete(item)
		Redirect(writer, request, strings.Replace(request.URL.Path, "/edit/delete", "", 1))
	}
}
