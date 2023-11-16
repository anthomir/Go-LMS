package dto

import (
	"github.com/go-playground/validator/v10"
)

// UserLoginDTO represents the data structure for user login requests.
type UserLoginDTO struct {
	Email    string `gorm:"column:email" validate:"required"`
	Password string `gorm:"column:password; size:100" validate:"required"`
}

// Validate performs data validation on the UserSignupDTO.
func (u UserLoginDTO) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}