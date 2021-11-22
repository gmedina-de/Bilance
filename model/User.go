package model

type User struct {
	Id       int
	Name     string
	Password string
	Role     UserRole
}

type UserRole int

const UserRoleNormal UserRole = 0
const UserRoleAdmin UserRole = 1
