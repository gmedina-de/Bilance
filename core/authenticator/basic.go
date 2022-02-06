package authenticator

import (
	"crypto/sha256"
	"crypto/subtle"
	"homecloud/core/models"
	"homecloud/core/repositories"
	"net/http"
	"strings"
)

type basic struct {
	users repositories.Repository[models.User]
}

func Basic(users repositories.Repository[models.User]) Authenticator {
	return &basic{users}
}

func (b *basic) Authenticate(writer http.ResponseWriter, request *http.Request) bool {
	username, password, ok := request.BasicAuth()
	if ok {
		usernameHash := sha256.Sum256([]byte(username))
		passwordHash := sha256.Sum256([]byte(password))

		user, found := b.retrieveUser(username)
		if found {
			expectedUsernameHash := sha256.Sum256([]byte(user.Name))
			expectedPasswordHash := sha256.Sum256([]byte(user.Password))

			usernameMatch := subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1
			passwordMatch := subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1

			if usernameMatch && passwordMatch {
				if !strings.HasPrefix(request.URL.Path, "/admin") || user.Role == models.UserRoleAdmin {
					return true
				}
			}
		}
	}

	writer.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	http.Error(writer, "Unauthorized", http.StatusUnauthorized)
	return false
}

func (b *basic) retrieveUser(username string) (models.User, bool) {
	users := b.users.List("name = ?", username)
	if len(users) > 0 {
		return users[0], true
	} else {
		return models.User{}, false
	}
}
