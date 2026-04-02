package models

type Investment struct {
	ID             int     `json:"id"`
	UserID         int     `json:"user_id"`
	AssetType      string  `json:"asset_type"`
	Amount         float64 `json:"amount"`
	InvestmentDate string  `json:"investment_date"` // format YYYY-MM-DD
}
