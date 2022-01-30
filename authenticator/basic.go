package authenticator

import (
	"Bilance/model"
	"Bilance/repository"
	"crypto/sha256"
	"crypto/subtle"
	"net/http"
	"strconv"
	"strings"
)

type basic struct {
	users    repository.Repository[model.User]
	projects repository.Repository[model.Project]
}

func Basic(users repository.Repository[model.User], projects repository.Repository[model.Project]) Authenticator {
	return &basic{users, projects}
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
				if !strings.HasPrefix(request.URL.Path, "/admin") || user.Role == model.UserRoleAdmin {
					projects := b.projects.List("UserIds LIKE '%" + strconv.FormatInt(user.Id, 10) + "%'")
					if b.isProjectAccessible(user, projects, writer, request) {
						request.Header.Add("user", model.Serialize(user))
						request.Header.Add("user", model.Serialize(projects))
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

func (b *basic) retrieveUser(username string) (model.User, bool) {
	users := b.users.List("name = ?", username)
	if len(users) > 0 {
		return users[0], true
	} else {
		return model.User{}, false
	}
}

func (b *basic) isProjectAccessible(user model.User, projects []model.Project, writer http.ResponseWriter, request *http.Request) bool {
	projectId := model.GetSelectedProjectId(request)
	for _, project := range projects {
		if project.Id == projectId {
			return true
		}
	}
	if len(projects) > 0 {
		selectedProjectId := strconv.FormatInt(projects[0].Id, 10)
		model.SetSelectedProjectId(writer, selectedProjectId)
		return true
	} else if user.Role == model.UserRoleAdmin {
		model.SetSelectedProjectId(writer, "0")
		return true
	}
	return false
}
