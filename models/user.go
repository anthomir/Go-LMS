package models

import (
	"errors"
	"time"

	"github.com/anthomir/GoProject/enums"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	UUIDModel
	FirstName           string       `gorm:"column:first_name" json:"firstName"`
	LastName            string       `gorm:"column:last_name" json:"lastName"`
	Email               string       `gorm:"column:email" validate:"required" json:"email"`
	Password            string       `gorm:"column:password; size:100" validate:"required" json:"password"`
	Role                enums.Role   `gorm:"column:role" json:"role"`
	Balance             float64      `gorm:"column:balance;default:100" json:"balance"`
	CourseSubscriptions []Course     `gorm:"many2many:subscription;" json:"subscriptions"`
	CreatedCourses      []Course     `gorm:"foreignKey:CreatedBy" json:"createdCourses"`
	Groups				[]Group 	 `gorm:"many2many:user_group;" json:"groups"`
	CompletedChapters 	[]UserChapter `gorm:"foreignKey:UserID" json:"completedChapters"`
}

func (User) TableName() string {
    return "users"
}
// ========================================================= //
// ======================= Validation ====================== //
// ========================================================= //

func (u User) Validate() error {
	validate := validator.New()

	if !u.Role.IsValid() {
		return errors.New("invalid role")
	}
	return validate.Struct(u)
}

// ========================================================= //
// ================== CreatedAt UpdatedAt ================== //
// ========================================================= //

func (m *User) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()
	tx.Statement.SetColumn("created_at", now)
	tx.Statement.SetColumn("updated_at", now)
	return nil
}

func (m *User) BeforeUpdate(tx *gorm.DB) error {
	tx.Statement.SetColumn("updated_at", time.Now())
	return nil
}
