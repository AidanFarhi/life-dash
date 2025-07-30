package model

type Expense struct {
	Date     string
	Category string
	Amount   float64
}

type AggregatedExpense struct {
	Category string  `json:"category"`
	Amount   float64 `json:"amount"`
}
