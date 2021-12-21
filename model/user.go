package model

import (
	"database/sql"
	"net/http"
	"strconv"
)

type User struct {
	Id       int64
	Name     string
	Password string
	Role     UserRole
}

func (u User) FromQuery(row *sql.Rows) *User {
	ScanAndPanic(row, &u.Id, &u.Name, &u.Password, &u.Role)
	return &u
}

func (u User) FromRequest(request *http.Request, id int64) *User {
	admin, _ := strconv.Atoi(request.Form.Get("Role"))
	u.Id = id
	u.Name = request.Form.Get("Name")
	u.Password = request.Form.Get("Password")
	u.Role = UserRole(admin)
	return &u
}

type UserRole int64

const UserRoleNormal UserRole = 0
const UserRoleAdmin UserRole = 1
