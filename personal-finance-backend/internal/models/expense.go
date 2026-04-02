package models

import "time"

type Expense struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Category    string    `json:"category"`
	Amount      float64   `json:"amount"`
	ExpenseDate string    `json:"expense_date"` // format YYYY-MM-DD
	CreatedAt   time.Time `json:"created_at"`
}
