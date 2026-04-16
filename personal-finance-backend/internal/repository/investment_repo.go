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

func GetInvestments(userID int) ([]models.Investment, error) {
	query := `SELECT id, user_id, asset_type, amount, COALESCE(investment_date::TEXT, '') FROM investments WHERE user_id=$1 ORDER BY investment_date DESC`
	rows, err := database.DB.Query(context.Background(), query, userID)
	if err != nil {
		return []models.Investment{}, err
	}
	defer rows.Close()

	var invs []models.Investment
	for rows.Next() {
		var i models.Investment
		if err := rows.Scan(&i.ID, &i.UserID, &i.AssetType, &i.Amount, &i.InvestmentDate); err == nil {
			invs = append(invs, i)
		}
	}
	return invs, nil
}

