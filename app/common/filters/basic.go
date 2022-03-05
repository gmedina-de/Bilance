package filters

import (
	"genuine/app/common/models"
	"genuine/core/filter"
	"genuine/core/repositories"
	"net/http"
)

type basic struct {
	users repositories.Repository[models.User]
}

func Basic(users repositories.Repository[models.User]) filter.Filter {
	return &basic{users}
}

func (b *basic) Filter(writer http.ResponseWriter, request *http.Request) bool {
	return filter.Basic(writer, request, func(username, password string) bool {
		users := b.users.List("Name = '" + username + "'")
		if len(users) > 0 {
			user := &users[0]
			if username == user.Name && password == user.Password {
				return true
			}
		}
		return false
	})
}
