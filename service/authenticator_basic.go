package service

import (
	"Bilance/model"
	"crypto/sha256"
	"crypto/subtle"
	"database/sql"
	"net/http"
	"strings"
)

type basicAuthenticator struct {
	database Database
}

func BasicAuthenticator(database Database) Authenticator {
	return &basicAuthenticator{database}
}

func (b *basicAuthenticator) Authenticate(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if ok {
		usernameHash := sha256.Sum256([]byte(username))
		passwordHash := sha256.Sum256([]byte(password))

		var users []model.User
		b.database.Select(&users, func(row *sql.Rows) interface{} {
			var id int64
			var Name string
			var password string
			var role model.UserRole
			row.Scan(&id, &Name, &password, &role)
			return &model.User{id, Name, password, role}
		}, "WHERE Name = '"+username+"'")
		if len(users) > 0 {
			user := users[0]
			expectedUsernameHash := sha256.Sum256([]byte(user.Name))
			expectedPasswordHash := sha256.Sum256([]byte(user.Password))

			usernameMatch := subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1
			passwordMatch := subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1

			if usernameMatch && passwordMatch {
				isAdmin := user.Role == model.UserRoleAdmin
				if !strings.HasPrefix(r.URL.Path, "/admin") || isAdmin {
					r.Header.Add("user", user.Serialize())
					return true
				}
			}
		}
	}

	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
	return false
}
