package authenticator

import (
	"crypto/sha256"
	"crypto/subtle"
	"genuine/apps/users/models"
	"genuine/core/authenticator"
	"genuine/core/repositories"
	"net/http"
)

type basic struct {
	Users repositories.Repository[models.User]
}

func Basic() authenticator.Authenticator {
	return &basic{}
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
				return true
			}
		}
	}
	writer.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	http.Error(writer, "Unauthorized", http.StatusUnauthorized)
	return false
}

func (b *basic) retrieveUser(username string) (*models.User, bool) {
	users := b.Users.List("WHERE Name = '" + username + "'")
	if len(users) > 0 {
		return &users[0], true
	} else {
		return nil, false
	}
}
