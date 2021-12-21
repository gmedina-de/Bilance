package repository

import (
	"Bilance/database"
	"Bilance/model"
	"database/sql"
	"net/http"
	"strconv"
)

type users struct {
	generic[model.User]
}

func Users(database database.Database) Repository[model.User] {
	return &users{
		generic[model.User]{
			database: database,
			model:    model.User{},
		},
	}
}

func (u users) FromQuery(row *sql.Rows) *model.User {
	user := model.User{}
	model.ScanAndPanic(row, &user.Id, &user.Name, &user.Password, &user.Role)
	return &user
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
