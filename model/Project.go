package model

import "database/sql"

type Project struct {
	Id          int
	Name        string
	Description string
}

func ProjectQuery(row *sql.Rows) interface{} {
	var id int
	var name string
	var description string
	row.Scan(&id, &name, &description)
	return &Project{id, name, description}
}
