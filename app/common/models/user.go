package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Password string
	Role     UserRole
}

type UserRole int64

const UserRoleAdmin UserRole = 0
const UserRoleNormal UserRole = 1
