package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"errors"
	"github.com/labstack/gommon/log"
)

type User struct {
	gorm.Model

	Name string `json:"name"`
	Password string `json:"password"`
}

func (u *User) BeforeCreate(scope *gorm.Scope) error {
	log.Info(u.Password)
	hashedPassword, err := hashPassword(u.Password)

	if err != nil {
		return errors.New("cannot hash password")
	}

	scope.SetColumn("password", hashedPassword)

	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func CreateUser(user User) (*User, error) {
	var err error
	err = db.Create(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, err
}

