package auth

import (
	// "context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"personal-finance-backend/internal/models"

	"personal-finance-backend/internal/repository"
)

func RegisterUser(user *models.User, password string) (int, error) {
	// 1. Hash the password :
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return 0, errors.New("failed to hash password")
	}

	user.PasswordHashed = string(hashedPassword)

	// 2. Store User :
	userID, err := repository.CreateUser(user)

	if err != nil {
		return 0, err
	}
	return userID, nil

}
