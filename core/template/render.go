package template

import (
	"github.com/beego/beego/v2/server/web"
	"homecloud/core/localization"
	"homecloud/core/model"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type Context struct {
	User               *model.User
	Path               string
	Title              string
	Navigation         []*MenuItem
	Navigation2        []*MenuItem
	CurrentNavigation1 *MenuItem
	CurrentNavigation2 *MenuItem
	Parameters         *Parameters
}

type Parameters struct {
	Model      interface{}
	Data       interface{}
	Pagination *Pagination
	Toast      string
}

func init() {
	web.AddFuncMap("translate", localization.Translate)
	web.AddFuncMap("paginate", paginate)
	web.AddFuncMap("inputs", inputs)
	web.AddFuncMap("td", td)
	web.AddFuncMap("th", th)
	web.AddFuncMap("sum", func(a int, b int) int { return a + b })
	web.AddFuncMap("contains", func(a string, b int64) bool { return strings.Contains(a, strconv.FormatInt(b, 10)) })
}

var funcMap = template.FuncMap{
	"translate": localization.Translate,
	"paginate":  paginate,
	"inputs":    inputs,
	"td":        td,
	"th":        th,
	"sum": func(a int, b int) int {
		return a + b
	},
	"contains": func(a string, b int64) bool {
		return strings.Contains(a, strconv.FormatInt(b, 10))
	},
}

func Render(writer http.ResponseWriter, request *http.Request, title string, parameters *Parameters, templates ...string) {
	// prepare templates (optimized?)
	templates = append(templates,
		"core/template/base.gohtml",
		"core/template/navigation1.gohtml",
		"core/template/navigation2.gohtml",
		"core/template/pagination.gohtml",
	)
	tmpl := template.New("")

	tmpl.Funcs(funcMap)
	tmpl, err := tmpl.ParseFiles(templates...)
	if err != nil {
		panic(err)
	}
	var user model.User
	model.Deserialize(request.Header.Get("user"), &user)
	currentNavigation := GetCurrentNavigation(request.URL.Path, Navigation)
	err = tmpl.ExecuteTemplate(writer, "base", &Context{
		&user,
		request.URL.Path,
		title,
		Navigation,
		currentNavigation.SubMenu,
		currentNavigation,
		GetCurrentNavigation(request.URL.Path, currentNavigation.SubMenu),
		parameters,
	})
	if err != nil {
		panic(err)
	}
}

func Render2(c web.Controller, t string) {

}
