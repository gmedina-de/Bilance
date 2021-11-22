package model

import (
	"database/sql"
	"strconv"
)

type Project struct {
	Id          int
	Name        string
	Description string
}

func (p *Project) Users() []User {
	list := MyUserRepository.List("WHERE Id IN (SELECT Id FROM ProjectUser WHERE projectId = " + strconv.Itoa(p.Id) + ")").([]User)
	return list
}

func ProjectQuery(row *sql.Rows) interface{} {
	var id int
	var name string
	var description string
	row.Scan(&id, &name, &description)
	return &Project{id, name, description}
}
