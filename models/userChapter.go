package models

import (
	"time"

	"github.com/google/uuid"
)

type UserChapter struct {
    ID          uuid.UUID      `gorm:"primary_key" json:"id"`
    UserID      uuid.UUID      `gorm:"column:user_id" json:"userId"`
    ChapterID   uuid.UUID      `gorm:"column:chapter_id" json:"chapterId"`
    Completed   bool           `gorm:"column:completed" json:"completed"`
    CompletedAt time.Time      `gorm:"column:completed_at" json:"completedAt"`
    User        User           `gorm:"foreignkey:UserID" json:"user"`
    Chapter     Chapter        `gorm:"foreignkey:ChapterID" json:"chapter"`
}

func (UserChapter) TableName() string {
	return "user_chapter"
}
