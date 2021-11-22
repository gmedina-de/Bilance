package model

import (
	"Bilance/service/database"
	"database/sql"
	"net/http"
)

type tagRepository struct {
	baseRepository
}

func TagRepository(database database.Database) Repository {
	return &tagRepository{baseRepository{
		database: database,
	}}
}

func (r *tagRepository) NewEmpty() interface{} {
	return &Tag{}
}

func (r *tagRepository) NewFromQuery(row *sql.Rows) interface{} {
	var id int
	var name string
	row.Scan(&id, &name)
	return &Tag{id, name}
}

func (r *tagRepository) NewFromRequest(request *http.Request, id int) interface{} {
	return &Tag{
		id,
		request.Form.Get("Name"),
	}
}

func (r *tagRepository) Find(id string) interface{} {
	var result []Tag
	r.database.Query(&result, r.NewFromQuery, "WHERE Id = "+id)
	return &result[0]
}

func (r *tagRepository) List(conditions ...string) interface{} {
	var result []Tag
	conditions = append(conditions, "ORDER BY Id")
	r.database.Query(&result, r.NewFromQuery, conditions...)
	return result
}
