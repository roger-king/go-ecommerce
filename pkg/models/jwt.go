package models

import (
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"time"
)

type JwtToken struct {
	Token string `json:"token"`
}

type UserClaims struct {
	jwt.StandardClaims

	Name string `json:name`
	Email string `json:email`
}

const jwtSecret = "secret"

func CreateJWTToken(user User) JwtToken {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{
		jwt.StandardClaims {
			ExpiresAt: time.Now().Unix() + 604800,
			Issuer: "rking",
		},
		user.Name,
		user.Email,
	})

	tokenString, error := token.SignedString([]byte(jwtSecret))

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
		return []byte(jwtSecret), nil
	})

	if error != nil {
		return false
	}

	if token.Valid && !isExpired(token.Claims.()) {
		// context.Set(req, "decoded", token.Claims)
		return true
	}

	return false
}

func at(t time.Time, f func()) {
	jwt.TimeFunc = func() time.Time {
		return t
	}
	f()
	jwt.TimeFunc = time.Now
}
