package controller

import (
	. "Bilance/model"
	"Bilance/service/database"
	"Bilance/service/router"
	"github.com/joncalhoun/form"
	"html/template"
	"net/http"
)

var inputTpl = `
<label {{with .ID}}for="{{.}}"{{end}}>
	{{.Label}}
</label>
<input {{with .ID}}id="{{.}}"{{end}} type="{{.Type}}" name="{{.Name}}" placeholder="{{.Placeholder}}" {{with .Value}}value="{{.}}"{{end}}>
{{with .Footer}}
  <p>{{.}}</p>
{{end}}
`

type userController struct {
	database database.Database
}

func UserController(database database.Database) Controller {
	return &userController{database: database}
}

func (this *userController) Routing(router router.Router) {
	router.Get("/Hello", this.SayHello)
}

func (this *userController) SayHello(writer http.ResponseWriter, request *http.Request) {
	newUser := User{Username: "Admin", Password: "asdf"}
	this.database.Create(&newUser)
	var users []User
	this.database.RetrieveAll(&users)

	tpl := template.Must(template.New("").Parse(inputTpl))
	fb := form.Builder{
		InputTemplate: tpl,
	}

	tmpl, err := template.New("user.html").Funcs(fb.FuncMap()).ParseFiles("view/user.html")

	err = tmpl.Execute(writer,
		struct {
			Users   []User
			NewUser User
		}{users, newUser},
	)
	if err != nil {
		panic(err)
	}
}
