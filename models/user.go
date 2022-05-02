package models

type User struct {
	Model
	Name     string
	Password string
	IsAdmin  bool
}

func (u User) String() string {
	return u.Name
}
