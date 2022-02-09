package authenticator

import (
	"genuine/core/models"
	"genuine/core/repositories"
	"github.com/beego/beego/v2/server/web"
	auth2 "github.com/beego/beego/v2/server/web/filter/auth"
)

type basic struct {
	repository repositories.Repository[models.User]
}

func Basic(repository repositories.Repository[models.User]) Authenticator {
	b := &basic{repository}
	web.InsertFilter("*", web.BeforeRouter, auth2.NewBasicAuthenticator(b.Authenticate, "Authorization Required"))
	return b
}

func (b *basic) Authenticate(username, password string) bool {
	u := b.repository.List("WHERE name = ?", username)
	if len(u) > 0 {
		if u[0].Password == password {
			return true
		}
	}
	return false
}
