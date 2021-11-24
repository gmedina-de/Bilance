package controller

import (
	"Bilance/localization"
	"Bilance/model"
	"Bilance/repository"
	"Bilance/service"
	"html/template"
	"net/http"
	"strconv"
)

type Controller interface {
	Routing(router service.Router)
}

type Context struct {
	Data        interface{}
	Title       string
	User        *model.User
	CurrentPath string
}

var MyLog service.Log
var MyUserRepository repository.Repository

func render(writer http.ResponseWriter, request *http.Request, title string, data interface{}, templates ...string) {
	for i, iTemplate := range templates {
		templates[i] = "view/" + iTemplate + ".gohtml"
	}
	templates = append(templates, "view/base.gohtml", "view/navbar.gohtml", "view/navigation.gohtml")
	tmpl := template.New("")
	tmpl.Funcs(template.FuncMap{
		"translate": localization.Translate,
	})
	tmpl, err := tmpl.ParseFiles(templates...)
	parseInt, _ := strconv.ParseInt(request.Header.Get("userId"), 10, 64)
	context := &Context{
		data,
		title,
		MyUserRepository.Find(parseInt).(*model.User),
		request.URL.Path,
	}
	err = tmpl.ExecuteTemplate(writer, "base", context)
	if err != nil {
		MyLog.Error(err.Error())
	}
}
