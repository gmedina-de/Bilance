package repository

import (
	"Bilance/model"
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
	return &model.Tag{}
}

func (r *tagRepository) NewFromQuery(row *sql.Rows) interface{} {
	var id int
	var name string
	row.Scan(&id, &name)
	return &model.Tag{id, name}
}

func (r *tagRepository) NewFromRequest(request *http.Request, id int) interface{} {
	return &model.Tag{
		id,
		request.Form.Get("Name"),
	}
}

func (r *tagRepository) List(conditions ...string) interface{} {
	var result []model.Tag
	conditions = append(conditions, "ORDER BY Id")
	r.database.Query(&result, r.NewFromQuery, conditions...)
	return result
}
