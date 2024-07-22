package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Password string `json:"password" validate:"required,min=8"`
	Email    string `json:"email" validate:"email,required" gorm:"unique"`
	Token    string `json:"token"`
}

type CreateUserInput struct {
	Password string `json:"password" validate:"required,min=8"`
	Email    string `json:"email" validate:"email,required" gorm:"unique"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
