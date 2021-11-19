package controller

import (
	"Bilance/service/router"
	"html/template"
	"net/http"
)

type Controller interface {
	Routing(router router.Router)
}

func render(writer http.ResponseWriter, templateName string, data interface{}) {
	tmpl, err := template.ParseFiles("view/"+templateName+".html", "view/base.html")
	err = tmpl.ExecuteTemplate(writer, "base", data)
	if err != nil {
		panic(err)
	}
}
