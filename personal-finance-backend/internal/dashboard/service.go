package dashboard

import (
	"personal-finance-backend/internal/mlclient"
	"personal-finance-backend/internal/models"
)

type DashboardResponse struct {
	TotalNetWorth    float64               `json:"total_net_worth"`
	TotalIncome      float64               `json:"total_income"`
	TotalExpenses    float64               `json:"total_expenses"`
	TotalInvestments float64               `json:"total_investments"`
	MonthlyBurnRate  float64               `json:"monthly_burn_rate"`
	ExpenseBreakdown map[string]float64    `json:"expense_breakdown"`
	MLSpenderType    string                `json:"ml_spender_type"`
	ActiveGoals      []models.Goal         `json:"active_goals"`
}

func GetDashboardData(userID int) (*DashboardResponse, error) {
	breakdown, err := GetExpenseBreakdown(userID)
	if err != nil {
		return nil, err
	}

	metrics, err := GetFinancialMetrics(userID)
	if err != nil {
		return nil, err // Should never fail entirely based on the logic, but safety catch
	}

	goals, _ := GetActiveGoals(userID)

	input := mlclient.ExpenseInput{
		FoodAndDrink:     breakdown["food_and_drink"],
		Rent:             breakdown["rent"],
		Utilities:        breakdown["utilities"],
		Entertainment:    breakdown["entertainment"],
		Travel:           breakdown["travel"],
		HealthAndFitness: breakdown["health_and_fitness"],
		Shopping:         breakdown["shopping"],
		Other:            breakdown["other"],
	}

	// 0.0 values naturally feed successfully if categories are empty.
	spenderType := "Unknown"
	mlResult, mlErr := mlclient.PredictSpender(input)
	if mlErr == nil && mlResult != nil {
		spenderType = mlResult.SpenderType
	}

	netWorth := metrics["total_income"] + metrics["total_investments"] - metrics["total_expenses"]

	return &DashboardResponse{
		TotalNetWorth:    netWorth,
		TotalIncome:      metrics["total_income"],
		TotalExpenses:    metrics["total_expenses"],
		TotalInvestments: metrics["total_investments"],
		MonthlyBurnRate:  metrics["monthly_burn_rate"],
		ExpenseBreakdown: breakdown,
		MLSpenderType:    spenderType,
		ActiveGoals:      goals,
	}, nil
}
