package model

import "database/sql"

type Tag struct {
	Id   int
	Name string
}

func TagQuery(row *sql.Rows) interface{} {
	var id int
	var name string
	row.Scan(&id, &name)
	return &Tag{id, name}
}
