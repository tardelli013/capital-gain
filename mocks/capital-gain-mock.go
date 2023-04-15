package mocks

import (
	"desafio-nu/core/domain"
	"desafio-nu/core/ports"
	"errors"
)

type OperationUseCaseMock struct {
}

func (o OperationUseCaseMock) CalcCapitalGain(_ []*domain.Oper) (interface{}, error) {
	return "", errors.New("error")
}

func NewOperationUseCaseMock() ports.OperationUseCase {
	return &OperationUseCaseMock{}
}
