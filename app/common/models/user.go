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

type UserRole int64

const UserRoleAdmin UserRole = 0
const UserRoleNormal UserRole = 1
