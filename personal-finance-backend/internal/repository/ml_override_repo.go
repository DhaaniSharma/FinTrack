package repository

import (
	"context"
	"personal-finance-backend/pkg/database"
)

func UpsertOverride(userID int, rawText, correctedCategory string) error {
	query := `
		INSERT INTO ml_category_overrides (user_id, raw_text, corrected_category)
		VALUES ($1, $2, $3)
		ON CONFLICT (user_id, raw_text) 
		DO UPDATE SET corrected_category = EXCLUDED.corrected_category;
	`
	_, err := database.DB.Exec(context.Background(), query, userID, rawText, correctedCategory)
	if err == nil {
		// Retrospectively update all existing matching expenses for this user
		updatePast := `UPDATE expenses SET category = $1 WHERE user_id = $2 AND LOWER(description) = LOWER($3)`
		database.DB.Exec(context.Background(), updatePast, correctedCategory, userID, rawText)
	}
	return err
}

func GetOverride(userID int, rawText string) (string, bool) {
	var corrected string
	query := `SELECT corrected_category FROM ml_category_overrides WHERE user_id=$1 AND raw_text=$2`
	err := database.DB.QueryRow(context.Background(), query, userID, rawText).Scan(&corrected)
	if err != nil {
		return "", false
	}
	return corrected, true
}
