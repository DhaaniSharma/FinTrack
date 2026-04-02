package dashboard

import (
	"strings"
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

	// Normalizer: Converting wild-card string names into the 8 specific ML model keys
	normalized := make(map[string]float64)
	for k, v := range breakdown {
		lower := strings.ToLower(k)
		if strings.Contains(lower, "food") || strings.Contains(lower, "grocer") || strings.Contains(lower, "drink") || strings.Contains(lower, "restaurant") || strings.Contains(lower, "coffee") || strings.Contains(lower, "cafe") || strings.Contains(lower, "snack") {
			normalized["food_and_drink"] += v
		} else if strings.Contains(lower, "rent") || strings.Contains(lower, "mortgage") || strings.Contains(lower, "housing") || strings.Contains(lower, "lease") {
			normalized["rent"] += v
		} else if strings.Contains(lower, "utilit") || strings.Contains(lower, "bill") || strings.Contains(lower, "electric") || strings.Contains(lower, "water") {
			normalized["utilities"] += v
		} else if strings.Contains(lower, "entertain") || strings.Contains(lower, "movie") || strings.Contains(lower, "game") {
			normalized["entertainment"] += v
		} else if strings.Contains(lower, "travel") || strings.Contains(lower, "transit") || strings.Contains(lower, "gas") || strings.Contains(lower, "car") || strings.Contains(lower, "transport") {
			normalized["travel"] += v
		} else if strings.Contains(lower, "health") || strings.Contains(lower, "fitness") || strings.Contains(lower, "gym") || strings.Contains(lower, "medical") {
			normalized["health_and_fitness"] += v
		} else if strings.Contains(lower, "shop") || strings.Contains(lower, "cloth") || strings.Contains(lower, "amazon") || strings.Contains(lower, "ecom") {
			normalized["shopping"] += v
		} else {
			normalized["other"] += v
		}
	}

	input := mlclient.ExpenseInput{
		FoodAndDrink:     normalized["food_and_drink"],
		Rent:             normalized["rent"],
		Utilities:        normalized["utilities"],
		Entertainment:    normalized["entertainment"],
		Travel:           normalized["travel"],
		HealthAndFitness: normalized["health_and_fitness"],
		Shopping:         normalized["shopping"],
		Other:            normalized["other"],
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
		ExpenseBreakdown: normalized,
		MLSpenderType:    spenderType,
		ActiveGoals:      goals,
	}, nil
}
