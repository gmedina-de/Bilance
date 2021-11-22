package authenticator

import (
	"Bilance/model"
	"crypto/sha256"
	"crypto/subtle"
	"net/http"
	"strconv"
	"strings"
)

type basicAuthenticator struct {
	userRepository model.Repository
}

func BasicAuthenticator(userRepository model.Repository) Authenticator {
	return &basicAuthenticator{userRepository}
}

func (b *basicAuthenticator) Authenticate(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if ok {
		usernameHash := sha256.Sum256([]byte(username))
		passwordHash := sha256.Sum256([]byte(password))

		var users []model.User
		users = b.userRepository.List("WHERE Name = '" + username + "'").([]model.User)
		if len(users) > 0 {
			user := users[0]
			expectedUsernameHash := sha256.Sum256([]byte(user.Name))
			expectedPasswordHash := sha256.Sum256([]byte(user.Password))

			usernameMatch := subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1
			passwordMatch := subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1

			if usernameMatch && passwordMatch {
				isAdmin := user.Role == model.UserRoleAdmin
				if !strings.HasPrefix(r.URL.Path, "/admin") || isAdmin {
					r.Header.Add("userId", strconv.Itoa(user.Id))
					return true
				}
			}
		}
	}

	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
	return false
}
