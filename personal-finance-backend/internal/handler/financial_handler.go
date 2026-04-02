package handler

import (
	"encoding/json"
	"net/http"
	"personal-finance-backend/internal/middleware"
	"personal-finance-backend/internal/models"
	"personal-finance-backend/internal/repository"
)

func OnboardingBatchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userIDValue := r.Context().Value(middleware.UserIDKey)
	if userIDValue == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userID := userIDValue.(int)

	var payload models.OnboardingPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if payload.Age > 0 {
		_ = repository.UpdateUserAge(userID, payload.Age)
	}

	for i := range payload.Expenses {
		payload.Expenses[i].UserID = userID
	}
	_ = repository.CreateExpensesBatch(payload.Expenses)

	for i := range payload.Incomes {
		payload.Incomes[i].UserID = userID
	}
	_ = repository.CreateIncomesBatch(payload.Incomes)

	for i := range payload.Investments {
		payload.Investments[i].UserID = userID
	}
	_ = repository.CreateInvestmentsBatch(payload.Investments)

	for i := range payload.Goals {
		payload.Goals[i].UserID = userID
	}
	_ = repository.CreateGoalsBatch(payload.Goals)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message":"Onboarding data saved successfully"}`))
}

func CreateExpenseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := r.Context().Value(middleware.UserIDKey).(int)

	var expense models.Expense
	if err := json.NewDecoder(r.Body).Decode(&expense); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	expense.UserID = userID

	if err := repository.CreateExpense(expense); err != nil {
		http.Error(w, "Failed to create expense", http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message":"Expense created successfully"}`))
}

func CreateIncomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := r.Context().Value(middleware.UserIDKey).(int)

	var income models.Income
	if err := json.NewDecoder(r.Body).Decode(&income); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	income.UserID = userID

	if err := repository.CreateIncome(income); err != nil {
		http.Error(w, "Failed to create income", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message":"Income created successfully"}`))
}

func CreateInvestmentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := r.Context().Value(middleware.UserIDKey).(int)

	var investment models.Investment
	if err := json.NewDecoder(r.Body).Decode(&investment); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	investment.UserID = userID

	if err := repository.CreateInvestment(investment); err != nil {
		http.Error(w, "Failed to create investment", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message":"Investment created successfully"}`))
}

func CreateGoalHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := r.Context().Value(middleware.UserIDKey).(int)

	var goal models.Goal
	if err := json.NewDecoder(r.Body).Decode(&goal); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	goal.UserID = userID

	if err := repository.CreateGoal(goal); err != nil {
		http.Error(w, "Failed to create goal", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message":"Goal created successfully"}`))
}
