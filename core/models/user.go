package models

type User struct {
	Id       int64
	Name     string
	Password string
	Role     UserRole
}

type UserRole int64

const UserRoleNormal UserRole = 0
const UserRoleAdmin UserRole = 1
