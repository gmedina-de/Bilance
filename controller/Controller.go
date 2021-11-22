package controller

import (
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
	Admin       bool
	CurrentPath string
}

func render(writer http.ResponseWriter, request *http.Request, title string, data interface{}, templates ...string) {
	for i, iTemplate := range templates {
		templates[i] = "view/" + iTemplate + ".html"
	}
	templates = append(templates, "view/base.html", "view/navigation.html")
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		panic(err)
	}
	context := &Context{
		data,
		title,
		request.Header.Get("isAdmin") == "true",
		request.URL.Path,
	}
	err = tmpl.ExecuteTemplate(writer, "base", context)
	if err != nil {
		panic(err)
	}
}
