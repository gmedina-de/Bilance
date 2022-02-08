package models

type User struct {
	Id       int64
	Name     string
	Password string
	Role     UserRole
}

type UserRole int64

const UserRoleAdmin UserRole = 0
const UserRoleNormal UserRole = 1
