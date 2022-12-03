package domain

type OperationType string

const (
	Buy  OperationType = "buy"
	Sell OperationType = "sell"
)

type Oper struct {
	Operation OperationType `json:"operation"`
	UnitCost  float64       `json:"unit-cost"`
	Quantity  int           `json:"quantity"`
}

func NewOperation(operation OperationType, unitCost float64, quantity int) *Oper {
	return &Oper{
		Operation: operation,
		UnitCost:  unitCost,
		Quantity:  quantity,
	}
}
