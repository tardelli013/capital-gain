package calc

import (
	"desafio-nu/core/domain"
)

type Sell struct {
}

type OperationSell interface {
	CalcSell(operation domain.Oper, averagePrice float64, totalLoss float64) (float64, float64)
}

func NewOperationSell() OperationSell {
	return &Sell{}
}

func (s Sell) CalcSell(operation domain.Oper, averagePrice float64, totalLoss float64) (float64, float64) {
	loss, taxToPay := calculateTaxInOperation(operation, averagePrice, totalLoss)

	return taxToPay, loss
}

func calculateTaxInOperation(operation domain.Oper, averagePrice float64, totalLoss float64) (float64, float64) {
	const maxTaxFreeProfit = 20000.00
	var loss, taxToPay float64
	var taxPercentPaid float64 = 20

	isTaxFreeOperation := float64(operation.Quantity)*operation.UnitCost <= maxTaxFreeProfit

	if operation.UnitCost > averagePrice {
		valorLucroNaOperacao := (operation.UnitCost - averagePrice) * float64(operation.Quantity)

		if valorLucroNaOperacao > totalLoss {
			lucroMenosPreju := valorLucroNaOperacao - totalLoss
			taxToPay = (lucroMenosPreju * taxPercentPaid) / 100
			loss = 0
		} else {
			loss = totalLoss - valorLucroNaOperacao
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
