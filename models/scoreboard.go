package models

import (
	"github.com/google/uuid"
)

type Scoreboard struct {
	UUIDModel
    UserID   uuid.UUID `gorm:"column:user_id" json:"userId"`
    GroupID  uuid.UUID `gorm:"column:group_id" json:"groupId"`
    Score    int       `gorm:"column:score" json:"score"`
}