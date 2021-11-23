package controller

import (
	"Bilance/model"
	"Bilance/service/router"
	"html/template"
	"net/http"
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

func render(writer http.ResponseWriter, request *http.Request, title string, data interface{}, templates ...string) {
	for i, iTemplate := range templates {
		templates[i] = "view/" + iTemplate + ".html"
	}
	templates = append(templates, "view/base.html", "view/navigation.html", "view/header.html")
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		panic(err)
	}
	context := &Context{
		data,
		title,
		model.MyUserRepository.Find(request.Header.Get("userId")).(*model.User),
		request.URL.Path,
	}
	err = tmpl.ExecuteTemplate(writer, "base", context)
	if err != nil {
		panic(err)
	}
}
