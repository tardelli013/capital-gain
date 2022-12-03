package domain

type FeeResponse struct {
	Tax float64 `json:"tax"`
}

func NewResponse(tax float64) *FeeResponse {
	return &FeeResponse{
		Tax: tax,
	}
}
