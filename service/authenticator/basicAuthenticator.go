package authenticator

import (
	"Bilance/model"
	"Bilance/service/database"
	"crypto/sha256"
	"crypto/subtle"
	"net/http"
	"strings"
)

type basicAuthenticator struct {
	database database.Database
}

func BasicAuthenticator(database database.Database) Authenticator {
	return &basicAuthenticator{database}
}

func (b *basicAuthenticator) Authenticate(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if ok {
		usernameHash := sha256.Sum256([]byte(username))
		passwordHash := sha256.Sum256([]byte(password))

		var users []model.User
		b.database.Query(&users, model.UserQuery, "WHERE Username = '"+username+"'")
		if len(users) > 0 {
			expectedUsernameHash := sha256.Sum256([]byte(users[0].Username))
			expectedPasswordHash := sha256.Sum256([]byte(users[0].Password))

			usernameMatch := subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1
			passwordMatch := subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1

			if usernameMatch && passwordMatch {
				if !strings.HasPrefix(r.URL.Path, "/admin") || users[0].Role == model.UserRoleAdmin {
					return true
				}
			}
		}
	}

	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
	return false
}
