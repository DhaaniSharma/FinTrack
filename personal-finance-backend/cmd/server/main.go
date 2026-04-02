package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"personal-finance-backend/pkg/database"

	"personal-finance-backend/internal/auth"

	"personal-finance-backend/internal/middleware"

	"personal-finance-backend/internal/dashboard"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

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
	mux := http.NewServeMux()

    mux.HandleFunc("/auth/register", auth.RegisterHandler)
    mux.HandleFunc("/auth/login", auth.LoginHandler)
    mux.HandleFunc("/api/profile", middleware.AuthMiddleware(auth.ProfileHandler))
    mux.HandleFunc("/api/dashboard", middleware.AuthMiddleware(dashboard.DashboardHandler))

	 mux.HandleFunc("/auth/google/login", auth.GoogleLoginHandler)
     mux.HandleFunc("/auth/google/callback", auth.GoogleCallbackHandler)

    http.ListenAndServe(":8080", enableCORS(mux))

}
