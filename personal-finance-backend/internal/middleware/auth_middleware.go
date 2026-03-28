package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserIDKey contextKey = "user_id"

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. Get Authorization header :
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "Missing Token", http.StatusUnauthorized)
			return
		}

		// 2. Format : Bearer <token>
		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}
		tokenString := parts[1]

		// 3; Parse Token :
		secret := os.Getenv("JWT_SECRET")
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			http.Error(w, "Invalid signing method", http.StatusUnauthorized)
			return
		}

		// Extract user_id from token :
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		userIDFloat := claims["user_id"].(float64)
		userID := int(userIDFloat)

		// Store in context :
		ctx := context.WithValue(r.Context(), UserIDKey, userID)

		next(w, r.WithContext(ctx))

	}
}
