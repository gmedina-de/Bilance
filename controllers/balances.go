package controllers

import (
	"genuine/models"
	"genuine/repositories"
	"math"
)

type balances struct {
	payments repositories.Repository[models.Payment]
	users    repositories.Repository[models.User]
}

func Balances(payments repositories.Repository[models.Payment], users repositories.Repository[models.User]) Controller {
	return &balances{payments: payments, users: users}
}

func (b *balances) Routes() map[string]Handler {
	return map[string]Handler{
		"GET /accounting/balances": b.Balances,
	}
}

func (b *balances) Balances(Request) Response {
	balanceData := b.prepareBalanceData()
	return Response{
		"Template":    "balances",
		"BalanceData": balanceData,
	}
}

func (b *balances) prepareBalanceData() []*BalanceData {
	var balanceData []*BalanceData
	totalExpenses := -models.SumAmounts(b.payments.List("payee_id = 0"))
	users := b.users.All()
	userAmount := len(users)
	proportionalExpenses := models.Currency(int64(totalExpenses) / int64(userAmount))
	var maxBalance float64
	for _, user := range users {
		userID := user.ID
		sentExpenses := models.SumAmounts(b.payments.List("payer_id = ? AND payee_id = 0", userID))
		sentTransfer := models.SumAmounts(b.payments.List("payer_id = ? AND payee_id != 0", userID))
		receivedTransfer := models.SumAmounts(b.payments.List("payer_id != ? AND payee_id = ?", userID, userID))
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
	TotalExpenses        models.Currency
	UserAmount           int
	ProportionalExpenses models.Currency
	SentExpenses         models.Currency
	SentTransfer         models.Currency
	ReceivedTransfer     models.Currency
	Result               models.Currency
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
