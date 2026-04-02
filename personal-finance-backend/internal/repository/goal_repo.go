package repository

import (
	"context"
	"personal-finance-backend/internal/models"
	"personal-finance-backend/pkg/database"
)

func CreateGoal(goal models.Goal) error {
	query := `INSERT INTO goals (user_id, goal_name, target_amount, target_date) VALUES ($1, $2, $3, $4)`
	_, err := database.DB.Exec(context.Background(), query, goal.UserID, goal.GoalName, goal.TargetAmount, goal.TargetDate)
	return err
}

func CreateGoalsBatch(goals []models.Goal) error {
	for _, goal := range goals {
		if err := CreateGoal(goal); err != nil {
			return err
		}
	}
	return nil
}
