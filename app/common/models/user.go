package models

import (
	"genuine/core/models"
)

type User struct {
	models.Model
	Name     string
	Password string
	Role     UserRole
}

func (u User) String() string {
	return u.Name
}

type UserRole int64

const UserRoleAdmin UserRole = 0
const UserRoleNormal UserRole = 1
