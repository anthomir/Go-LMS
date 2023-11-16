package models

import "github.com/google/uuid"

type Subscription struct {
	UserID   uuid.UUID `gorm:"primaryKey" json:"userId"`
	CourseID uuid.UUID `gorm:"primaryKey" json:"courseId"`
}

func (Subscription) TableName() string {
	return "subscription"
}
