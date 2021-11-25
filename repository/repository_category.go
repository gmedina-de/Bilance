package repository

import (
	"Bilance/model"
	"Bilance/service"
	"database/sql"
	"net/http"
	"strconv"
)

type typeRepository struct {
	baseRepository
}

func CategoryRepository(database service.Database) Repository {
	return &typeRepository{baseRepository{database: database}}
}

func (r *typeRepository) NewEmpty() interface{} {
	return &model.Category{}
}

func (r *typeRepository) NewFromQuery(row *sql.Rows) interface{} {
	var id int64
	var name string
	var color string
	var projectId int64
	ScanAndPanic(row, &id, &name, &color, &projectId)
	return &model.Category{id, name, color, projectId}
}

func (r *typeRepository) NewFromRequest(request *http.Request, id int64) interface{} {
	cookie, _ := request.Cookie(model.SelectedProjectIdCookie)
	projectId, _ := strconv.ParseInt(cookie.Value, 10, 64)
	return &model.Category{
		id,
		request.Form.Get("Name"),
		request.Form.Get("Color"),
		projectId,
	}
}

func (r *typeRepository) Find(id int64) interface{} {
	var result []model.Category
	r.database.Select(&result, r.NewFromQuery, "WHERE Id = "+strconv.FormatInt(id, 10))
	if len(result) > 0 {
		return &result[0]
	} else {
		return r.NewEmpty()
	}
}

func (r *typeRepository) List(conditions ...string) interface{} {
	var result []model.Category
	conditions = append(conditions, "ORDER BY Id")
	r.database.Select(&result, r.NewFromQuery, conditions...)
	return result
}
