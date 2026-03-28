package dashboard

import (
	"personal-finance-backend/internal/mlclient"
)

func GetDashboardData(userID int) (map[string]interface{}, error) {
	// 1. Expense Breakdown :
	breakdown, err := GetExpenseBreakdown(userID)

	if err != nil {
		return nil, err
	}

	// 2. Convert to ML input :
	input := mlclient.ExpenseInput{
		FoodAndDrink:     breakdown["food"],
		Rent:             breakdown["rent"],
		Utilities:        breakdown["utilities"],
		Entertainment:    breakdown["entertainment"],
		Travel:           breakdown["travel"],
		HealthAndFitness: breakdown["health"],
		Shopping:         breakdown["shopping"],
		Other:            breakdown["other"],
	}

	// 3. Call ML :
	spenderType, _ := mlclient.PredictSpender(input)

	return map[string]interface{}{
		"expense_breakdown": breakdown,
		"spender_type":      spenderType,
	}, nil

}
