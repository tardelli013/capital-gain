package adapter

import "desafio-nu/core/domain"

type Scan interface {
	ScanOperations() ([]*domain.Oper, error)
}
