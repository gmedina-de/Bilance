package model

func init() {
	AddModel(Person{}, "user")
	AddModel(Note{}, "edit")
	AddModel(Book{}, "book")
}

type Person struct {
	Id       int64
	Name     string `form:"required"`
	Password string `form:"required"`
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
	Note        Note
}
