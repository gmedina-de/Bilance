package controller

import (
	"Bilance/model"
	"Bilance/repository"
	"net/http"
	"strconv"
	"strings"
)

type crudController struct {
	repository   repository.Repository
	basePath     string
	dataProvider func(request *http.Request) interface{}
}

func (c *crudController) List(writer http.ResponseWriter, request *http.Request) {
	var toast string
	if request.URL.Query().Has("success") {
		toast = "record_saved_successfully"
	}

	projectIdString := model.GetSelectedProjectIdString(request)
	modelName := c.repository.ModelNamePlural()
	var conditions []string
	switch modelName {
	case "categories":
		conditions = append(conditions, "WHERE ProjectId = "+projectIdString)
	case "payments":
		conditions = append(conditions, "WHERE ProjectId = "+projectIdString)
		conditions = append(conditions, "ORDER BY Date DESC")
	}

	pagination, limitCondition := c.handlePagination(request, projectIdString)
	render(
		writer,
		request,
		&Parameters{Model: c.repository.List(append(conditions, limitCondition)...), Toast: toast, Data: pagination},
		modelName,
		"crud_table",
		modelName,
	)
}

type Pagination struct {
	Pages  int64
	Active int64
}

func (c *crudController) handlePagination(request *http.Request, projectIdString string) (*Pagination, string) {
	if strings.HasPrefix(request.URL.Path, "/admin/") {
		return nil, ""
	}
	var limit int64 = 10
	var page int64 = 1
	if request.URL.Query().Has("page") {
		page, _ = strconv.ParseInt(request.URL.Query().Get("page"), 10, 64)
	}
	var offset = limit * (page - 1)
	var pages = c.repository.Count("WHERE ProjectId = "+projectIdString) / limit
	pages++
	return &Pagination{pages, page},
		"LIMIT " + strconv.FormatInt(limit, 10) + " OFFSET " + strconv.FormatInt(offset, 10)
}

func (c *crudController) Edit(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		var data interface{}
		if c.dataProvider != nil {
			data = c.dataProvider(request)
		}
		modelName := c.repository.ModelNamePlural()
		if request.URL.Query().Has("Id") {
			idString := request.URL.Query().Get("Id")
			id, _ := strconv.ParseInt(idString, 10, 64)
			model := c.repository.Find(id)
			render(writer, request, &Parameters{Model: model, Data: data}, "edit", "crud_form", modelName)
		} else {
			render(writer, request, &Parameters{Model: c.repository.NewEmpty(), Data: data}, "new", "crud_form", modelName)
		}
	} else if request.Method == "POST" {
		err := request.ParseForm()
		if err != nil {
			panic(err)
		}
		if request.Form.Has("Id") {
			id, _ := strconv.ParseInt(request.Form.Get("Id"), 10, 64)
			c.repository.Update(c.repository.NewFromRequest(request, id))
		} else {
			c.repository.Insert(c.repository.NewFromRequest(request, 0))
		}
		redirect(writer, request, c.basePath+"?success")
	}
}

func (c *crudController) Delete(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" && request.URL.Query().Has("Id") {
		id, _ := strconv.ParseInt(request.URL.Query().Get("Id"), 10, 64)
		item := c.repository.Find(id)
		c.repository.Delete(item)
		redirect(writer, request, c.basePath)
	}
}
