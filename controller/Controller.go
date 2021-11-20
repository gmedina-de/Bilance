package controller

import (
	"Bilance/service/router"
	"html/template"
	"net/http"
	"strings"
)

type Controller interface {
	Routing(router router.Router)
}

type Context struct {
	Data  interface{}
	Admin bool
}

func render(writer http.ResponseWriter, request *http.Request, templateName string, data interface{}) {
	tmpl, err := template.ParseFiles("view/"+templateName+".html", "view/base.html", "view/navigation.html")

	err = tmpl.ExecuteTemplate(writer, "base", &Context{data, strings.HasPrefix(request.URL.Path, "/admin")})
	if err != nil {
		panic(err)
	}
}
