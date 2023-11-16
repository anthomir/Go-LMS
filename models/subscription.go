package models

import "github.com/google/uuid"

type Subscription struct {
    UserID   uuid.UUID `gorm:"primaryKey;uniqueIndex:idx_user_course" json:"userId"`
    CourseID uuid.UUID `gorm:"primaryKey;uniqueIndex:idx_user_course" json:"courseId"`
}

func (Subscription) TableName() string {
	return "subscription"
}
