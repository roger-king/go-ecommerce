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
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(168)).Unix(),
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

//func Validate(jwtToken JwtToken) bool {
//	validationError := &jwt.ValidationError{}
//	// now := time.Now().Unix()
//
//	token, _ := jwt.Parse(jwtToken.Token, func(token *jwt.Token) (interface{}, error) {
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, fmt.Errorf("there was an error")
//		}
//		return []byte(jwtSecret), nil
//	})
//
//	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
//		log.Info(claims["ExpiresAt"])
//		//if now > int64(claims["ExpiresAt"]) {
//		//	return false
//		//}
//
//		return true
//	} else {
//		validationError.Errors |= jwt.ValidationErrorIssuer
//		return false
//	}
//}

