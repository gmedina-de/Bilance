package repository

import (
	"Bilance/model"
	"Bilance/service"
	"database/sql"
	"net/http"
	"strconv"
)

type categoryRepository struct {
	database service.Database
}

func CategoryRepository(database service.Database) Repository {
	return &categoryRepository{database}
}

func (r *categoryRepository) ModelName() string {
	return "category"
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
	scanAndPanic(row, &id, &name, &color, &projectId)
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
	r.database.Select(r.ModelName(), &result, "*", r.NewFromQuery, "WHERE Id = "+strconv.FormatInt(id, 10))
	if len(result) > 0 {
		return &result[0]
	} else {
		return r.NewEmpty()
	}
}

func (r *categoryRepository) List(conditions ...string) interface{} {
	var result []model.Category
	r.database.Select(r.ModelName(), &result, "*", r.NewFromQuery, conditions...)
	return result
}

func (r *categoryRepository) Count(conditions ...string) int64 {
	var result []int64
	r.database.Select(r.ModelName(), &result, "COUNT(*)", countQueryFunc, conditions...)
	return result[0]
}

func (r *categoryRepository) Insert(entity interface{}) {
	r.database.Insert(r.ModelName(), entity)
}

func (r *categoryRepository) Update(entity interface{}) {
	r.database.Update(r.ModelName(), entity)
}

func (r *categoryRepository) Delete(entity interface{}) {
	r.database.Delete(r.ModelName(), entity)
}
