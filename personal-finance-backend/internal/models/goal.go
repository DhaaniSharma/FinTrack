package models

type Goal struct {
	ID           int     `json:"id"`
	UserID       int     `json:"user_id"`
	GoalName     string  `json:"goal_name"`
	TargetAmount float64 `json:"target_amount"`
	TargetDate   string  `json:"target_date"` // format YYYY-MM-DD
}
