package models

import (
	"golang.org/x/crypto/bcrypt"
)

func Authenticate(user User) (*JwtToken, error) {
	var authedUser User

	err := db.Where("email = ?", user.Email).First(&authedUser).Error

	if err != nil {
		return nil, err
	}

	authErr := bcrypt.CompareHashAndPassword([]byte(authedUser.Password), []byte(user.Password))

	if authErr != nil {
		return nil, authErr
	}

	jwtToken := CreateJWTToken(user)

	return &jwtToken, nil
}
