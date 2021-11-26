package controller

import (
	"Bilance/localization"
	"Bilance/model"
	"Bilance/repository"
	"Bilance/service"
	"net/http"
	"strconv"
	"strings"
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
	router.Get("/by_category/", c.Expenses)
}

func (c *expensesController) Expenses(writer http.ResponseWriter, request *http.Request) {
	render(writer, request, &Parameters{Model: c.prepareGraphData(request), Data: c.prepareYears()}, "expenses", "expenses")
}

type GraphData struct {
	Filter string
	Type   string
	X      []string
	Y      []model.EUR
	Z      []string
	Total  model.EUR
}

const neutralColor = "#6c757d"

func (c *expensesController) prepareGraphData(request *http.Request) *GraphData {
	var graphData = GraphData{}
	projectId := model.GetSelectedProjectIdString(request)
	start, end, step := c.calculateBoundaries(request, &graphData)
	if strings.HasPrefix(request.URL.Path, "/expenses/") {
		graphData.Type = "bar"
		c.fillExpensesGraphData(start, end, step, &graphData, projectId)
	} else if strings.HasPrefix(request.URL.Path, "/by_category/") {
		graphData.Type = "doughnut"
		categories := c.categoryRepository.List("WHERE ProjectId = " + projectId).([]model.Category)
		categories = append(categories, model.Category{0, localization.Translate("uncategorized"), neutralColor, 0})
		c.fillByCategoryGraphData(start, end, categories, &graphData, projectId)
	}

	return &graphData
}

func (c *expensesController) calculateBoundaries(request *http.Request, data *GraphData) (time.Time, time.Time, model.TimeUnit) {
	location, _ := time.LoadLocation("Europe/Berlin")
	now := time.Now().In(location)
	var start time.Time
	var end time.Time
	var step model.TimeUnit
	data.Filter = request.URL.Query().Get("filter")
	if strings.HasPrefix(data.Filter, "year") {
		yearString := strings.Replace(data.Filter, "year", "", 1)
		data.Filter = yearString
		year, _ := strconv.Atoi(yearString)
		start = time.Date(year, time.January, 1, 0, 0, 0, 0, location)
		step = model.TimeUnitMonth
		end = time.Date(year, time.December, 31, 0, 0, 0, 0, location)
	} else {
		switch data.Filter {
		case "this_year":
			start = time.Date(now.Year(), 1, 1, 0, 0, 0, 0, location)
			step = model.TimeUnitMonth
		case "this_month":
			start = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, location)
			step = model.TimeUnitMonthday
		case "this_week":
			fallthrough
		case "":
			data.Filter = "this_week"
			start = time.Date(now.Year(), now.Month(), now.Day()-model.NormalWeekday(now.Weekday()), 0, 0, 0, 0, location)
			step = model.TimeUnitWeekday
		}
		end = now
	}
	return start, end, step
}

func (c *expensesController) fillExpensesGraphData(start time.Time, end time.Time, step model.TimeUnit, data *GraphData, projectId string) {
	switch step {
	case model.TimeUnitMonth:
		for i := start.Month(); i <= end.Month(); i++ {
			t := start.AddDate(0, int(i)-1, 0)
			data.X = append(data.X, localization.Translate(t.Month().String()))
			y := c.sumExpenses(c.paymentRepository.List(
				"WHERE ProjectId = "+projectId,
				"AND Date LIKE '"+t.Format("2006-01")+"%'",
			).([]model.Payment))
			data.Y = append(data.Y, y)
			data.Z = append(data.Z, neutralColor)
			data.Total += y
		}
	case model.TimeUnitMonthday:
		for i := start.Day(); i <= end.Day(); i++ {
			t := start.AddDate(0, 0, i-1)
			data.X = append(data.X, t.Format(model.DateLayoutDE))
			y := c.sumExpenses(c.paymentRepository.List(
				"WHERE ProjectId = "+projectId,
				"AND Date = '"+t.Format(model.DateLayoutISO)+"'",
			).([]model.Payment))
			data.Y = append(data.Y, y)
			data.Z = append(data.Z, neutralColor)
			data.Total += y
		}
	case model.TimeUnitWeekday:
		for i := model.NormalWeekday(start.Weekday()); i <= model.NormalWeekday(end.Weekday()); i++ {
			t := start.AddDate(0, 0, i)
			data.X = append(data.X, localization.Translate(t.Weekday().String()))
			y := c.sumExpenses(c.paymentRepository.List(
				"WHERE ProjectId = "+projectId,
				"AND Date = '"+t.Format(model.DateLayoutISO)+"'",
			).([]model.Payment))
			data.Y = append(data.Y, y)
			data.Z = append(data.Z, neutralColor)
			data.Total += y
		}
	}
}

func (c *expensesController) fillByCategoryGraphData(start time.Time, end time.Time, categories []model.Category, data *GraphData, projectId string) {
	startDate := start.Format(model.DateLayoutISO)
	endDate := end.Format(model.DateLayoutISO)
	for _, category := range categories {
		data.X = append(data.X, category.Name)
		var y model.EUR
		if category.Id == 0 {
			y = c.sumExpenses(c.paymentRepository.List(
				"WHERE ProjectId = "+projectId,
				"AND CategoryId NOT IN ("+extractCategoryIds(categories)+")",
				"AND Date BETWEEN '"+startDate+"' AND '"+endDate+"'",
			).([]model.Payment))
		} else {
			y = c.sumExpenses(c.paymentRepository.List(
				"WHERE ProjectId = "+projectId,
				"AND CategoryId = '"+strconv.FormatInt(category.Id, 10)+"'",
				"AND Date BETWEEN '"+startDate+"' AND '"+endDate+"'",
			).([]model.Payment))
		}
		data.Y = append(data.Y, y)
		data.Z = append(data.Z, category.Color)
		data.Total += y
	}
}

func extractCategoryIds(categories []model.Category) string {
	var result []string
	for _, category := range categories {
		result = append(result, strconv.FormatInt(category.Id, 10))
	}
	return strings.Join(result, ",")
}

func (c *expensesController) sumExpenses(payments []model.Payment) model.EUR {
	var result model.EUR
	for _, payment := range payments {
		result += payment.Amount
	}
	return result
}

func (c *expensesController) prepareYears() []int {
	var result []int
	currentYear := time.Now().Year()
	for i := 1; i < 11; i++ {
		result = append(result, currentYear-i)
	}
	return result
}
