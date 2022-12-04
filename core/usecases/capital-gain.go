package usecases

import (
	"desafio-nu/core/domain"
	"desafio-nu/core/ports"
	"desafio-nu/core/usecases/calc"
	"github.com/shopspring/decimal"
)

type OperationUseCase struct {
}

func NewOperationUseCase() ports.OperationUseCase {
	return &OperationUseCase{}
}

func (o OperationUseCase) CalcCapitalGain(operations []*domain.Oper) ([]domain.FeeResponse, error) {
	opBuy := calc.NewOperationBuy()
	opSell := calc.NewOperationSell()
	var averagePrice, lastUnitCost, lastDamage, taxPaid decimal.Decimal
	var currentTotalStocks int64
	listFees := make([]domain.FeeResponse, len(operations))

	for i, op := range operations {
		switch op.Operation {
		case domain.Buy:
			taxPaid, averagePrice = opBuy.CalcBuy(*op, currentTotalStocks, averagePrice)
			currentTotalStocks += op.Quantity
		case domain.Sell:
			taxPaid, lastDamage = opSell.CalcSell(*op, currentTotalStocks, averagePrice, lastUnitCost, lastDamage)
			currentTotalStocks -= op.Quantity
		}

		lastUnitCost = op.UnitCost
		listFees[i] = domain.NewResponse(taxPaid)
	}

	return listFees, nil
}
