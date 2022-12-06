package calc

import (
	"desafio-nu/core/domain"
)

type Buy struct {
}

type OperationBuy interface {
	CalcBuy(operation domain.Oper, currentTotalStocks int64, currentAveragePrice float64) (float64, float64)
}

func NewOperationBuy() OperationBuy {
	return &Buy{}
}

func (b Buy) CalcBuy(operation domain.Oper, currentTotalStocks int64, currentAveragePrice float64) (float64, float64) {
	const taxOperationBuy = 0.00

	newWeightedAverage := calculateWeightedAverage(operation, currentTotalStocks, currentAveragePrice)
	return taxOperationBuy, newWeightedAverage
}

func calculateWeightedAverage(operation domain.Oper, currentTotalStocks int64, currentWeightedAverage float64) float64 {
	if currentWeightedAverage == 0 {
		return operation.UnitCost
	}
	op1 := float64(currentTotalStocks) * currentWeightedAverage
	op2 := float64(operation.Quantity) * operation.UnitCost
	op3 := float64(currentTotalStocks + operation.Quantity)

	newWeightedAverage := (op1 + op2) / op3

	return newWeightedAverage
}
