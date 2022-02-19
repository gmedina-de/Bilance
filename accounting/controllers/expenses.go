package controllers

import (
	models3 "genuine/accounting/models"
	"genuine/core/controllers"
	"genuine/core/inject"
	"genuine/core/models"
	"genuine/core/repositories"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type expenses struct {
	controllers.Base
	Payments   repositories.Repository[models3.Payment]
	Categories repositories.Repository[models3.Category]
}

func Expenses() controllers.Controller {
	return inject.Inject(&expenses{})
}

func (c *expenses) Routes() map[string]string {
	return map[string]string{}
}

func (c *expenses) Expenses() {
	//var title string
	//switch request.URL.Path {
	//case "/expenses/by_period/":
	//	title = localization.Translate("expenses") + " " + localization.Translate("by_period")
	//case "/expenses/by_category/":
	//	title = localization.Translate("expenses") + " " + localization.Translate("by_category")
	//}
	//template.Render(
	//	writer,
	//	request,
	//	title,
	//	&template.Parameters{models: c.prepareGraphData(request), Data: c.prepareYears()},
	//	"accounting/template/expenses.gohtml",
	//)
}

type GraphData struct {
	Filter string
	Type   string
	X      []string
	Y      []models3.EUR
	Z      []string
	Total  models3.EUR
}

const primaryColor = "#007bff"
const neutralColor = "#6c757d"

func (c *expenses) prepareGraphData(request *http.Request) *GraphData {
	var graphData = GraphData{}
	start, end, step := c.calculateBoundaries(request, &graphData)
	switch request.URL.Path {
	case "/expenses/by_period/":
		graphData.Type = "bar"
		c.fillExpensesByPeriodGraphData(start, end, step, &graphData)
	case "/expenses/by_category/":
		graphData.Type = "doughnut"
		categories := c.Categories.All()
		categories = append(categories, models3.Category{0, "uncategorized", neutralColor})
		c.fillExpensesByCategoryGraphData(start, end, categories, &graphData)
	}
	return &graphData
}

func (c *expenses) calculateBoundaries(request *http.Request, data *GraphData) (time.Time, time.Time, models.TimeUnit) {
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
			y := models3.SumAmounts(c.Payments.List(
				"PayeeId = 0",
				"AND Date LIKE '"+t.Format("2006-01")+"%'",
			))
			data.Y = append(data.Y, y)
			data.Z = append(data.Z, primaryColor)
			data.Total += y
		}
	case models.TimeUnitMonthday:
		for i := start.Day(); i <= end.Day(); i++ {
			t := start.AddDate(0, 0, i-1)
			data.X = append(data.X, t.Format(models.DateLayoutDE))
			y := models3.SumAmounts(c.Payments.List(
				"PayeeId = 0",
				"AND Date = '"+t.Format(models.DateLayoutISO)+"'",
			))
			data.Y = append(data.Y, y)
			data.Z = append(data.Z, primaryColor)
			data.Total += y
		}
	case models.TimeUnitWeekday:
		for i := models.NormalWeekday(start.Weekday()); i <= models.NormalWeekday(end.Weekday()); i++ {
			t := start.AddDate(0, 0, i)
			data.X = append(data.X, t.Weekday().String())
			y := models3.SumAmounts(c.Payments.List(
				"PayeeId = 0",
				"AND Date = '"+t.Format(models.DateLayoutISO)+"'",
			))
			data.Y = append(data.Y, y)
			data.Z = append(data.Z, primaryColor)
			data.Total += y
		}
	}
}

func (c *expenses) fillExpensesByCategoryGraphData(start time.Time, end time.Time, categories []models3.Category, data *GraphData) {
	startDate := start.Format(models.DateLayoutISO)
	endDate := end.Format(models.DateLayoutISO)
	for _, category := range categories {
		data.X = append(data.X, category.Name)
		var y models3.EUR
		if category.Id == 0 {
			y = models3.SumAmounts(c.Payments.List(
				"PayeeId = 0",
				"AND CategoryId NOT IN ("+ExtractCategoryIds(categories)+")",
				"AND Date BETWEEN '"+startDate+"' AND '"+endDate+"'",
			))
		} else {
			y = models3.SumAmounts(c.Payments.List(
				"PayeeId = 0",
				"AND CategoryId = '"+strconv.FormatInt(category.Id, 10)+"'",
				"AND Date BETWEEN '"+startDate+"' AND '"+endDate+"'",
			))
		}
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

func ExtractCategoryIds(categories []models3.Category) string {
	var result []string
	for _, category := range categories {
		result = append(result, strconv.FormatInt(category.Id, 10))
	}
	return strings.Join(result, ",")
}
