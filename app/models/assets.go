package models

import (
	"genuine/app/models/register"
)

func init() {
	register.Register(Person{}, "user")
	register.Register(Note{}, "edit")
	register.Register(Book{}, "book")
}

type Person struct {
	Id       int64  `form:"-"`
	Name     string `required:"true" class:"form-control"`
	Password string `form:"Password,password" required:"true" class:"form-control"`
}

type Note struct {
	Id          int64
	Name        string
	Description string
}

type Book struct {
	Id          int64
	Name        string
	Description string
	//Note        *Note
}
