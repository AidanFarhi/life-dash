package model

type AggregatedExpense struct {
	Category string  `json:"category"`
	Amount   float64 `json:"amount"`
}
