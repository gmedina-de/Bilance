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
	baseRepository
	userRepository Repository
	typeRepository Repository
}

func PaymentRepository(database service.Database, userRepository Repository, typeRepository Repository) Repository {
	return &paymentRepository{baseRepository{database: database}, userRepository, typeRepository}
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
	ScanAndPanic(row, &id, &name, &amount, &date, &projectId, &categoryId, &payerId, &payeeId)
	category := r.typeRepository.Find(categoryId).(*model.Category)
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
		r.typeRepository.Find(categoryId).(*model.Category),
		r.userRepository.Find(payerId).(*model.User),
		r.userRepository.Find(payeeId).(*model.User),
	}
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
