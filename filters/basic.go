package filters

import (
	"genuine/models"
	"genuine/repositories"
	"net/http"
)

type basic struct {
	users repositories.Repository[models.User]
}

func Basic(users repositories.Repository[models.User]) Filter {
	return &basic{users}
}

func (b *basic) Filter(writer http.ResponseWriter, request *http.Request) bool {
	username, password, ok := request.BasicAuth()
	if ok {
		users := b.users.List("Name = ?", username)
		if len(users) > 0 {
			user := &users[0]
			if username == user.Name && password == user.Password {
				// todo be more specific which urls are not allowed for non admins
				if request.URL.Path != "/settings/users" {
					return true
				} else if user.IsAdmin {
					return true
				}
			}
		}
	}
	writer.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	http.Error(writer, "Unauthorized", http.StatusUnauthorized)
	return false
}
