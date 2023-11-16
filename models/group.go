package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Group struct {
    UUIDModel
    CourseID    uuid.UUID    `gorm:"column:course_id" validate:"required" json:"courseId"`
    Title       string    `gorm:"column:title" validate:"required" json:"title"`
    Description string    `gorm:"column:description" json:"description"`
    Users       []User    `gorm:"many2many:user_group;" json:"users"`
    CreatedBy   uuid.UUID `gorm:"column:created_by;index" json:"createdBy"`
}

func (Group) TableName() string {
    return "group" // specify the actual table name
}

// ========================================================= //
// ======================= Validation ====================== //
// ========================================================= //

func (u Group) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

// ========================================================= //
// ================== CreatedAt UpdatedAt ================== //
// ========================================================= //

func (m *Group) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()
	tx.Statement.SetColumn("created_at", now)
	tx.Statement.SetColumn("updated_at", now)
	return nil
}

func (m *Group) BeforeUpdate(tx *gorm.DB) error {
	tx.Statement.SetColumn("updated_at", time.Now())
	return nil
}
