package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"errors"
	log "github.com/sirupsen/logrus"
)

type User struct {
	gorm.Model

	Name string `json:"name"`
	Email string `gorm:"type:varchar(100);unique_index"`
	Password string `json:"password"`
}

type userDTO struct {
	Name string `json:"name"`
	Email string `json:"email"`
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

func CreateUser(user User) (*userDTO, error) {
	var err error
	err = db.Create(&user).Error

	if err != nil {
		return nil, err
	}

	createdUser := userDTO{
		Name: user.Name,
		Email: user.Email,
	}

	return &createdUser, nil
}

func FindUserByEmail(email string) (*userDTO, error) {
	var user User

	err := db.Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}


	return &userDTO{
		Name: user.Name,
		Email: user.Email,
	}, nil
}

