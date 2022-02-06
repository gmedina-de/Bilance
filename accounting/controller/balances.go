package controller

import (
	model2 "homecloud/accounting/model"
	"homecloud/core/controllers"
	"homecloud/core/model"
	"homecloud/core/repositories"
	"net/http"
)

type balances struct {
	controllers.BaseController
	payments repositories.Repository[model2.Payment]
	users    repositories.Repository[model.User]
}

func Balances(payments repositories.Repository[model2.Payment], users repositories.Repository[model.User]) controllers.Controller {
	return &balances{payments: payments, users: users}
}
func (b *balances) Balances(writer http.ResponseWriter, request *http.Request) {
	//balanceData := b.prepareBalanceData()
	//template.Render(writer, request, "balances", &template.Parameters{model: &balanceData}, "accounting/template/balances.gohtml")
	return
}

func (b *balances) prepareBalanceData() []*BalanceData {
	//var balanceData []*BalanceData
	//totalExpenses := -model2.SumAmounts(b.payments.List(
	//	"PayeeId = 0",
	//))
	//userIds := model.StringToIds(project.UserIds)
	//userAmount := len(userIds)
	//proportionalExpenses := model2.EUR(int64(totalExpenses) / int64(userAmount))
	//var maxBalance float64
	//for _, userId := range userIds {
	//	userIdString := strconv.FormatInt(userId, 10)
	//	sentExpenses := model2.SumAmounts(b.payments.List(
	//		"PayerId = "+userIdString,
	//		"AND PayeeId = 0",
	//	))
	//	sentTransfer := model2.SumAmounts(b.payments.List(
	//		"PayerId = "+userIdString,
	//		"AND PayeeId != 0",
	//	))
	//	receivedTransfer := model2.SumAmounts(b.payments.List(
	//		"PayerId != "+userIdString,
	//		"AND PayeeId = "+userIdString,
	//	))
	//	result := proportionalExpenses + sentExpenses + sentTransfer - receivedTransfer
	//	if math.Abs(float64(result)) > float64(maxBalance) {
	//		maxBalance = float64(result)
	//	}
	//	balanceData = append(balanceData, &BalanceData{
	//		UserName:             b.users.Find(userId).Name,
	//		TotalExpenses:        totalExpenses,
	//		UserAmount:           userAmount,
	//		ProportionalExpenses: proportionalExpenses,
	//		SentExpenses:         sentExpenses,
	//		SentTransfer:         sentTransfer,
	//		ReceivedTransfer:     receivedTransfer,
	//		Result:               result,
	//	})
	//}
	//b.calculateWidths(balanceData, maxBalance)
	//return balanceData
	return nil
}

type BalanceData struct {
	ProjectName          string
	UserName             string
	TotalExpenses        model2.EUR
	UserAmount           int
	ProportionalExpenses model2.EUR
	SentExpenses         model2.EUR
	SentTransfer         model2.EUR
	ReceivedTransfer     model2.EUR
	Result               model2.EUR
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
