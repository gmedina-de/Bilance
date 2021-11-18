package controller

import (
	. "Bilance/model"
	"Bilance/service"
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
	repository service.Repository
}

func UserController(repository service.Repository) Controller {
	return &userController{repository: repository}
}

func (this *userController) Routing(router service.Router) {
	router.Get("/Hello", this.SayHello)
}

func (this *userController) SayHello(writer http.ResponseWriter, request *http.Request) {
	newUser := User{Username: "Admin", Password: "asdf"}
	this.repository.Create(&newUser)
	var users []User
	this.repository.RetrieveAll(&users)

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
