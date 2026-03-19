package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"personal-finance-backend/pkg/database"

	"personal-finance-backend/internal/auth"

	"personal-finance-backend/internal/middleware"
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
	http.ListenAndServe(":8080", nil)

}
