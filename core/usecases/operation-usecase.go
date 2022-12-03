package usecases

import (
	"desafio-nu/core/domain"
	"desafio-nu/core/ports"
)

type OperationUseCase struct {
}

func NewOperationUseCase() ports.OperationUseCase {
	return &OperationUseCase{}
}

func (o OperationUseCase) CalcOperations(operations []*domain.Oper) ([]*domain.FeeResponse, error) {
	/*for _, operation := range operations {

	}*/

	asd := domain.NewResponse(34232)
	asd2 := domain.NewResponse(789346577)
	list := []*domain.FeeResponse{
		asd,
		asd2,
	}
	return list, nil
}
