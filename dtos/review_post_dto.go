package dto

import "github.com/go-playground/validator/v10"

type ReviewCreateDto struct {
    Rating      float64 `json:"rating" binding:"required,min=1,max=5"`
    Description string  `json:"description" binding:"required"`
}

func (u ReviewCreateDto) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
