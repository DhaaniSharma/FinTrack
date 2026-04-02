package repository

import (
	"context"
	"personal-finance-backend/internal/models"
	"personal-finance-backend/pkg/database"
)

func CreateIncome(income models.Income) error {
	query := `INSERT INTO incomes (user_id, source, amount, income_date) VALUES ($1, $2, $3, $4)`
	_, err := database.DB.Exec(context.Background(), query, income.UserID, income.Source, income.Amount, income.IncomeDate)
	return err
}

func CreateIncomesBatch(incomes []models.Income) error {
	for _, inc := range incomes {
		if err := CreateIncome(inc); err != nil {
			return err
		}
	}
	return nil
}
