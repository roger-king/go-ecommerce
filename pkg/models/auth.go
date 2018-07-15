package models

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"fmt"
)

type AuthUser struct {
	Email string
	Password string
	Token string
}

func Authenticate (authUser AuthUser) (*AuthUser, error) {
	tokenString := authUser.Token

	if len(tokenString) > 0 {
		token, err := jwt.Parse(authUser.Token,  func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return hmacSampleSecret, nil
		})
	}

	var user User

	err := db.Where("email = ?", authUser.Email).First(&user).Error

	if err != nil {
		return nil, err
	}

	authErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authUser.Password))

	if authErr != nil {
		return nil, authErr
	}


}
