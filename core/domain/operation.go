package domain

type OperationType string

const (
	Buy  OperationType = "buy"
	Sell OperationType = "sell"
)

type Oper struct {
	Operation OperationType `json:"operation"`
	UnitCost  float64       `json:"unit-cost"`
	Quantity  int64         `json:"quantity"`
}
