package model

func init() {
	AddModel(Person{}, "user")
	AddModel(Note{}, "edit")
	AddModel(Book{}, "book")
}

type Person struct {
	Id       int64  `form:"-"`
	Name     string `form:"Password,password" required:"true" class:"form-control"`
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
