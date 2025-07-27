package models

type Payment struct {
	Amount         float64 `json:"amount"`
	Correlation_id string  `json:"id"`
}
