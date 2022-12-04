package domain

import "github.com/shopspring/decimal"

type FeeResponse struct {
	Tax decimal.Decimal `json:"tax"`
}

func NewResponse(tax decimal.Decimal) FeeResponse {
	return FeeResponse{
		Tax: tax,
	}
}
