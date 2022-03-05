package models

import (
	"genuine/core/models"
)

type User struct {
	models.Model
	Name     string
	Password string
	IsAdmin  bool
}

func (u User) String() string {
	return u.Name
}
