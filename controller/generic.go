package controller

import (
	"Bilance/model"
	"Bilance/repository"
	"Bilance/server"
	"net/http"
	"strconv"
	"strings"
)

type generic[T model.Model] struct {
	repository   repository.Repository[T]
	basePath     string
	dataProvider func(request *http.Request) interface{}
}

func (g *generic[T]) Routing(server server.Server) {
	server.Get(g.basePath, g.List)
	server.Post(g.basePath, g.List)
	server.Get(g.basePath+"edit", g.Edit)
	server.Post(g.basePath+"edit", g.Edit)
	server.Get(g.basePath+"edit/delete", g.Delete)
}

func (g *generic[T]) List(writer http.ResponseWriter, request *http.Request) {

	var data interface{}
	if g.dataProvider != nil {
		data = g.dataProvider(request)
	}

	var toast string
	if request.URL.Query().Get("success") != "" {
		toast = "record_saved_successfully"
	}

	projectIdString := model.GetSelectedProjectIdString(request)
	modelName := g.repository.ModelNamePlural()
	var conditions []string
	switch modelName {
	case "categories":
		conditions = append(conditions, "ProjectId = "+projectIdString)
	case "payments":
		conditions = append(conditions, "ProjectId = "+projectIdString)
		conditions = append(conditions, "ORDER BY Date DESC")
	}

	pagination, limitCondition := g.handlePagination(request, projectIdString)
	parameters := &Parameters{
		Model:      g.repository.Raw(strings.Join(conditions, " AND ") + " " + limitCondition),
		Data:       data,
		Pagination: pagination,
		Toast:      toast,
	}
	render(
		writer,
		request,
		parameters,
		modelName,
		"crud_table",
		modelName,
	)
}

func (g *generic[T]) handlePagination(request *http.Request, projectIdString string) (*Pagination, string) {
	if strings.HasPrefix(request.URL.Path, "/admin/") {
		return nil, ""
	}
	var limit int64 = 10
	var page int64 = 1
	if request.URL.Query().Get("page") != "" {
		page, _ = strconv.ParseInt(request.URL.Query().Get("page"), 10, 64)
	}
	var offset = limit * (page - 1)
	var pages = g.repository.Count("ProjectId = "+projectIdString) / limit
	pages++
	return &Pagination{pages, page},
		"LIMIT " + strconv.FormatInt(limit, 10) + " OFFSET " + strconv.FormatInt(offset, 10)
}

func (g *generic[T]) Edit(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		var data interface{}
		if g.dataProvider != nil {
			data = g.dataProvider(request)
		}
		modelName := g.repository.ModelNamePlural()
		if request.URL.Query().Get("Id") != "" {
			idString := request.URL.Query().Get("Id")
			id, _ := strconv.ParseInt(idString, 10, 64)
			model := g.repository.Find(id)
			render(writer, request, &Parameters{Model: model, Data: data}, "edit", "crud_form", modelName)
		} else {
			render(writer, request, &Parameters{Model: g.repository.NewEmpty(), Data: data}, "new", "crud_form", modelName)
		}
	} else if request.Method == "POST" {
		err := request.ParseForm()
		if err != nil {
			panic(err)
		}
		if request.Form.Get("Id") != "" {
			id, _ := strconv.ParseInt(request.Form.Get("Id"), 10, 64)
			g.repository.Update(g.repository.FromRequest(request, id))
		} else {
			g.repository.Insert(g.repository.FromRequest(request, 0))
		}
		redirect(writer, request, g.basePath+"?success")
	}
}

func (g *generic[T]) Delete(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" && request.URL.Query().Get("Id") != "" {
		id, _ := strconv.ParseInt(request.URL.Query().Get("Id"), 10, 64)
		item := g.repository.Find(id)
		g.repository.Delete(item)
		redirect(writer, request, g.basePath)
	}
}

type Pagination struct {
	Pages  int64
	Active int64
}
