package ports

import "desafio-nu/core/domain"

type OperationUseCase interface {
	CalcOperations(operations []*domain.Oper) ([]*domain.FeeResponse, error)
}
