package usecases

import (
	"desafio-nu/core/domain"
	"desafio-nu/core/ports"
	"github.com/shopspring/decimal"
)

type OperationUseCase struct {
}

func NewOperationUseCase() ports.OperationUseCase {
	return &OperationUseCase{}
}

func (o OperationUseCase) CalcCapitalGain(operations []*domain.Oper) ([]domain.FeeResponse, error) {
	var weightedAveragePrice, lastUnitCost, lastDamage, taxPaid decimal.Decimal
	var currentTotalActions int64
	listFees := make([]domain.FeeResponse, len(operations))

	for i, op := range operations {
		switch op.Operation {
		case domain.Buy:
			taxPaid, weightedAveragePrice = calcBuy(*op, currentTotalActions, weightedAveragePrice)
			currentTotalActions += op.Quantity
		case domain.Sell:
			taxPaid, lastDamage = calcSell(*op, currentTotalActions, weightedAveragePrice, lastUnitCost, lastDamage)
			currentTotalActions -= op.Quantity
		}

		lastUnitCost = op.UnitCost
		listFees[i] = domain.NewResponse(taxPaid)
	}

	return listFees, nil
}

func calcSell(operation domain.Oper, currentTotalActions int64,
	weightedAveragePrice, lastUnitCost, lastDamage decimal.Decimal) (decimal.Decimal, decimal.Decimal) {

	loss, taxToPay := calculateTaxOperation(operation, currentTotalActions, weightedAveragePrice, lastUnitCost, lastDamage)

	return taxToPay, loss
}

func calcBuy(operation domain.Oper, currentTotalActions int64, currentWeightedAverage decimal.Decimal) (decimal.Decimal, decimal.Decimal) {
	const taxOperationBuy = 0

	newWeightedAverage := calculateWeightedAverage(operation, currentTotalActions, currentWeightedAverage)
	return decimal.NewFromInt(taxOperationBuy), newWeightedAverage
}

func calculateWeightedAverage(operation domain.Oper, currentTotalActions int64, currentWeightedAverage decimal.Decimal) decimal.Decimal {
	newWeightedAverage := ((currentTotalActions * currentWeightedAverage.IntPart()) +
		(operation.Quantity * operation.UnitCost.IntPart())) / (currentTotalActions + operation.Quantity)

	return decimal.NewFromInt(newWeightedAverage)
}

func calculateTaxOperation(operation domain.Oper, currentTotalActions int64, weightedAveragePrice, lastUnitCost, lastDamage decimal.Decimal) (decimal.Decimal, decimal.Decimal) {
	const maxTaxFreeProfit = 20000.00
	var profits, loss, taxToPay decimal.Decimal
	var taxPercentPaid int64 = 20

	actionsDecreased := currentTotalActions - operation.Quantity
	if actionsDecreased == 0 {
		actionsDecreased = currentTotalActions
	}

	if transactionWithProfits := decimal.NewFromInt(operation.UnitCost.IntPart()).GreaterThan(lastUnitCost); transactionWithProfits == true {
		profits = decimal.NewFromInt(actionsDecreased*(operation.UnitCost.IntPart()-lastUnitCost.IntPart()) - lastDamage.IntPart())
		loss = decimal.NewFromInt(0)
		taxToPay = decimal.NewFromInt(profits.IntPart() * taxPercentPaid).Div(decimal.NewFromInt(100))
	} else if decimal.NewFromInt(operation.UnitCost.IntPart()).Equals(weightedAveragePrice) {
		loss = decimal.NewFromInt(0)
		profits = decimal.NewFromInt(0)
	} else {
		loss = decimal.NewFromInt(actionsDecreased * operation.UnitCost.IntPart())
		profits = decimal.NewFromInt(0)
	}

	if isTaxFreeOperation := operation.Quantity * operation.UnitCost.IntPart(); isTaxFreeOperation <= maxTaxFreeProfit {
		return loss, decimal.NewFromInt(0)
	}

	return loss, taxToPay
}
