package controller

import (
	"Bilance/model"
	"Bilance/repository"
	"Bilance/service"
	"Bilance/static"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type expenses struct {
	payments   repository.GRepository[model.Payment]
	categories repository.GRepository[model.Category]
}

func Expenses(payments repository.GRepository[model.Payment], categories repository.GRepository[model.Category]) Controller {
	return &expenses{payments, categories}
}

func (c *expenses) Routing(router service.Router) {
	router.Get("/expenses/by_period/", c.Expenses)
	router.Get("/expenses/by_category/", c.Expenses)
}

func (c *expenses) Expenses(writer http.ResponseWriter, request *http.Request) {
	var title string
	switch request.URL.Path {
	case "/expenses/by_period/":
		title = static.Translate("expenses") + " " + static.Translate("by_period")
	case "/expenses/by_category/":
		title = static.Translate("expenses") + " " + static.Translate("by_category")
	}
	render(
		writer,
		request,
		&Parameters{Model: c.prepareGraphData(request), Data: c.prepareYears()},
		title,
		"expenses",
	)
}

type GraphData struct {
	Filter string
	Type   string
	X      []string
	Y      []model.EUR
	Z      []string
	Total  model.EUR
}

const primaryColor = "#007bff"
const neutralColor = "#6c757d"

func (c *expenses) prepareGraphData(request *http.Request) *GraphData {
	var graphData = GraphData{}
	projectId := model.GetSelectedProjectIdString(request)
	start, end, step := c.calculateBoundaries(request, &graphData)
	switch request.URL.Path {
	case "/expenses/by_period/":
		graphData.Type = "bar"
		c.fillExpensesByPeriodGraphData(start, end, step, &graphData, projectId)
	case "/expenses/by_category/":
		graphData.Type = "doughnut"
		categories := c.categories.List("WHERE ProjectId = " + projectId)
		categories = append(categories, model.Category{0, static.Translate("uncategorized"), neutralColor, 0})
		c.fillExpensesByCategoryGraphData(start, end, categories, &graphData, projectId)
	}
	return &graphData
}

func (c *expenses) calculateBoundaries(request *http.Request, data *GraphData) (time.Time, time.Time, model.TimeUnit) {
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

func (c *expenses) fillExpensesByPeriodGraphData(start time.Time, end time.Time, step model.TimeUnit, data *GraphData, projectId string) {
	switch step {
	case model.TimeUnitMonth:
		for i := start.Month(); i <= end.Month(); i++ {
			t := start.AddDate(0, int(i)-1, 0)
			data.X = append(data.X, static.Translate(t.Month().String()))
			y := model.SumAmounts(c.payments.List(
				"WHERE ProjectId = "+projectId,
				"AND PayeeId = 0",
				"AND Date LIKE '"+t.Format("2006-01")+"%'",
			))
			data.Y = append(data.Y, y)
			data.Z = append(data.Z, primaryColor)
			data.Total += y
		}
	case model.TimeUnitMonthday:
		for i := start.Day(); i <= end.Day(); i++ {
			t := start.AddDate(0, 0, i-1)
			data.X = append(data.X, t.Format(model.DateLayoutDE))
			y := model.SumAmounts(c.payments.List(
				"WHERE ProjectId = "+projectId,
				"AND PayeeId = 0",
				"AND Date = '"+t.Format(model.DateLayoutISO)+"'",
			))
			data.Y = append(data.Y, y)
			data.Z = append(data.Z, primaryColor)
			data.Total += y
		}
	case model.TimeUnitWeekday:
		for i := model.NormalWeekday(start.Weekday()); i <= model.NormalWeekday(end.Weekday()); i++ {
			t := start.AddDate(0, 0, i)
			data.X = append(data.X, static.Translate(t.Weekday().String()))
			y := model.SumAmounts(c.payments.List(
				"WHERE ProjectId = "+projectId,
				"AND PayeeId = 0",
				"AND Date = '"+t.Format(model.DateLayoutISO)+"'",
			))
			data.Y = append(data.Y, y)
			data.Z = append(data.Z, primaryColor)
			data.Total += y
		}
	}
}

func (c *expenses) fillExpensesByCategoryGraphData(start time.Time, end time.Time, categories []model.Category, data *GraphData, projectId string) {
	startDate := start.Format(model.DateLayoutISO)
	endDate := end.Format(model.DateLayoutISO)
	for _, category := range categories {
		data.X = append(data.X, category.Name)
		var y model.EUR
		if category.Id == 0 {
			y = model.SumAmounts(c.payments.List(
				"WHERE ProjectId = "+projectId,
				"AND PayeeId = 0",
				"AND CategoryId NOT IN ("+ExtractCategoryIds(categories)+")",
				"AND Date BETWEEN '"+startDate+"' AND '"+endDate+"'",
			))
		} else {
			y = model.SumAmounts(c.payments.List(
				"WHERE ProjectId = "+projectId,
				"AND PayeeId = 0",
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

func ExtractCategoryIds(categories []model.Category) string {
	var result []string
	for _, category := range categories {
		result = append(result, strconv.FormatInt(category.Id, 10))
	}
	return strings.Join(result, ",")
}
