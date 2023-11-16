package dto

import (
	"github.com/go-playground/validator/v10"
)

// UserLoginDTO represents the data structure for user login requests.
type CourseCreateDto struct {
	Title       string `gorm:"column:title" validate:"required" json:"title"`
	Description string `gorm:"column:description" validate:"required" json:"description"`
	Price       string `gorm:"column:price" validate:"required" json:"price"`
}


func (u CourseCreateDto) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

