package models

import "time"

type User struct {
	UserID         int       `json:"user_id"`
	UserName       string    `json:"user_name"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	Address        string    `json:"address"`
	PasswordHashed string    `json:"password_hashed"`
	Currency       string    `json:"currency"`
	GoogleID       *string   `json:"google_id"`
	AuthProvider   string    `json:"auth_provider"`
	CreatedAt      time.Time `json:"created_at"`
}
