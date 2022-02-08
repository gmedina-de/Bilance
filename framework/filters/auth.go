package filters

import (
	"genuine/framework/models"
	"genuine/framework/repositories"
	"github.com/beego/beego/v2/server/web"
	auth2 "github.com/beego/beego/v2/server/web/filter/auth"
)

type auth struct {
	repository repositories.Repository[models.User]
}

func Auth(repository repositories.Repository[models.User]) Filter {
	return &auth{repository}
}

func (a *auth) Pattern() string {
	return "*"
}

func (a *auth) Func() web.FilterFunc {
	return auth2.NewBasicAuthenticator(func(username, password string) bool {
		u := a.repository.List("WHERE name = ?", username)
		if len(u) > 0 {
			if u[0].Password == password {
				return true
			}
		}
		return false
	}, "Authorization Required")
}

func (a *auth) Pos() int {
	return web.BeforeRouter
}

func (a *auth) Opts() []web.FilterOpt {
	return nil
}
