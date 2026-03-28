package dashboard

import (
	"encoding/json"
	"net/http"
	"personal-finance-backend/internal/middleware"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	// Extract userID from context (JWT Middleware) :
	userIDValue := r.Context().Value(middleware.UserIDKey)
	if userIDValue == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userID, ok := userIDValue.(int)
	if !ok {
		http.Error(w, "Invalid user ID", http.StatusInternalServerError)
		return
	}

	// Call service layer :
	data, err := GetDashboardData(userID)
	if err != nil {
		http.Error(w, "Failed to fetch dashboard", http.StatusInternalServerError)
		return
	}

	// Send Response :
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
