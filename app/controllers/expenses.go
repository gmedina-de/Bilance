package controllers

import (
	"genuine/app/models"
	"genuine/core/controllers"
	"genuine/core/repositories"
	"strconv"
	"strings"
	"time"
)

type expenses struct {
	payments   repositories.Repository[models.Payment]
	categories repositories.Repository[models.Category]
}

func Expenses(payments repositories.Repository[models.Payment], categories repositories.Repository[models.Category]) controllers.Controller {
	return &expenses{payments: payments, categories: categories}
}

func (c *expenses) Routes() map[string]controllers.Handler {
	return map[string]controllers.Handler{
		"GET /accounting/expenses/by_period":   c.ExpensesByPeriod,
		"GET /accounting/expenses/by_category": c.ExpensesByCategory,
	}
}

func (c *expenses) ExpensesByPeriod(request controllers.Request) controllers.Response {
	var graphData = GraphData{}
	start, end, step := c.calculateBoundaries(request, &graphData)
	graphData.Type = "bar"
	c.fillExpensesByPeriodGraphData(start, end, step, &graphData)
	return controllers.Response{
		"Template":  "expenses",
		"GraphData": graphData,
		"Years":     c.prepareYears(),
	}
}

func (c *expenses) ExpensesByCategory(request controllers.Request) controllers.Response {
	var graphData = GraphData{}
	start, end, _ := c.calculateBoundaries(request, &graphData)
	graphData.Type = "doughnut"
	categories := c.categories.All()
	categories = append(categories, models.Category{Name: "uncategorized", Color: neutralColor})
	c.fillExpensesByCategoryGraphData(start, end, categories, &graphData)
	return controllers.Response{
		"Template":  "expenses",
		"GraphData": graphData,
		"Years":     c.prepareYears(),
	}
}

type GraphData struct {
	Filter string
	Type   string
	X      []string
	Y      []models.Currency
	Z      []string
	Total  models.Currency
}

const primaryColor = "#007bff"
const neutralColor = "#6c757d"

func (c *expenses) calculateBoundaries(request controllers.Request, data *GraphData) (time.Time, time.Time, models.TimeUnit) {
	location, _ := time.LoadLocation("Europe/Berlin")
	now := time.Now().In(location)
	var start time.Time
	var end time.Time
	var step models.TimeUnit
	data.Filter = request.URL.Query().Get("filter")
	if strings.HasPrefix(data.Filter, "year") {
		yearString := strings.Replace(data.Filter, "year", "", 1)
		data.Filter = yearString
		year, _ := strconv.Atoi(yearString)
		start = time.Date(year, time.January, 1, 0, 0, 0, 0, location)
		step = models.TimeUnitMonth
		end = time.Date(year, time.December, 31, 0, 0, 0, 0, location)
	} else {
		switch data.Filter {
		case "this_year":
			start = time.Date(now.Year(), 1, 1, 0, 0, 0, 0, location)
			step = models.TimeUnitMonth
		case "this_month":
			start = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, location)
			step = models.TimeUnitMonthday
		case "this_week":
			fallthrough
		case "":
			data.Filter = "this_week"
			start = time.Date(now.Year(), now.Month(), now.Day()-models.NormalWeekday(now.Weekday()), 0, 0, 0, 0, location)
			step = models.TimeUnitWeekday
		}
		end = now
	}
	return start, end, step
}

func (c *expenses) fillExpensesByPeriodGraphData(start time.Time, end time.Time, step models.TimeUnit, data *GraphData) {
	switch step {
	case models.TimeUnitMonth:
		for i := start.Month(); i <= end.Month(); i++ {
			t := start.AddDate(0, int(i)-1, 0)
			data.X = append(data.X, t.Month().String())
			y := models.SumAmounts(c.payments.List("payee_id = 0 AND date LIKE ?", t.Format("2006-01")+"%"))
			data.Y = append(data.Y, y)
			data.Z = append(data.Z, primaryColor)
			data.Total += y
		}
	case models.TimeUnitMonthday:
		for i := start.Day(); i <= end.Day(); i++ {
			t := start.AddDate(0, 0, i-1)
			data.X = append(data.X, t.Format(models.DateLayoutDE))
			y := models.SumAmounts(c.payments.List("payee_id = 0 AND date = ?", t.Format(models.DateLayoutISO)))
			data.Y = append(data.Y, y)
			data.Z = append(data.Z, primaryColor)
			data.Total += y
		}
	case models.TimeUnitWeekday:
		for i := models.NormalWeekday(start.Weekday()); i <= models.NormalWeekday(end.Weekday()); i++ {
			t := start.AddDate(0, 0, i)
			data.X = append(data.X, t.Weekday().String())
			y := models.SumAmounts(c.payments.List("payee_id = 0 AND date = ?", t.Format(models.DateLayoutISO)))
			data.Y = append(data.Y, y)
			data.Z = append(data.Z, primaryColor)
			data.Total += y
		}
	}
}

func (c *expenses) fillExpensesByCategoryGraphData(start time.Time, end time.Time, categories []models.Category, data *GraphData) {
	startDate := start.Format(models.DateLayoutISO)
	endDate := end.Format(models.DateLayoutISO)
	for _, category := range categories {
		data.X = append(data.X, category.Name)
		var y = models.SumAmounts(c.payments.List("payee_id = 0 AND category_id = ? AND date BETWEEN ? AND ?",
			category.ID,
			startDate,
			endDate,
		))
		data.Y = append(data.Y, y)
		data.Z = append(data.Z, category.Color)
		data.Total += y
	}
}

func (c *expenses) prepareYears() []int {
	var result []int
	currentYear := time.Now().Year()
	for i := 1; i < 11; i++ {
		result = append(result, currentYear-i)
	}
	return result
}

func extractCategoryIds(categories []models.Category) string {
	var result []string
	for _, category := range categories {
		result = append(result, strconv.FormatUint(uint64(category.ID), 10))
	}
	return strings.Join(result, ",")
}
