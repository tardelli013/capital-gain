package ports

import "desafio-nu/core/domain"

type OperationUseCase interface {
	CalcCapitalGain(operations []*domain.Oper) (interface{}, error)
}
