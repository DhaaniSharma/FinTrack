package models

type OnboardingPayload struct {
	Age         int          `json:"age"`
	Expenses    []Expense    `json:"expenses"`
	Incomes     []Income     `json:"incomes"`
	Investments []Investment `json:"investments"`
	Goals       []Goal       `json:"goals"`
}
