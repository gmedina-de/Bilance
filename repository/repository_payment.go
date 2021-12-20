package repository

import (
	"Bilance/model"
	"Bilance/service"
	"database/sql"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type paymentRepository struct {
	database           service.Database
	userRepository     Repository
	categoryRepository GenericRepository[model.Category]
}

func PaymentRepository(database service.Database, userRepository Repository, categoryRepository GenericRepository[model.Category]) Repository {
	return &paymentRepository{database, userRepository, categoryRepository}
}

func (r *paymentRepository) ModelName() string {
	return "payment"
}

func (r *paymentRepository) ModelNamePlural() string {
	return "payments"
}

func (r *paymentRepository) NewEmpty() interface{} {
	return &model.Payment{
		Category: &model.Category{},
		Payer:    &model.User{},
		Payee:    &model.User{},
		Date:     model.Date(time.Now().Format(model.DateLayoutISO)),
	}
}

func (r *paymentRepository) NewFromQuery(row *sql.Rows) interface{} {
	var id int64
	var name string
	var amount int64
	var date string
	var projectId int64
	var categoryId int64
	var payerId int64
	var payeeId int64
	scanAndPanic(row, &id, &name, &amount, &date, &projectId, &categoryId, &payerId, &payeeId)
	category := r.categoryRepository.Find(categoryId)
	payer := r.userRepository.Find(payerId).(*model.User)
	payee := r.userRepository.Find(payeeId).(*model.User)
	return &model.Payment{
		id,
		name,
		model.EUR(amount),
		model.Date(date),
		projectId,
		category,
		payer,
		payee,
	}
}

func (r *paymentRepository) NewFromRequest(request *http.Request, id int64) interface{} {
	name := request.Form.Get("Name")
	date := request.Form.Get("Date")
	projectId := model.GetSelectedProjectId(request)
	amountInput := request.Form.Get("Amount")
	amountString := strings.Replace(amountInput, ".", "", 1)
	amount, _ := strconv.ParseInt(amountString, 10, 64)
	if !strings.Contains(amountInput, ".") {
		amount *= 100
	}
	categoryId, _ := strconv.ParseInt(request.Form.Get("CategoryId"), 10, 64)
	payerId, _ := strconv.ParseInt(request.Form.Get("PayerId"), 10, 64)
	payeeId, _ := strconv.ParseInt(request.Form.Get("PayeeId"), 10, 64)
	return &model.Payment{
		id,
		name,
		model.EUR(amount),
		model.Date(date),
		projectId,
		r.categoryRepository.Find(categoryId),
		r.userRepository.Find(payerId).(*model.User),
		r.userRepository.Find(payeeId).(*model.User),
	}
}

func (r *paymentRepository) Find(id int64) interface{} {
	var result []model.Payment
	r.database.Select(r.ModelName(), &result, "*", r.NewFromQuery, "WHERE Id = "+strconv.FormatInt(id, 10))
	if len(result) > 0 {
		return &result[0]
	} else {
		return r.NewEmpty()
	}
}

func (r *paymentRepository) List(conditions ...string) interface{} {
	var result []model.Payment
	r.database.Select(r.ModelName(), &result, "*", r.NewFromQuery, conditions...)
	return result
}

func (r *paymentRepository) Count(conditions ...string) int64 {
	var result []int64
	r.database.Select(r.ModelName(), &result, "COUNT(*)", countQueryFunc, conditions...)
	return result[0]
}

func (r *paymentRepository) Insert(entity interface{}) {
	r.database.Insert(r.ModelName(), entity)
}

func (r *paymentRepository) Update(entity interface{}) {
	r.database.Update(r.ModelName(), entity)
}

func (r *paymentRepository) Delete(entity interface{}) {
	r.database.Delete(r.ModelName(), entity)
}
