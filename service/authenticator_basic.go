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

func (b *basicAuthenticator) Authenticate(writer http.ResponseWriter, request *http.Request) bool {
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
				if !strings.HasPrefix(request.URL.Path, "/admin") || user.Role == model.UserRoleAdmin {
					if b.isProjectAccessible(user, writer, request) {
						request.Header.Add("user", user.Serialize())
						return true
					}
				}
			}
		}
	}

	writer.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	http.Error(writer, "Unauthorized", http.StatusUnauthorized)
	return false
}

func (b *basicAuthenticator) retrieveUser(username string) (model.User, bool) {
	var users []model.User
	b.database.Select(
		"User",
		&users,
		"*",
		func(row *sql.Rows) interface{} {
			var id int64
			var Name string
			var password string
			var role model.UserRole
			row.Scan(&id, &Name, &password, &role)
			var projects []model.Project
			b.database.Select(
				"Project",
				&projects,
				"*",
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

func (b *basicAuthenticator) isProjectAccessible(user model.User, writer http.ResponseWriter, request *http.Request) bool {
	projectId := model.GetSelectedProjectId(request)
	for _, project := range user.Projects {
		if project.Id == projectId {
			return true
		}
	}
	if len(user.Projects) > 0 {
		selectedProjectId := strconv.FormatInt(user.Projects[0].Id, 10)
		model.SetSelectedProjectId(writer, selectedProjectId)
		return true
	} else if user.Role == model.UserRoleAdmin {
		model.SetSelectedProjectId(writer, "0")
		return true
	}
	return false
}
