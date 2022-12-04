package calc

import (
	"desafio-nu/core/domain"
	"github.com/shopspring/decimal"
)

type Buy struct {
}

type OperationBuy interface {
	CalcBuy(operation domain.Oper, currentTotalStocks int64, currentAveragePrice decimal.Decimal) (decimal.Decimal, decimal.Decimal)
}

func NewOperationBuy() OperationBuy {
	return &Buy{}
}

func (b Buy) CalcBuy(operation domain.Oper, currentTotalStocks int64, currentAveragePrice decimal.Decimal) (decimal.Decimal, decimal.Decimal) {
	const taxOperationBuy = 0

	newWeightedAverage := calculateWeightedAverage(operation, currentTotalStocks, currentAveragePrice)
	return decimal.NewFromInt(taxOperationBuy), newWeightedAverage
}

func calculateWeightedAverage(operation domain.Oper, currentTotalStocks int64, currentWeightedAverage decimal.Decimal) decimal.Decimal {
	newWeightedAverage := ((currentTotalStocks * currentWeightedAverage.IntPart()) +
		(operation.Quantity * operation.UnitCost.IntPart())) / (currentTotalStocks + operation.Quantity)

	return decimal.NewFromInt(newWeightedAverage)
}
