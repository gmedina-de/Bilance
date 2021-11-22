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

func render(writer http.ResponseWriter, request *http.Request, title string, templateName string, data interface{}) {
	tmpl, err := template.ParseFiles("view/"+templateName+".html", "view/base.html", "view/navigation.html")
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
