package controller

import (
	"Bilance/localization"
	"Bilance/model"
	"Bilance/service"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Controller interface {
	Routing(router service.Router)
}

type Context struct {
	User              *model.User
	SelectedProjectId int64
	Path              string
	Title             string
	Toast             string
	Data              interface{}
}

func render(writer http.ResponseWriter, request *http.Request, title string, data interface{}, templates ...string) {
	// prepare templates
	for i, iTemplate := range templates {
		templates[i] = "view/" + iTemplate + ".gohtml"
	}
	templates = append(templates, "view/base.gohtml", "view/navbar.gohtml", "view/navigation.gohtml")
	tmpl := template.New("")
	tmpl.Funcs(template.FuncMap{
		"translate": localization.Translate,
		"active": func(currentPath string, linkPath string) string {
			if currentPath == "/" && linkPath == "/" || strings.HasPrefix(currentPath, linkPath) && linkPath != "/" {
				return " active"
			}
			return ""
		},
	})
	tmpl, err := tmpl.ParseFiles(templates...)
	if err != nil {
		panic(err)
	}

	// deserialize user
	user := model.DeserializeUser(request.Header.Get("user"))
	cookie, err := request.Cookie(model.SelectedProjectIdCookie)
	var selectedProjectId int64
	if err != nil {
		selectedProjectId = user.Projects[0].Id
		expiration := time.Now().Add(365 * 24 * time.Hour)
		http.SetCookie(writer, &http.Cookie{
			Name:    model.SelectedProjectIdCookie,
			Value:   strconv.FormatInt(selectedProjectId, 10),
			Path:    "/",
			Expires: expiration,
		},
		)
	} else {
		selectedProjectId, _ = strconv.ParseInt(cookie.Value, 10, 64)
	}

	// execute templates
	context := &Context{
		user,
		selectedProjectId,
		request.URL.Path,
		title,
		"record_saved_successfully",
		data,
	}
	err = tmpl.ExecuteTemplate(writer, "base", context)
	if err != nil {
		panic(err)
	}
}

func redirect(writer http.ResponseWriter, request *http.Request, path string) {
	http.Redirect(writer, request, path, http.StatusTemporaryRedirect)
}
