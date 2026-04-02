package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"personal-finance-backend/pkg/database"

	"personal-finance-backend/internal/auth"

	"personal-finance-backend/internal/middleware"

	"personal-finance-backend/internal/dashboard"

	"personal-finance-backend/internal/handler"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = database.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to the database")
	}

	log.Println("Server running on : 8080")

	http.HandleFunc("/auth/register", auth.RegisterHandler)
	http.HandleFunc("/auth/login", auth.LoginHandler)
	http.HandleFunc("/api/profile", middleware.AuthMiddleware(auth.ProfileHandler))
	http.HandleFunc("/api/dashboard", middleware.AuthMiddleware(dashboard.DashboardHandler))
	http.HandleFunc("/api/onboarding", middleware.AuthMiddleware(handler.OnboardingBatchHandler))
	http.HandleFunc("/api/expenses", middleware.AuthMiddleware(handler.CreateExpenseHandler))
	http.HandleFunc("/api/incomes", middleware.AuthMiddleware(handler.CreateIncomeHandler))
	http.HandleFunc("/api/investments", middleware.AuthMiddleware(handler.CreateInvestmentHandler))
	http.HandleFunc("/api/goals", middleware.AuthMiddleware(handler.CreateGoalHandler))
	http.ListenAndServe(":8080", nil)

}
