package controllers

import (
	model2 "genuine/app/models"
	"genuine/core/controllers"
	"genuine/core/repositories"
	"math"
)

type balances struct {
	payments repositories.Repository[model2.Payment]
	users    repositories.Repository[model2.User]
}

func Balances(payments repositories.Repository[model2.Payment], users repositories.Repository[model2.User]) controllers.Controller {
	return &balances{payments: payments, users: users}
}

func (b *balances) Routes() map[string]controllers.Handler {
	return map[string]controllers.Handler{
		"GET /accounting/balances": b.Balances,
	}
}

func (b *balances) Balances(controllers.Request) controllers.Response {
	balanceData := b.prepareBalanceData()
	return controllers.Response{
		"Template":    "balances",
		"BalanceData": balanceData,
	}
}

func (b *balances) prepareBalanceData() []*BalanceData {
	var balanceData []*BalanceData
	totalExpenses := -model2.SumAmounts(b.payments.List("payee_id = 0"))
	users := b.users.All()
	userAmount := len(users)
	proportionalExpenses := model2.Currency(int64(totalExpenses) / int64(userAmount))
	var maxBalance float64
	for _, user := range users {
		userID := user.ID
		sentExpenses := model2.SumAmounts(b.payments.List("payer_id = ? AND payee_id = 0", userID))
		sentTransfer := model2.SumAmounts(b.payments.List("payer_id = ? AND payee_id != 0", userID))
		receivedTransfer := model2.SumAmounts(b.payments.List("payer_id != ? AND payee_id = ?", userID, userID))
		result := proportionalExpenses + sentExpenses + sentTransfer - receivedTransfer
		if math.Abs(float64(result)) > float64(maxBalance) {
			maxBalance = float64(result)
		}
		balanceData = append(balanceData, &BalanceData{
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
	b.calculateWidths(balanceData, maxBalance)
	return balanceData
}

type BalanceData struct {
	ProjectName          string
	UserName             string
	TotalExpenses        model2.Currency
	UserAmount           int
	ProportionalExpenses model2.Currency
	SentExpenses         model2.Currency
	SentTransfer         model2.Currency
	ReceivedTransfer     model2.Currency
	Result               model2.Currency
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
