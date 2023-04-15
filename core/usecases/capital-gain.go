package usecases

import (
	"desafio-nu/core/domain"
	"desafio-nu/core/ports"
	"desafio-nu/core/usecases/calc"
)

type OperationUseCase struct {
}

func NewOperationUseCase() ports.OperationUseCase {
	return &OperationUseCase{}
}

func (o OperationUseCase) CalcCapitalGain(operations []*domain.Oper) (interface{}, error) {
	opBuy := calc.NewOperationBuy()
	opSell := calc.NewOperationSell()
	var taxPaid, totalLoss, averagePrice float64
	var currentTotalStocks int64
	var calcSellError error
	listFees := make([]interface{}, len(operations))

	for i, op := range operations {
		switch op.Operation {
		case domain.Buy:
			taxPaid, averagePrice = opBuy.CalcBuy(*op, currentTotalStocks, averagePrice)
			currentTotalStocks += op.Quantity
		case domain.Sell:
			taxPaid, totalLoss, calcSellError = opSell.CalcSell(*op, averagePrice, totalLoss, currentTotalStocks)
			currentTotalStocks -= op.Quantity
		}

		if calcSellError != nil {
			listFees[i] = domain.BuildResponse(calcSellError)
		} else {
			listFees[i] = domain.BuildResponse(taxPaid)
		}
	}

	return listFees, nil
}
