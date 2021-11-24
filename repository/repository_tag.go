package repository

import (
	"Bilance/model"
	"Bilance/service"
	"database/sql"
	"net/http"
	"strconv"
)

type tagRepository struct {
	baseRepository
}

func TagRepository(database service.Database) Repository {
	return &tagRepository{baseRepository{database: database}}
}

func (r *tagRepository) NewEmpty() interface{} {
	return &model.Tag{}
}

func (r *tagRepository) NewFromQuery(row *sql.Rows) interface{} {
	var id int64
	var name string
	var projectId int64
	row.Scan(&id, &name, &projectId)
	return &model.Tag{id, name, projectId}
}

func (r *tagRepository) NewFromRequest(request *http.Request, id int64) interface{} {
	cookie, _ := request.Cookie("SelectedProjectId")
	projectId, _ := strconv.ParseInt(cookie.Value, 10, 64)
	return &model.Tag{
		id,
		request.Form.Get("Name"),
		projectId,
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
