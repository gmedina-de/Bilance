package repository

import (
	"Bilance/model"
	"Bilance/service"
	"database/sql"
	"net/http"
	"strconv"
)

type paymentRepository struct {
	baseRepository
	userRepository Repository
	tagRepository  Repository
}

func PaymentRepository(database service.Database, userRepository Repository, tagRepository Repository) Repository {
	return &paymentRepository{baseRepository{database: database}, userRepository, tagRepository}
}

func (r *paymentRepository) NewEmpty() interface{} {
	return &model.Payment{}
}

func (r *paymentRepository) NewFromQuery(row *sql.Rows) interface{} {
	var id int64
	var name string
	var amount model.EUR
	var date model.Date
	var projectId int64
	var tagId int64
	var payerId int64
	var payeeId int64
	row.Scan(&id, &name, &amount, &date, &projectId, &tagId, &payerId, &payeeId)
	return &model.Payment{
		id,
		name,
		amount,
		date,
		projectId,
		r.tagRepository.Find(tagId).(*model.Tag),
		r.userRepository.Find(payerId).(*model.User),
		r.userRepository.Find(payeeId).(*model.User),
	}
}

func (r *paymentRepository) NewFromRequest(request *http.Request, id int64) interface{} {
	//cookie, _ := request.Cookie(model.SelectedProjectIdCookie)
	//projectId, _ := strconv.ParseInt(cookie.Value, 10, 64)
	return nil
}

func (r *paymentRepository) Find(id int64) interface{} {
	var result []model.Payment
	r.database.Select(&result, r.NewFromQuery, "WHERE Id = "+strconv.FormatInt(id, 10))
	if len(result) > 0 {
		return &result[0]
	} else {
		return r.NewEmpty()
	}
}

func (r *paymentRepository) List(conditions ...string) interface{} {
	var result []model.Payment
	conditions = append(conditions, "ORDER BY Id")
	r.database.Select(&result, r.NewFromQuery, conditions...)
	return result
}
