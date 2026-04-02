package dashboard

import (
	"context"
	"personal-finance-backend/internal/models"
	"personal-finance-backend/pkg/database"
)

func GetExpenseBreakdown(userID int) (map[string]float64, error) {
	query := `
	 Select category, SUM(amount)
	 From expenses
	 where user_id=$1
	 Group by category
	`
	rows, err := database.DB.Query(context.Background(), query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[string]float64)

	for rows.Next() {
		var category string
		var amount float64

		err := rows.Scan(&category, &amount)
		if err != nil {
			return nil, err
		}
		result[category] = amount
	}
	return result, nil
}

func GetFinancialMetrics(userID int) (map[string]float64, error) {
	metrics := map[string]float64{
		"total_income": 0,
		"total_expenses": 0,
		"total_investments": 0,
		"monthly_burn_rate": 0,
	}

	var tIncome, tExp, tInv, burnRate float64

	// Since errors are skipped gracefully, defaults to 0
	_ = database.DB.QueryRow(context.Background(), "SELECT COALESCE(SUM(amount), 0) FROM incomes WHERE user_id=$1", userID).Scan(&tIncome)
	_ = database.DB.QueryRow(context.Background(), "SELECT COALESCE(SUM(amount), 0) FROM expenses WHERE user_id=$1", userID).Scan(&tExp)
	_ = database.DB.QueryRow(context.Background(), "SELECT COALESCE(SUM(amount), 0) FROM investments WHERE user_id=$1", userID).Scan(&tInv)
	
	// Burn rate: expenses over the last 30 days
	burnQuery := "SELECT COALESCE(SUM(amount), 0) FROM expenses WHERE user_id=$1 AND expense_date >= CURRENT_DATE - INTERVAL '30 days'"
	_ = database.DB.QueryRow(context.Background(), burnQuery, userID).Scan(&burnRate)

	metrics["total_income"] = tIncome
	metrics["total_expenses"] = tExp
	metrics["total_investments"] = tInv
	metrics["monthly_burn_rate"] = burnRate

	return metrics, nil
}

func GetActiveGoals(userID int) ([]models.Goal, error) {
	query := `SELECT id, user_id, goal_name, target_amount, COALESCE(target_date::TEXT, '') FROM goals WHERE user_id=$1`
	rows, err := database.DB.Query(context.Background(), query, userID)
	if err != nil {
		return []models.Goal{}, nil // return empty slice on error so it doesn't crash dashboard
	}
	defer rows.Close()

	goals := make([]models.Goal, 0)
	for rows.Next() {
		var g models.Goal
		if err := rows.Scan(&g.ID, &g.UserID, &g.GoalName, &g.TargetAmount, &g.TargetDate); err != nil {
			continue
		}
		goals = append(goals, g)
	}
	return goals, nil
}
