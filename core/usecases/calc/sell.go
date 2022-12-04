package calc

import (
	"desafio-nu/core/domain"
	"github.com/shopspring/decimal"
)

type Sell struct {
}

type OperationSell interface {
	CalcSell(operation domain.Oper, currentTotalStocks int64, averagePrice, lastUnitCost, lastDamage decimal.Decimal) (decimal.Decimal, decimal.Decimal)
}

func NewOperationSell() OperationSell {
	return &Sell{}
}

func (s Sell) CalcSell(operation domain.Oper, currentTotalStocks int64, averagePrice, lastUnitCost, lastDamage decimal.Decimal) (decimal.Decimal, decimal.Decimal) {
	loss, taxToPay := calculateTaxOperation(operation, currentTotalStocks, averagePrice, lastUnitCost, lastDamage)

	return taxToPay, loss
}

func calculateTaxOperation(operation domain.Oper, currentTotalStocks int64, averagePrice, lastUnitCost, lastDamage decimal.Decimal) (decimal.Decimal, decimal.Decimal) {
	const maxTaxFreeProfit = 20000.00
	var profits, loss, taxToPay decimal.Decimal
	var taxPercentPaid int64 = 20

	stocksDecreased := currentTotalStocks - operation.Quantity
	if stocksDecreased == 0 {
		stocksDecreased = currentTotalStocks
	}

	if transactionWithProfits := decimal.NewFromInt(operation.UnitCost.IntPart()).GreaterThan(lastUnitCost); transactionWithProfits == true {
		profits = decimal.NewFromInt(stocksDecreased*(operation.UnitCost.IntPart()-lastUnitCost.IntPart()) - lastDamage.IntPart())
		loss = decimal.NewFromInt(0)
		taxToPay = decimal.NewFromInt(profits.IntPart() * taxPercentPaid).Div(decimal.NewFromInt(100))
	} else if decimal.NewFromInt(operation.UnitCost.IntPart()).Equals(averagePrice) {
		loss = decimal.NewFromInt(0)
		profits = decimal.NewFromInt(0)
	} else {
		loss = decimal.NewFromInt(stocksDecreased * operation.UnitCost.IntPart())
		profits = decimal.NewFromInt(0)
	}

	if isTaxFreeOperation := operation.Quantity * operation.UnitCost.IntPart(); isTaxFreeOperation <= maxTaxFreeProfit {
		return loss, decimal.NewFromInt(0)
	}

	return loss, taxToPay
}
