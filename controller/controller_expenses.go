package controller

import (
	"Bilance/localization"
	"Bilance/model"
	"Bilance/repository"
	"Bilance/service"
	"net/http"
	"strconv"
	"time"
)

type expensesController struct {
	paymentRepository  repository.Repository
	categoryRepository repository.Repository
}

func ExpensesController(paymentRepository repository.Repository, categoryRepository repository.Repository) Controller {
	return &expensesController{paymentRepository, categoryRepository}
}

func (c *expensesController) Routing(router service.Router) {
	router.Get("/expenses/", c.Expenses)
	router.Get("/by_category/", c.ByCategory)
}

func (c *expensesController) Expenses(writer http.ResponseWriter, request *http.Request) {
	render(
		writer,
		request,
		&Parameters{
			Model: c.prepareExpensesGraphData(request),
		},
		"expenses",
		"expenses",
	)
}

func (c *expensesController) ByCategory(writer http.ResponseWriter, request *http.Request) {
	render(
		writer,
		request,
		&Parameters{
			Model: c.prepareCategoryGraphData(request),
		},
		"by_category",
		"expenses",
	)
}

type GraphData struct {
	GraphDataType string
	GraphDataX    []string
	GraphDataY    []model.EUR
	GraphDataZ    []string
	Total         model.EUR
}

func (c *expensesController) prepareExpensesGraphData(request *http.Request) *GraphData {
	var graphDataType = "bar"
	var graphDataX []string
	var graphDataY []model.EUR
	var graphDataZ []string
	var total model.EUR

	projectId := model.GetSelectedProjectIdString(request)
	location, _ := time.LoadLocation("Europe/Berlin")

	var start int
	//var step time.Duration
	//get := request.URL.Query().Get("filter")
	//switch get {
	//case "this_month":
	//	start = -30
	//	step = time.
	//case "this_week":
	//	start = -6
	//case "":
	//	start = -6
	//}

	now := time.Now().In(location)
	for i := start; i <= 0; i++ {
		day := now.AddDate(0, 0, i)

		graphDataX = append(graphDataX, localization.Translate(day.Weekday().String()))

		y := c.sumExpenses(c.paymentRepository.List(
			"WHERE ProjectId = "+projectId,
			"AND Date = '"+day.Format(model.DateLayoutISO)+"'",
		).([]model.Payment))
		graphDataY = append(graphDataY, y)
		graphDataZ = append(graphDataZ, "#007bff")
		total += y
	}

	return &GraphData{graphDataType, graphDataX, graphDataY, graphDataZ, total}
}

func (c *expensesController) prepareCategoryGraphData(request *http.Request) *GraphData {
	var graphDataType = "doughnut"
	var graphDataX []string
	var graphDataY []model.EUR
	var graphDataZ []string
	var total model.EUR

	projectId := model.GetSelectedProjectIdString(request)
	categories := c.categoryRepository.List("WHERE ProjectId = " + projectId).([]model.Category)

	for _, category := range categories {
		graphDataX = append(graphDataX, category.Name)
		y := c.sumExpenses(c.paymentRepository.List(
			"WHERE ProjectId = "+projectId,
			"AND CategoryId = '"+strconv.FormatInt(category.Id, 10)+"'",
		).([]model.Payment))
		graphDataY = append(graphDataY, y)
		graphDataZ = append(graphDataZ, category.Color)
		total += y

	}

	return &GraphData{graphDataType, graphDataX, graphDataY, graphDataZ, total}
}

func (c *expensesController) sumExpenses(payments []model.Payment) model.EUR {
	var result model.EUR
	for _, payment := range payments {
		result += payment.Amount
	}
	return result
}
