package models

type Activity struct {
	ID          int     `json:"id"`
	Type        string  `json:"type"`        // expense, income, investment, goal
	Description string  `json:"description"` // maps to description/source/asset
	Category    string  `json:"category"`    // only used for expenses, empty otherwise
	Amount      float64 `json:"amount"` 
	Date        string  `json:"date"`
}
