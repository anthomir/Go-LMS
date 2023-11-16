package dto

import (
	"github.com/anthomir/GoProject/enums"
	"github.com/go-playground/validator/v10"
)

// UserSignupDTO represents the data structure for user signup requests.
type UserSignupDTO struct {
	FirstName string     `gorm:"column:first_name" validate:"required"`
	LastName  string     `gorm:"column:last_name" validate:"required"`
	Email     string     `gorm:"column:email" validate:"required"`
	Role      enums.Role `json:"role"`
	Password  string     `gorm:"column:password; size:100" validate:"required"`
}

// Validate performs data validation on the UserSignupDTO.
func (u UserSignupDTO) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

func (UserSignupDTO) TableName() string {
	return "users"
}
