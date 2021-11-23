package repository

import (
	"Bilance/model"
	"Bilance/service/database"
	"database/sql"
	"net/http"
	"strconv"
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
	var id int64
	var name string
	row.Scan(&id, &name)
	return &model.Tag{id, name}
}

func (r *tagRepository) NewFromRequest(request *http.Request, id int64) interface{} {
	return &model.Tag{
		id,
		request.Form.Get("Name"),
	}
}

func (r *tagRepository) Find(id int64) interface{} {
	var result []model.Tag
	r.database.Select(&result, r.NewFromQuery, "WHERE Id = "+strconv.FormatInt(id, 10))
	return &result[0]
}

func (r *tagRepository) List(conditions ...string) interface{} {
	var result []model.Tag
	conditions = append(conditions, "ORDER BY Id")
	r.database.Select(&result, r.NewFromQuery, conditions...)
	return result
}
