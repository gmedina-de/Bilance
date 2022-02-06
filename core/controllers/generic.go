package controllers

import (
	"homecloud/core/repositories"
	"homecloud/core/template"
	"net/http"
	"strconv"
	"strings"
)

type generic[T any] struct {
	BaseController
	repository repositories.Repository[T]
	model      T
}

func Generic[T any](repository repositories.Repository[T], model T) Controller {
	return &generic[T]{BaseController: BaseController{}, repository: repository, model: model}
}

func (this *generic[T]) List() {
	var toast string
	if this.Ctx.Request.URL.Query().Get("success") != "" {
		toast = "record_saved_successfully"
	}

	//todo use framework pagination
	count := this.repository.Count()
	limit, offset, pagination := template.HandlePagination(this.Ctx.Request, count)
	this.Data["model"] = this.repository.Limit(limit, offset)
	this.Data["Pagination"] = pagination
	this.Data["Toast"] = toast
	this.TplName = "generic.gohtml"
}

func (this *generic[T]) Edit() {
	//if request.Method == "GET" {
	//	//var data interface{}
	//	//if this.DataProvider != nil {
	//	//	data = this.DataProvider(request)
	//	//}
	//	if request.URL.Query().Get("Id") != "" {
	//		idString := request.URL.Query().Get("Id")
	//		id, _ := strconv.ParseInt(idString, 10, 64)
	//		_ = this.repository.Find(id)
	//		//template.Render(writer, request, "edit", &template.Parameters{model: model, Data: data}, "core/template/form.gohtml")
	//	} else {
	//		//template.Render(writer, request, "new", &template.Parameters{model: this.model, Data: data}, "core/template/form.gohtml")
	//	}
	//} else if request.Method == "POST" {
	//	err := request.ParseForm()
	//	if err != nil {
	//		panic(err)
	//	}
	//	if request.Form.Get("Id") != "" {
	//		//this.repository.Update(this.FromRequest(request))
	//	} else {
	//		//this.repository.Insert(this.FromRequest(request))
	//	}
	//	http.Redirect(writer, request, strings.Replace(request.URL.Path, "/edit", "", 1)+"?success", http.StatusTemporaryRedirect)
	//}
}

func (this *generic[T]) Remove(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" && request.URL.Query().Get("Id") != "" {
		id, _ := strconv.ParseInt(request.URL.Query().Get("Id"), 10, 64)
		item := this.repository.Find(id)
		this.repository.Delete(item)
		http.Redirect(writer, request, strings.Replace(request.URL.Path, "/edit/delete", "", 1), http.StatusTemporaryRedirect)
	}
}
