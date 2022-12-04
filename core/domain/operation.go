package domain

import "github.com/shopspring/decimal"

type OperationType string

const (
	Buy  OperationType = "buy"
	Sell OperationType = "sell"
)

type Oper struct {
	Operation OperationType   `json:"operation"`
	UnitCost  decimal.Decimal `json:"unit-cost"`
	Quantity  int64           `json:"quantity"`
}
