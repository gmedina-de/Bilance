package model

type Project struct {
	Id          int64
	Name        string
	Description string
	Payments    []Payment
	Categories  []Category
	Users       []User
}

type ProjectUser struct {
	Id        int64
	ProjectId int64
	UserId    int64
}
