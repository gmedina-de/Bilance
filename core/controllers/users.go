package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"homecloud/core/database"
)

type USER struct {
	web.Controller
	Database database.Database
}

//func Users(repository repositories.Repository[model.User]) Controller {
//	return &users{&generic[model.User]{repository: repository, model: model.User{}}}
//}

func Users(database database.Database) *USER {
	return &USER{Database: database}
}

func (u *USER) URLMapping() {
	u.Mapping("List", u.List)
}

func (u *USER) List() {
	//var toast string
	//if u.Ctx.Request.URL.Query().Get("success") != "" {
	//	toast = "record_saved_successfully"
	//}

	//todo use framework pagination
	u.Database.Insert("asdf")
	//limit, offset, pagination := template.HandlePagination(u.Ctx.Request, count)
	//u.Data["model"] = u.repository.Limit(limit, offset)
	//u.Data["Pagination"] = pagination
	//u.Data["Toast"] = toast
	//u.TplName = "generic.gohtml"
}
