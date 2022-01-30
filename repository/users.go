package repository

import (
	"Bilance/database"
	"Bilance/model"
	"net/http"
	"strconv"
)

type users struct {
	*generic[model.User]
}

func Users(database database.Database) Repository[model.User] {
	u := &users{Generic(database, model.User{})}
	count := u.Count("name = admin")
	if count < 1 {
		u.Insert(&model.User{
			Id:       0,
			Name:     "admin",
			Password: "admin",
			Role:     1,
		})
	}
	return u
}

func (u users) FromRequest(request *http.Request, id int64) *model.User {
	admin, _ := strconv.Atoi(request.Form.Get("Role"))
	return &model.User{
		Id:       id,
		Name:     request.Form.Get("Name"),
		Password: request.Form.Get("Password"),
		Role:     model.UserRole(admin),
	}
}
