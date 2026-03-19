package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"personal-finance-backend/internal/repository"
)

/*  Login User :
Steps :
1. Find user in database
2. Compare password using bcrypt
3. Return user if valid
*/

func LoginUser(identifier string, password string) (string, error) {
	// 1. Find user by email or username :
	user, err := repository.GetUserByEmailOrUsername(identifier)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	// 2. Compare password :
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHashed),
		[]byte(password),
	)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	token, err := GenerateJWT(user.UserID)
	if err != nil {
		return "", err
	}
	return token, nil

}
