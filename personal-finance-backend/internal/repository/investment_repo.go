package repository

import (
	"context"
	"personal-finance-backend/internal/models"
	"personal-finance-backend/pkg/database"
)

func CreateInvestment(investment models.Investment) error {
	query := `INSERT INTO investments (user_id, asset_type, amount, investment_date) VALUES ($1, $2, $3, $4)`
	_, err := database.DB.Exec(context.Background(), query, investment.UserID, investment.AssetType, investment.Amount, investment.InvestmentDate)
	return err
}

func CreateInvestmentsBatch(investments []models.Investment) error {
	for _, inv := range investments {
		if err := CreateInvestment(inv); err != nil {
			return err
		}
	}
	return nil
}
