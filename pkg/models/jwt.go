package models

import (
	"github.com/dgrijalva/jwt-go"
	"fmt"
)

type JwtToken struct {
	Token string `json:"token"`
}

func CreateJWTToken(user User) JwtToken {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": user.Name,
		"email": user.Email,
		"password": user.Password,
	})

	tokenString, error := token.SignedString([]byte("secret"))

	if error != nil {
		fmt.Println(error)
	}

	return JwtToken{Token: tokenString}
}
