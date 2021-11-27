package repository

import (
	"Bilance/model"
	"Bilance/service"
	"database/sql"
	"net/http"
	"strconv"
)

type categoryRepository struct {
	baseRepository
}

func CategoryRepository(database service.Database) Repository {
	return &categoryRepository{baseRepository{database: database}}
}

func (r *categoryRepository) ModelNamePlural() string {
	return "categories"
}

func (r *categoryRepository) NewEmpty() interface{} {
	return &model.Category{}
}

func (r *categoryRepository) NewFromQuery(row *sql.Rows) interface{} {
	var id int64
	var name string
	var color string
	var projectId int64
	ScanAndPanic(row, &id, &name, &color, &projectId)
	return &model.Category{id, name, color, projectId}
}

func (r *categoryRepository) NewFromRequest(request *http.Request, id int64) interface{} {
	projectId := model.GetSelectedProjectId(request)
	return &model.Category{
		id,
		request.Form.Get("Name"),
		request.Form.Get("Color"),
		projectId,
	}
}

func (r *categoryRepository) Find(id int64) interface{} {
	var result []model.Category
	r.database.Select(&result, r.NewFromQuery, "WHERE Id = "+strconv.FormatInt(id, 10))
	if len(result) > 0 {
		return &result[0]
	} else {
		return r.NewEmpty()
	}
}

func (r *categoryRepository) List(conditions ...string) interface{} {
	var result []model.Category
	conditions = append(conditions, "ORDER BY Name")
	r.database.Select(&result, r.NewFromQuery, conditions...)
	return result
}
