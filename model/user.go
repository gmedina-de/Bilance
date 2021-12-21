package model

import (
	"bytes"
	"database/sql"
	"encoding/base64"
	"encoding/gob"
	"net/http"
	"strconv"
)

type User struct {
	Id       int64
	Name     string
	Password string
	Role     UserRole
	Projects []Project
}

func (u User) FromQuery(row *sql.Rows) *User {
	scanAndPanic(row, &u.Id, &u.Name, &u.Password, &u.Role)
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

func (user *User) Serialize() string {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(user)
	return base64.StdEncoding.EncodeToString(b.Bytes())
}

func DeserializeUser(str string) *User {
	user := User{}
	by, _ := base64.StdEncoding.DecodeString(str)
	b := bytes.Buffer{}
	b.Write(by)
	d := gob.NewDecoder(&b)
	_ = d.Decode(&user)
	return &user
}
