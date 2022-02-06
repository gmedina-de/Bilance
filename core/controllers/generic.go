package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"homecloud/core/repositories"
	"homecloud/core/template"
	"net/http"
	"strconv"
	"strings"
)

type Generic[T any] struct {
	BaseController
	Repository repositories.Repository[T]
	Model      T
	Route      string
}

func Generics[T any](repository repositories.Repository[T], model T, route string) *Generic[T] {
	return &Generic[T]{Repository: repository, Model: model, Route: route}
}

const GenericMethods = "get:List"

func (this *Generic[T]) Routing() {
	web.Router(this.Route, this, GenericMethods)
}

func (this *Generic[T]) List() {

	var toast string
	if this.Ctx.Request.URL.Query().Get("success") != "" {
		toast = "record_saved_successfully"
	}

	//todo use framework pagination
	count := this.Repository.Count()
	limit, offset, pagination := template.HandlePagination(this.Ctx.Request, count)
	i := this.Repository.Limit(limit, offset)
	this.Data["Model"] = i
	this.Data["Pagination"] = pagination
	this.Data["Toast"] = toast
	this.TplName = "generic.gohtml"
}

func (this *Generic[T]) Edit() {
	//if request.Method == "GET" {
	//	//var data interface{}
	//	//if this.DataProvider != nil {
	//	//	data = this.DataProvider(request)
	//	//}
	//	if request.URL.Query().Get("Id") != "" {
	//		idString := request.URL.Query().Get("Id")
	//		id, _ := strconv.ParseInt(idString, 10, 64)
	//		_ = this.Repository.Find(id)
	//		//template.Render(writer, request, "edit", &template.Parameters{Model: Model, Data: data}, "core/template/form.gohtml")
	//	} else {
	//		//template.Render(writer, request, "new", &template.Parameters{Model: this.Model, Data: data}, "core/template/form.gohtml")
	//	}
	//} else if request.Method == "POST" {
	//	err := request.ParseForm()
	//	if err != nil {
	//		panic(err)
	//	}
	//	if request.Form.Get("Id") != "" {
	//		//this.Repository.Update(this.FromRequest(request))
	//	} else {
	//		//this.Repository.Insert(this.FromRequest(request))
	//	}
	//	http.Redirect(writer, request, strings.Replace(request.URL.Path, "/edit", "", 1)+"?success", http.StatusTemporaryRedirect)
	//}
}

func (this *Generic[T]) Remove(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" && request.URL.Query().Get("Id") != "" {
		id, _ := strconv.ParseInt(request.URL.Query().Get("Id"), 10, 64)
		item := this.Repository.Find(id)
		this.Repository.Delete(item)
		http.Redirect(writer, request, strings.Replace(request.URL.Path, "/edit/delete", "", 1), http.StatusTemporaryRedirect)
	}
}
