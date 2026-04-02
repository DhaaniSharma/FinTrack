package repository

import (
	"context"
	"personal-finance-backend/internal/models"
	"personal-finance-backend/pkg/database"
)

func CreateExpense(expense models.Expense) error {
	query := `INSERT INTO expenses (user_id, category, amount, expense_date) VALUES ($1, $2, $3, $4)`
	_, err := database.DB.Exec(context.Background(), query, expense.UserID, expense.Category, expense.Amount, expense.ExpenseDate)
	return err
}

func CreateExpensesBatch(expenses []models.Expense) error {
	// Simple loop implementation that executes in a single connection context,
	// can be optimized with pgx batch in the future
	for _, exp := range expenses {
		if err := CreateExpense(exp); err != nil {
			return err
		}
	}
	return nil
}
