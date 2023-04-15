package calc

import (
	"desafio-nu/core/domain"
	"errors"
)

type Sell struct {
}

type OperationSell interface {
	CalcSell(operation domain.Oper, averagePrice float64, totalLoss float64, currentTotalStocks int64) (float64, float64, error)
}

func NewOperationSell() OperationSell {
	return &Sell{}
}

func (s Sell) CalcSell(operation domain.Oper, averagePrice float64, totalLoss float64, currentTotalStocks int64) (float64, float64, error) {
	if operation.Quantity > currentTotalStocks {
		return 0, 0, errors.New("Can't sell more stocks than you have")
	}

	loss, taxToPay := calculateTaxInOperation(operation, averagePrice, totalLoss)
	return taxToPay, loss, nil
}

func calculateTaxInOperation(operation domain.Oper, averagePrice float64, totalLoss float64) (float64, float64) {
	const maxProfitWithoutTaxes = 20000.00
	var loss, taxToPay float64
	var taxPercent float64 = 20

	isTaxFreeOperation := float64(operation.Quantity)*operation.UnitCost <= maxProfitWithoutTaxes

	if operation.UnitCost > averagePrice {
		profit := (operation.UnitCost - averagePrice) * float64(operation.Quantity)

		if profit > totalLoss {
			netProfit := profit - totalLoss
			taxToPay = (netProfit * taxPercent) / 100
			loss = 0
		} else {
			loss = totalLoss - profit
		}

	} else if operation.UnitCost == averagePrice {
		taxToPay = 0
		loss = 0
	} else if operation.UnitCost < averagePrice {
		taxToPay = 0
		loss = (averagePrice - operation.UnitCost) * float64(operation.Quantity)
	}

	if isTaxFreeOperation {
		return loss, 0
	}

	return loss, taxToPay
}
