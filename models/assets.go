package models

type Person struct {
	Model
	Name     string
	Password string
}

func (p Person) Icon() string {
	return "user"
}

func (p Person) String() string {
	return p.Name
}

type Note struct {
	Model
	Name        string
	Description string
}

func (n Note) String() string {
	return n.Name
}

func (n Note) Icon() string {
	return "edit"
}

type Book struct {
	Model
	Name        string
	Description string
	Note        Note
	NoteID      uint
}

func (b Book) Icon() string {
	return "book"
}

func (b Book) String() string {
	return b.Name
}

type Asset interface {
	Icon() string
}
