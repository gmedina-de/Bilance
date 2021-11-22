package model

import "database/sql"

type User struct {
	Id       int
	Name     string
	Password string
	Role     UserRole
}

type UserRole int

const UserRoleNormal UserRole = 0
const UserRoleAdmin UserRole = 1

func UserQuery(row *sql.Rows) interface{} {
	var id int
	var Name string
	var password string
	var role UserRole
	row.Scan(&id, &Name, &password, &role)
	return &User{id, Name, password, role}
}
