package controller

import (
	"homecloud/core/model"
	"homecloud/core/repository"
	"homecloud/core/server"
)

type users struct {
	Generic[model.User]
}

func Users(repository repository.Repository[model.User]) Controller {
	return &users{
		Generic[model.User]{
			BaseTemplate: "core/template/users.gohtml",
			Repository:   repository,
		},
	}
}

func (c *users) Routing(server server.Server) {
	c.Generic.Routing(server)
	AddNavigation1("users", "user", "/users")
}
