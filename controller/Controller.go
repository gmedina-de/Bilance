package controller

import (
	"Bilance/model"
	"Bilance/repository"
	"Bilance/service/router"
	"html/template"
	"net/http"
	"strconv"
)

type Controller interface {
	Routing(router router.Router)
}

type Context struct {
	Data        interface{}
	Title       string
	User        *model.User
	CurrentPath string
}

var MyUserRepository repository.Repository

func render(writer http.ResponseWriter, request *http.Request, title string, data interface{}, templates ...string) {
	for i, iTemplate := range templates {
		templates[i] = "view/" + iTemplate + ".html"
	}
	templates = append(templates, "view/base.html", "view/navigation.html", "view/header.html")
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		panic(err)
	}
	parseInt, _ := strconv.ParseInt(request.Header.Get("userId"), 10, 64)
	context := &Context{
		data,
		title,
		MyUserRepository.Find(parseInt).(*model.User),
		request.URL.Path,
	}
	err = tmpl.ExecuteTemplate(writer, "base", context)
	if err != nil {
		panic(err)
	}
}
