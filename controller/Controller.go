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
	Admin       bool
	CurrentPath string
}

func render(writer http.ResponseWriter, request *http.Request, templateName string, data interface{}) {
	tmpl, err := template.ParseFiles("view/"+templateName+".html", "view/base.html", "view/navigation.html")
	if err != nil {
		panic(err)
	}
	path := request.URL.Path
	context := &Context{data, request.Header.Get("isAdmin") == "true", path}
	err = tmpl.ExecuteTemplate(writer, "base", context)
	if err != nil {
		panic(err)
	}
}
