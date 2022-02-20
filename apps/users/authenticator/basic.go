package authenticator

import (
	"genuine/apps/users/models"
	"genuine/core/authenticator"
	"genuine/core/repositories"
	"net/http"
)

type basic struct {
	users repositories.Repository[models.User]
}

func Basic() authenticator.Authenticator {
	return &basic{}
}

func (b *basic) Authenticate(writer http.ResponseWriter, request *http.Request) bool {
	return authenticator.Basic(writer, request, func(username, password string) bool {
		users := b.users.List("WHERE Name = '" + username + "'")
		if len(users) > 0 {
			user := &users[0]
			if username == user.Name && password == user.Password {
				return true
			}
		}
		return false
	})
}
