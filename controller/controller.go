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
	Parameters        *Parameters
}

type Parameters struct {
	Model interface{}
	Data  interface{}
	Toast string
}

func render(writer http.ResponseWriter, request *http.Request, parameters *Parameters, title string, templates ...string) {
	// prepare templates
	for i, iTemplate := range templates {
		templates[i] = "view/" + iTemplate + ".gohtml"
	}
	templates = append(templates, "view/base.gohtml", "view/navbar.gohtml", "view/navigation.gohtml")
	tmpl := template.New("")
	f := func(currentPath string, linkPath string) string {
		if currentPath == "/" && linkPath == "/" || strings.HasPrefix(currentPath, linkPath) && linkPath != "/" {
			return " active"
		}
		return ""
	}
	tmpl.Funcs(template.FuncMap{
		"translate": localization.Translate,
		"active":    f,
	})
	tmpl, err := tmpl.ParseFiles(templates...)
	if err != nil {
		panic(err)
	}
	user := model.DeserializeUser(request.Header.Get("user"))

	// execute templates
	err = tmpl.ExecuteTemplate(writer, "base", &Context{
		user,
		handleSelectedProjectId(writer, request, user),
		request.URL.Path,
		title,
		parameters,
	})
	if err != nil {
		panic(err)
	}
}

func handleSelectedProjectId(writer http.ResponseWriter, request *http.Request, user *model.User) int64 {
	cookie, err := request.Cookie(model.SelectedProjectIdCookie)
	var selectedProjectId int64
	if err != nil {
		user := model.DeserializeUser(request.Header.Get("user"))
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
	return selectedProjectId
}

func redirect(writer http.ResponseWriter, request *http.Request, path string) {
	http.Redirect(writer, request, path, http.StatusTemporaryRedirect)
}
