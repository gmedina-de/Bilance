package controller

import (
	"Bilance/model"
	"Bilance/repository"
	"Bilance/service"
	"math"
	"net/http"
	"strconv"
)

type balancesController struct {
	projectRepository repository.Repository
	paymentRepository repository.Repository
}

func BalancesController(projectRepository repository.Repository, paymentRepository repository.Repository) Controller {
	return &balancesController{projectRepository, paymentRepository}
}

func (c *balancesController) Routing(router service.Router) {
	router.Get("/balances/", c.Balances)
}

type BalanceData struct {
	ProjectName          string
	UserName             string
	TotalExpenses        model.EUR
	UserAmount           int
	ProportionalExpenses model.EUR
	SentExpenses         model.EUR
	SentTransfer         model.EUR
	ReceivedTransfer     model.EUR
	Result               model.EUR
	Width                float64
}

func (c *balancesController) Balances(writer http.ResponseWriter, request *http.Request) {
	balanceData := c.prepareBalanceData(request)
	render(
		writer,
		request,
		&Parameters{Model: &balanceData},
		"balances",
		"balances",
	)
}

func (c *balancesController) prepareBalanceData(request *http.Request) []*BalanceData {
	var balanceData []*BalanceData
	project := c.projectRepository.Find(model.GetSelectedProjectId(request)).(*model.Project)
	projectIdString := model.GetSelectedProjectIdString(request)
	totalExpenses := -model.SumAmounts(c.paymentRepository.List(
		"WHERE ProjectId = "+projectIdString,
		"AND PayeeId = 0",
	).([]model.Payment))
	users := project.Users
	userAmount := len(users)
	proportionalExpenses := model.EUR(int64(totalExpenses) / int64(userAmount))
	var maxBalance float64
	for _, user := range users {
		userIdString := strconv.FormatInt(user.Id, 10)
		sentExpenses := model.SumAmounts(c.paymentRepository.List(
			"WHERE ProjectId = "+projectIdString,
			"AND PayerId = "+userIdString,
			"AND PayeeId = 0",
		).([]model.Payment))
		sentTransfer := model.SumAmounts(c.paymentRepository.List(
			"WHERE ProjectId = "+projectIdString,
			"AND PayerId = "+userIdString,
			"AND PayeeId != 0",
		).([]model.Payment))
		receivedTransfer := model.SumAmounts(c.paymentRepository.List(
			"WHERE ProjectId = "+projectIdString,
			"AND PayerId != "+userIdString,
			"AND PayeeId = "+userIdString,
		).([]model.Payment))
		result := proportionalExpenses + sentExpenses + sentTransfer - receivedTransfer
		if math.Abs(float64(result)) > float64(maxBalance) {
			maxBalance = float64(result)
		}
		balanceData = append(balanceData, &BalanceData{
			ProjectName:          project.Name,
			UserName:             user.Name,
			TotalExpenses:        totalExpenses,
			UserAmount:           userAmount,
			ProportionalExpenses: proportionalExpenses,
			SentExpenses:         sentExpenses,
			SentTransfer:         sentTransfer,
			ReceivedTransfer:     receivedTransfer,
			Result:               result,
		})
	}
	c.calculateWidths(balanceData, maxBalance)
	return balanceData
}

func (c *balancesController) calculateWidths(balanceData []*BalanceData, maxBalance float64) {
	for _, data := range balanceData {
		if data.Result > 0 {
			data.Width = float64(data.Result) / maxBalance * 100
		} else {
			data.Width = -float64(data.Result) / maxBalance * 100
		}
	}
}
