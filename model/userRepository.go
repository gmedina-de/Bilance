package model

import (
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
	return &User{}
}

func (r *userRepository) NewFromQuery(row *sql.Rows) interface{} {
	var id int
	var Name string
	var password string
	var role UserRole
	row.Scan(&id, &Name, &password, &role)
	return &User{id, Name, password, role}
}

func (r *userRepository) NewFromRequest(request *http.Request, id int) interface{} {
	admin, _ := strconv.Atoi(request.Form.Get("Role"))
	return &User{
		id,
		request.Form.Get("Name"),
		request.Form.Get("Password"),
		UserRole(admin),
	}
}

func (r *userRepository) Find(id string) interface{} {
	var result []User
	r.database.Query(&result, r.NewFromQuery, "WHERE Id = "+id)
	return &result[0]
}

func (r *userRepository) List(conditions ...string) interface{} {
	var result []User
	conditions = append(conditions, "ORDER BY Id")
	r.database.Query(&result, r.NewFromQuery, conditions...)
	return result
}
