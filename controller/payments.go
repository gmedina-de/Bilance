package controller

import (
	"Bilance/localization"
	"Bilance/model"
	"Bilance/repository"
	"Bilance/server"
	"net/http"
	"strconv"
)

type payments struct {
	generic[model.Payment]
}

func Payments(
	repository repository.Repository[model.Payment],
	categories repository.Repository[model.Category],
	users repository.Repository[model.User],
) Controller {
	return &payments{
		generic[model.Payment]{
			repository: repository,
			basePath:   "/payments/",
			dataProvider: func(request *http.Request) interface{} {
				projectId := model.GetSelectedProjectIdString(request)
				return struct {
					Categories map[int64]*model.Category
					Users      map[int64]*model.User
				}{
					categories.Map("ProjectId = " + projectId),
					users.Map("Id IN (SELECT UserIds FROM Project WHERE Id = " + projectId + ")"),
				}
			},
		},
	}
}

func (c *payments) Routing(server server.Server) {
	c.generic.Routing(server)
	server.Get(c.basePath, c.List)
	server.Post(c.basePath, c.List)
}

func (c *payments) List(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("search") != "" {
		term := request.URL.Query().Get("search")
		list := c.repository.List(
			"ProjectId = "+model.GetSelectedProjectIdString(request),
			"AND (Name LIKE '%"+term+"%'",
			"OR CategoryId IN (SELECT Id FROM Category WHERE Name LIKE '%"+term+"%')",
			"OR Date LIKE '"+term+"%')",
			"ORDER BY Date",
		)
		render(
			writer,
			request,
			&Parameters{
				Model: list,
				Toast: strconv.Itoa(len(list)) + " " + localization.Translate("records_found"),
			},
			"search_results",
			"crud_table",
			c.repository.ModelNamePlural(),
		)
	} else {
		c.generic.List(writer, request)
	}
}
