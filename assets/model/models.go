package model

func init() {
	AddModel(Person{}, "user")
	AddModel(Note{}, "edit")
}

type Person struct {
	Id   int64
	Name string
	Age  int
}

type Note struct {
	Id          int64
	Name        string
	Description string
}
