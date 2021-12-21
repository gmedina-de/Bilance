package controller

import (
	"Bilance/model"
	"Bilance/repository"
	"Bilance/service"
	"net/http"
	"strconv"
	"strings"
)

type generic struct {
	repository   repository.Repository
	basePath     string
	dataProvider func(request *http.Request) interface{}
}

func (g *generic) List(writer http.ResponseWriter, request *http.Request) {
	var toast string
	if request.URL.Query().Get("success") != "" {
		toast = "record_saved_successfully"
	}

	projectIdString := model.GetSelectedProjectIdString(request)
	modelName := g.repository.ModelNamePlural()
	var conditions []string
	switch modelName {
	case "categories":
		conditions = append(conditions, "WHERE ProjectId = "+projectIdString)
	case "payments":
		conditions = append(conditions, "WHERE ProjectId = "+projectIdString)
		conditions = append(conditions, "ORDER BY Date DESC")
	}

	pagination, limitCondition := g.handlePagination(request, projectIdString)
	render(
		writer,
		request,
		&Parameters{Model: g.repository.List(append(conditions, limitCondition)...), Toast: toast, Data: pagination},
		modelName,
		"crud_table",
		modelName,
	)
}

type Pagination struct {
	Pages  int64
	Active int64
}

func (g *generic) Routing(router service.Router) {
	router.Get(g.basePath, g.List)
	router.Post(g.basePath, g.List)
	router.Get(g.basePath+"edit", g.Edit)
	router.Post(g.basePath+"edit", g.Edit)
	router.Get(g.basePath+"edit/delete", g.Delete)
}

func (g *generic) handlePagination(request *http.Request, projectIdString string) (*Pagination, string) {
	if strings.HasPrefix(request.URL.Path, "/admin/") {
		return nil, ""
	}
	var limit int64 = 10
	var page int64 = 1
	if request.URL.Query().Get("page") != "" {
		page, _ = strconv.ParseInt(request.URL.Query().Get("page"), 10, 64)
	}
	var offset = limit * (page - 1)
	var pages = g.repository.Count("WHERE ProjectId = "+projectIdString) / limit
	pages++
	return &Pagination{pages, page},
		"LIMIT " + strconv.FormatInt(limit, 10) + " OFFSET " + strconv.FormatInt(offset, 10)
}

func (g *generic) Edit(writer http.ResponseWriter, request *http.Request) {
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
			g.repository.Update(g.repository.NewFromRequest(request, id))
		} else {
			g.repository.Insert(g.repository.NewFromRequest(request, 0))
		}
		redirect(writer, request, g.basePath+"?success")
	}
}

func (g *generic) Delete(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" && request.URL.Query().Get("Id") != "" {
		id, _ := strconv.ParseInt(request.URL.Query().Get("Id"), 10, 64)
		item := g.repository.Find(id)
		g.repository.Delete(item)
		redirect(writer, request, g.basePath)
	}
}

type generic2[T model.Model[T]] struct {
	repository   repository.GRepository[T]
	basePath     string
	dataProvider func(request *http.Request) interface{}
}

func (g *generic2[T]) Routing(router service.Router) {
	router.Get(g.basePath, g.List)
	router.Post(g.basePath, g.List)
	router.Get(g.basePath+"edit", g.Edit)
	router.Post(g.basePath+"edit", g.Edit)
	router.Get(g.basePath+"edit/delete", g.Delete)
}

func (g *generic2[T]) List(writer http.ResponseWriter, request *http.Request) {

	var toast string
	if request.URL.Query().Get("success") != "" {
		toast = "record_saved_successfully"
	}

	projectIdString := model.GetSelectedProjectIdString(request)
	modelName := g.repository.ModelNamePlural()
	var conditions []string
	switch modelName {
	case "categories":
		conditions = append(conditions, "WHERE ProjectId = "+projectIdString)
	case "payments":
		conditions = append(conditions, "WHERE ProjectId = "+projectIdString)
		conditions = append(conditions, "ORDER BY Date DESC")
	}

	pagination, limitCondition := g.handlePagination(request, projectIdString)
	parameters := &Parameters{
		g.repository.List(append(conditions, limitCondition)...),
		pagination,
		toast,
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

func (g *generic2[T]) handlePagination(request *http.Request, projectIdString string) (*Pagination, string) {
	if strings.HasPrefix(request.URL.Path, "/admin/") {
		return nil, ""
	}
	var limit int64 = 10
	var page int64 = 1
	if request.URL.Query().Get("page") != "" {
		page, _ = strconv.ParseInt(request.URL.Query().Get("page"), 10, 64)
	}
	var offset = limit * (page - 1)
	var pages = g.repository.Count("WHERE ProjectId = "+projectIdString) / limit
	pages++
	return &Pagination{pages, page},
		"LIMIT " + strconv.FormatInt(limit, 10) + " OFFSET " + strconv.FormatInt(offset, 10)
}

func (g *generic2[T]) Edit(writer http.ResponseWriter, request *http.Request) {
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
			render(writer, request, &Parameters{model, data, ""}, "edit", "crud_form", modelName)
		} else {
			render(writer, request, &Parameters{g.repository.NewEmpty(), data, ""}, "new", "crud_form", modelName)
		}
	} else if request.Method == "POST" {
		err := request.ParseForm()
		if err != nil {
			panic(err)
		}
		if request.Form.Get("Id") != "" {
			id, _ := strconv.ParseInt(request.Form.Get("Id"), 10, 64)
			g.repository.Update(g.repository.NewFromRequest(request, id))
		} else {
			g.repository.Insert(g.repository.NewFromRequest(request, 0))
		}
		redirect(writer, request, g.basePath+"?success")
	}
}

func (g *generic2[T]) Delete(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" && request.URL.Query().Get("Id") != "" {
		id, _ := strconv.ParseInt(request.URL.Query().Get("Id"), 10, 64)
		item := g.repository.Find(id)
		g.repository.Delete(item)
		redirect(writer, request, g.basePath)
	}
}
