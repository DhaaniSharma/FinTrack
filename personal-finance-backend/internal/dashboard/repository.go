package dashboard

import (
	"context"
	"personal-finance-backend/pkg/database"
)

func GetExpenseBreakdown(userID int) (map[string]float64, error) {
	query := `
	 Select category, SUM(amount)
	 From expense
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
