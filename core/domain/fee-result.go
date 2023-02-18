package domain

import (
	"encoding/json"
)

type FeeResponse struct {
	Tax float64 `json:"tax"`
}

type FeeResponseError struct {
	Error string `json:"error"`
}

func NewResponse(tax float64) FeeResponse {
	return FeeResponse{
		Tax: tax,
	}
}

func NewResponseError(error string) FeeResponseError {
	return FeeResponseError{
		Error: error,
	}
}

func BuildResponse(myInterface interface{}) interface{} {
	switch v := myInterface.(type) {
	case error:
		return NewResponseError(v.Error())
	case float64:
		return NewResponse(v)
	default:
		panic("deu ruim")
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
