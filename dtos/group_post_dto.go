package dto

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// UserLoginDTO represents the data structure for user login requests.
type GroupCreateDto struct {
	CourseID    string    `gorm:"column:course_id" validate:"required" json:"courseId"`
	Title       string    `gorm:"column:title" validate:"required" json:"title"`
	Description string    `gorm:"column:description" json:"description"`
	Users       []uuid.UUID    `json:"users"`
}

type GroupInternalDto struct {
    CourseID    string    `gorm:"column:course_id" validate:"required" json:"courseId"`
    Title       string    `gorm:"column:title" validate:"required" json:"title"`
    Description string    `gorm:"column:description" json:"description"`
    Users       []uuid.UUID    `gorm:"many2many:user_group;" json:"users"`
    CreatedBy   uuid.UUID `gorm:"column:created_by" json:"createdBy"`
}

func (u GroupCreateDto) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

func (GroupInternalDto) TableName() string {
	return "group"
}
