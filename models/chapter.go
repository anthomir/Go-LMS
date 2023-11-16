package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Chapter struct {
	UUIDModel
	CourseID    string `gorm:"column:course_id" json:"courseId"`
	Title       string `gorm:"column:title" json:"title"`
	Description string `gorm:"column:description" json:"description"`
	VideoUrl    string `gorm:"column:video_url" json:"videoUrl"`
	Course      Course `gorm:"column:course" json:"course"`
}

func (Chapter) TableName() string {
	return "chapter"
}

// ========================================================= //
// ======================= Validation ====================== //
// ========================================================= //

func (u Chapter) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

// ========================================================= //
// ================== CreatedAt UpdatedAt ================== //
// ========================================================= //

func (m *Chapter) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()
	tx.Statement.SetColumn("created_at", now)
	tx.Statement.SetColumn("updated_at", now)
	return nil
}

func (m *Chapter) BeforeUpdate(tx *gorm.DB) error {
	tx.Statement.SetColumn("updated_at", time.Now())
	return nil
}
