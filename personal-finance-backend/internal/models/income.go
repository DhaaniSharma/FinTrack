package models

type Income struct {
	ID         int     `json:"id"`
	UserID     int     `json:"user_id"`
	Source     string  `json:"source"`
	Amount     float64 `json:"amount"`
	IncomeDate string  `json:"income_date"` // format YYYY-MM-DD
}
