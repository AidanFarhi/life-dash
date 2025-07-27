package model

type Expense struct {
	Date     string  `json:"date"`
	Category string  `json:"category"`
	Amount   float64 `json:"amount"`
}
