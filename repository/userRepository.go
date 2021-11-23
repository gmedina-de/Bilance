package repository

import (
	"Bilance/model"
	"Bilance/service/database"
	"database/sql"
	"net/http"
	"strconv"
)

type userRepository struct {
	baseRepository
}

func UserRepository(database database.Database) Repository {
	return &userRepository{baseRepository{database: database}}
}

func (r *userRepository) NewEmpty() interface{} {
	return &model.User{}
}

func (r *userRepository) NewFromQuery(row *sql.Rows) interface{} {
	var id int64
	var Name string
	var password string
	var role model.UserRole
	row.Scan(&id, &Name, &password, &role)
	return &model.User{id, Name, password, role}
}

func (r *userRepository) NewFromRequest(request *http.Request, id int64) interface{} {
	admin, _ := strconv.Atoi(request.Form.Get("Role"))
	return &model.User{
		id,
		request.Form.Get("Name"),
		request.Form.Get("Password"),
		model.UserRole(admin),
	}
}

func (r *userRepository) Find(id int64) interface{} {
	var result []model.User
	r.database.Select(&result, r.NewFromQuery, "WHERE Id = "+strconv.FormatInt(id, 10))
	return &result[0]
}

func (r *userRepository) List(conditions ...string) interface{} {
	var result []model.User
	conditions = append(conditions, "ORDER BY Id")
	r.database.Select(&result, r.NewFromQuery, conditions...)
	return result
}
