package models

import "github.com/google/uuid"

type Review struct {
	UUIDModel
	CourseID    uuid.UUID 	`gorm:"column:course_id" json:"courseId"` // Updated foreign key
	Course      Course  	`gorm:"foreignKey:CourseID" json:"course"`
	Rating      float64 	`gorm:"type:double precision" validate:"min=1,max=5" json:"rating"`
	Description string  	`gorm:"type:text" json:"description"`
	CreatedBy   uuid.UUID    	 `gorm:"column:created_by_id" json:"createdBy"`
}

func (Review) TableName() string {
	return "review"
}