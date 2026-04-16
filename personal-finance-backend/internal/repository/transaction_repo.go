package repository

import (
	"context"
	"errors"
	"fmt"
	"personal-finance-backend/pkg/database"
)

func DeleteTransaction(userID int, txType string, txID int) error {
	var table string
	switch txType {
	case "expense":
		table = "expenses"
	case "income":
		table = "incomes"
	case "investment":
		table = "investments"
	default:
		return errors.New("invalid transaction type")
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 AND user_id=$2", table)
	_, err := database.DB.Exec(context.Background(), query, txID, userID)
	return err
}
