package controller

import (
	"Bilance/model"
	"Bilance/repository"
	"Bilance/server"
	"math"
	"net/http"
	"strconv"
)

type balances struct {
	projects repository.Repository[model.Project]
	payments repository.Repository[model.Payment]
	users    repository.Repository[model.User]
}

func Balances(
	projects repository.Repository[model.Project],
	payments repository.Repository[model.Payment],
	users repository.Repository[model.User],
) Controller {
	return &balances{projects, payments, users}
}

func (b *balances) Routing(server server.Server) {
	server.Get("/balances/", b.Balances)
}

func (b *balances) Balances(writer http.ResponseWriter, request *http.Request) {
	balanceData := b.prepareBalanceData(request)
	render(
		writer,
		request,
		&Parameters{Model: &balanceData},
		"balances",
		"balances",
	)
}

func (b *balances) prepareBalanceData(request *http.Request) []*BalanceData {
	var balanceData []*BalanceData
	project := b.projects.Find(model.GetSelectedProjectId(request))
	projectIdString := model.GetSelectedProjectIdString(request)
	totalExpenses := -model.SumAmounts(b.payments.List(
		"ProjectId = "+projectIdString,
		"AND PayeeId = 0",
	))
	userIds := model.StringToIds(project.UserIds)
	userAmount := len(userIds)
	proportionalExpenses := model.EUR(int64(totalExpenses) / int64(userAmount))
	var maxBalance float64
	for _, userId := range userIds {
		userIdString := strconv.FormatInt(userId, 10)
		sentExpenses := model.SumAmounts(b.payments.List(
			"ProjectId = "+projectIdString,
			"AND PayerId = "+userIdString,
			"AND PayeeId = 0",
		))
		sentTransfer := model.SumAmounts(b.payments.List(
			"ProjectId = "+projectIdString,
			"AND PayerId = "+userIdString,
			"AND PayeeId != 0",
		))
		receivedTransfer := model.SumAmounts(b.payments.List(
			"ProjectId = "+projectIdString,
			"AND PayerId != "+userIdString,
			"AND PayeeId = "+userIdString,
		))
		result := proportionalExpenses + sentExpenses + sentTransfer - receivedTransfer
		if math.Abs(float64(result)) > float64(maxBalance) {
			maxBalance = float64(result)
		}
		balanceData = append(balanceData, &BalanceData{
			ProjectName:          project.Name,
			UserName:             b.users.Find(userId).Name,
			TotalExpenses:        totalExpenses,
			UserAmount:           userAmount,
			ProportionalExpenses: proportionalExpenses,
			SentExpenses:         sentExpenses,
			SentTransfer:         sentTransfer,
			ReceivedTransfer:     receivedTransfer,
			Result:               result,
		})
	}
	b.calculateWidths(balanceData, maxBalance)
	return balanceData
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

func (b *balances) calculateWidths(balanceData []*BalanceData, maxBalance float64) {
	for _, data := range balanceData {
		if data.Result > 0 {
			data.Width = float64(data.Result) / maxBalance * 100
		} else {
			data.Width = -float64(data.Result) / maxBalance * 100
		}
	}
}
