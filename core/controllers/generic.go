package controllers

import (
	"github.com/beego/beego/v2/server/web/pagination"
	"homecloud/core/models"
	"homecloud/core/repositories"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type generic[T any] struct {
	BaseController
	Repository repositories.Repository[T]
	Model      T
	Route      string
}

func Generic[T any](repository repositories.Repository[T], model T, route string) Controller {
	return &generic[T]{Repository: repository, Model: model, Route: route}
}

func (this *generic[T]) Routing(router Router) {
	router.Add(this.Route, "get:List(p)")
	router.Add(this.Route+"/:Id", "get:Edit();post:Save()")
}

const PageSize = 1
const PageSizeAll = 9223372036854775807

func (this *generic[T]) List(p string) {
	pageSize := PageSize
	if this.GetString("p") == "all" {
		pageSize = PageSizeAll
	}
	paginator := pagination.SetPaginator(this.Ctx, pageSize, this.Repository.Count())
	this.Data["Model"] = this.Repository.Limit(pageSize, paginator.Offset())
	this.Data["Paginator"] = paginator
	this.Data["Title"] = models.Plural(this.Model)
	this.TplName = "generic_list.gohtml"
}

func (this *generic[T]) Edit() {
	param := this.Ctx.Input.Param(":id")
	if param == "new" {
		this.Data["Form"] = this.Model
		this.Data["Title"] = "new"
	} else {
		id, _ := strconv.ParseInt(param, 10, 64)
		this.Data["Form"] = this.Repository.Find(id)
		this.Data["Title"] = "edit"
	}
	this.TplName = "generic_edit.gohtml"
}

func (this *generic[T]) Save() {
	param := this.Ctx.Input.Param(":id")
	record := &this.Model
	if err := this.ParseForm(record); err != nil {
		panic(err)
	}
	if param == "new" {
		this.Repository.Insert(record)
	} else {
		id, _ := strconv.ParseInt(param, 10, 64)
		reflect.ValueOf(record).Elem().FieldByName("Id").SetInt(id)
		this.Repository.Update(record)
	}
	this.Redirect(this.Route+"?success", http.StatusTemporaryRedirect)
}

func (this *generic[T]) Remove(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" && request.URL.Query().Get("Id") != "" {
		id, _ := strconv.ParseInt(request.URL.Query().Get("Id"), 10, 64)
		item := this.Repository.Find(id)
		this.Repository.Delete(item)
		http.Redirect(writer, request, strings.Replace(request.URL.Path, "/edit/delete", "", 1), http.StatusTemporaryRedirect)
	}
}
