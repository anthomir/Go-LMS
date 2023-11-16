package models

import (
	"time"

	"github.com/anthomir/GoProject/enums"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Course struct {
	UUIDModel
	Title       string    `gorm:"column:title" json:"title"`
	Description string    `gorm:"column:description" json:"description"`
	Price       string    `gorm:"column:price" json:"price"`
	ImageUrl    string    `gorm:"column:image_url" json:"imageUrl"`
	Chapters    []Chapter `gorm:"foreignKey:CourseID" json:"chapters"`
	Subscribers []User    `gorm:"many2many:subscription;" json:"subscribers"`
	Reviews 	[]Review  `gorm:"foreignKey:CourseId" json:"reviews"`
	Category    enums.Category  `gorm:"column:category" json:"category"`
	CreatedBy   uuid.UUID  `gorm:"column:created_by" json:"createdBy"`
}

type CourseWithReviews struct {
	Course
	Reviews int `gorm:"column:reviews_count" json:"reviewsCount"` // Count of reviews for each post
}

func (Course) TableName() string {
	return "course"
}

// ========================================================= //
// ======================= Validation ====================== //
// ========================================================= //

func (u Course) Validate() error {
	validate := validator.New()

	return validate.Struct(u)
}

// ========================================================= //
// ================== CreatedAt UpdatedAt ================== //
// ========================================================= //

func (m *Course) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()
	tx.Statement.SetColumn("created_at", now)
	tx.Statement.SetColumn("updated_at", now)
	return nil
}

func (m *Course) BeforeUpdate(tx *gorm.DB) error {
	tx.Statement.SetColumn("updated_at", time.Now())
	return nil
}
