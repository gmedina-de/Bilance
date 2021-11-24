package service

import (
	"Bilance/model"
	"crypto/sha256"
	"crypto/subtle"
	"database/sql"
	"net/http"
	"strconv"
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

		user, found := b.retrieveUser(username)
		if found {
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

func (b *basicAuthenticator) retrieveUser(username string) (model.User, bool) {
	var users []model.User
	b.database.Select(
		&users,
		func(row *sql.Rows) interface{} {
			var id int64
			var Name string
			var password string
			var role model.UserRole
			row.Scan(&id, &Name, &password, &role)
			var projects []model.Project
			b.database.Select(
				&projects,
				func(row *sql.Rows) interface{} {
					var id int64
					var name string
					var description string
					row.Scan(&id, &name, &description)
					project := model.Project{
						id,
						name,
						description,
						nil,
						nil,
						nil,
						nil,
					}
					return &project
				},
				"WHERE Id IN (SELECT ProjectId FROM ProjectUser WHERE UserId = "+strconv.FormatInt(id, 10)+")",
			)
			return &model.User{id, Name, password, role, projects}
		},
		"WHERE Name = '"+username+"'",
	)
	if len(users) > 0 {
		return users[0], true
	} else {
		return model.User{}, false
	}
}
