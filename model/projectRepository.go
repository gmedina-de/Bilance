package model

import (
	"Bilance/service/database"
	"database/sql"
	"net/http"
)

type projectRepository struct {
	baseRepository
}

func ProjectRepository(database database.Database) Repository {
	return &projectRepository{baseRepository{
		database: database,
	}}
}

func (r *projectRepository) NewEmpty() interface{} {
	return &Project{}
}

func (r *projectRepository) NewFromQuery(row *sql.Rows) interface{} {
	var id int
	var name string
	var description string
	row.Scan(&id, &name, &description)
	return &Project{id, name, description}
}

func (r *projectRepository) NewFromRequest(request *http.Request, id int) interface{} {
	return &Project{
		id,
		request.Form.Get("Name"),
		request.Form.Get("Description"),
	}
}

func (r *projectRepository) Find(id string) interface{} {
	var result []Project
	r.database.Query(&result, r.NewFromQuery, "WHERE Id = "+id)
	return &result[0]
}

func (r *projectRepository) List(conditions ...string) interface{} {
	var result []Project
	conditions = append(conditions, "ORDER BY Id")
	r.database.Query(&result, r.NewFromQuery, conditions...)
	return result
}
