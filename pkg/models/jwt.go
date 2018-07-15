package models

import (
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"time"
)

type JwtToken struct {
	Token string `json:"token"`
}

func CreateJWTToken(user User) JwtToken {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": user.Name,
		"email": user.Email,
		"password": user.Password,
		"exp": time.Now().Unix() + 604800,
	})

	tokenString, error := token.SignedString([]byte("secret"))

	if error != nil {
		fmt.Println(error)
	}

	return JwtToken{Token: tokenString}
}

func Validate(jwtToken JwtToken) bool {
	token, error := jwt.Parse(jwtToken.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte("secret"), nil
	})

	if error != nil {
		return false
	}

	if token.Valid {
		// context.Set(req, "decoded", token.Claims)
		return true
	}

	return false
}
