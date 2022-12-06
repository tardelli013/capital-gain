package domain

import (
	"encoding/json"
)

type FeeResponse struct {
	Tax float64 `json:"tax"`
}

func NewResponse(tax float64) FeeResponse {
	return FeeResponse{
		Tax: tax,
	}
}

func StringToListFeeResponse(s string) []FeeResponse {
	var f []FeeResponse
	err := json.Unmarshal([]byte(s), &f)
	if err != nil {
		panic(err)
	}
	return f
}
