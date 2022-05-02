package controllers

import (
	models2 "genuine/models"
	"genuine/repositories"
	"strconv"
	"strings"
	"time"
)

type expenses struct {
	payments   repositories.Repository[models2.Payment]
	categories repositories.Repository[models2.Category]
}

func Expenses(payments repositories.Repository[models2.Payment], categories repositories.Repository[models2.Category]) Controller {
	return &expenses{payments: payments, categories: categories}
}

func (c *expenses) Routes() map[string]Handler {
	return map[string]Handler{
		"GET /accounting/expenses/by_period":   c.ExpensesByPeriod,
		"GET /accounting/expenses/by_category": c.ExpensesByCategory,
	}
}

func (c *expenses) ExpensesByPeriod(request Request) Response {
	var graphData = GraphData{}
	start, end, step := c.calculateBoundaries(request, &graphData)
	graphData.Type = "bar"
	c.fillExpensesByPeriodGraphData(start, end, step, &graphData)
	return Response{
		"Template":  "expenses",
		"GraphData": graphData,
		"Years":     c.prepareYears(),
	}
}

func (c *expenses) ExpensesByCategory(request Request) Response {
	var graphData = GraphData{}
	start, end, _ := c.calculateBoundaries(request, &graphData)
	graphData.Type = "doughnut"
	categories := c.categories.All()
	categories = append(categories, models2.Category{Name: "uncategorized", Color: neutralColor})
	c.fillExpensesByCategoryGraphData(start, end, categories, &graphData)
	return Response{
		"Template":  "expenses",
		"GraphData": graphData,
		"Years":     c.prepareYears(),
	}
}

type GraphData struct {
	Filter string
	Type   string
	X      []string
	Y      []models2.Currency
	Z      []string
	Total  models2.Currency
}

const primaryColor = "#007bff"
const neutralColor = "#6c757d"

func (c *expenses) calculateBoundaries(request Request, data *GraphData) (time.Time, time.Time, models2.TimeUnit) {
	location, _ := time.LoadLocation("Europe/Berlin")
	now := time.Now().In(location)
	var start time.Time
	var end time.Time
	var step models2.TimeUnit
	data.Filter = request.URL.Query().Get("filter")
	if strings.HasPrefix(data.Filter, "year") {
		yearString := strings.Replace(data.Filter, "year", "", 1)
		data.Filter = yearString
		year, _ := strconv.Atoi(yearString)
		start = time.Date(year, time.January, 1, 0, 0, 0, 0, location)
		step = models2.TimeUnitMonth
		end = time.Date(year, time.December, 31, 0, 0, 0, 0, location)
	} else {
		switch data.Filter {
		case "this_year":
			start = time.Date(now.Year(), 1, 1, 0, 0, 0, 0, location)
			step = models2.TimeUnitMonth
		case "this_month":
			start = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, location)
			step = models2.TimeUnitMonthday
		case "this_week":
			fallthrough
		case "":
			data.Filter = "this_week"
			start = time.Date(now.Year(), now.Month(), now.Day()-models2.NormalWeekday(now.Weekday()), 0, 0, 0, 0, location)
			step = models2.TimeUnitWeekday
		}
		end = now
	}
	return start, end, step
}

func (c *expenses) fillExpensesByPeriodGraphData(start time.Time, end time.Time, step models2.TimeUnit, data *GraphData) {
	switch step {
	case models2.TimeUnitMonth:
		for i := start.Month(); i <= end.Month(); i++ {
			t := start.AddDate(0, int(i)-1, 0)
			data.X = append(data.X, t.Month().String())
			y := models2.SumAmounts(c.payments.List("payee_id = 0 AND date LIKE ?", t.Format("2006-01")+"%"))
			data.Y = append(data.Y, y)
			data.Z = append(data.Z, primaryColor)
			data.Total += y
		}
	case models2.TimeUnitMonthday:
		for i := start.Day(); i <= end.Day(); i++ {
			t := start.AddDate(0, 0, i-1)
			data.X = append(data.X, t.Format(models2.DateLayoutDE))
			y := models2.SumAmounts(c.payments.List("payee_id = 0 AND date = ?", t.Format(models2.DateLayoutISO)))
			data.Y = append(data.Y, y)
			data.Z = append(data.Z, primaryColor)
			data.Total += y
		}
	case models2.TimeUnitWeekday:
		for i := models2.NormalWeekday(start.Weekday()); i <= models2.NormalWeekday(end.Weekday()); i++ {
			t := start.AddDate(0, 0, i)
			data.X = append(data.X, t.Weekday().String())
			y := models2.SumAmounts(c.payments.List("payee_id = 0 AND date = ?", t.Format(models2.DateLayoutISO)))
			data.Y = append(data.Y, y)
			data.Z = append(data.Z, primaryColor)
			data.Total += y
		}
	}
}

func (c *expenses) fillExpensesByCategoryGraphData(start time.Time, end time.Time, categories []models2.Category, data *GraphData) {
	startDate := start.Format(models2.DateLayoutISO)
	endDate := end.Format(models2.DateLayoutISO)
	for _, category := range categories {
		data.X = append(data.X, category.Name)
		var y = models2.SumAmounts(c.payments.List("payee_id = 0 AND category_id = ? AND date BETWEEN ? AND ?",
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

func extractCategoryIds(categories []models2.Category) string {
	var result []string
	for _, category := range categories {
		result = append(result, strconv.FormatUint(uint64(category.ID), 10))
	}
	return strings.Join(result, ",")
}
