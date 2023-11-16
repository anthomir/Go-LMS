package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UUIDModel struct {
	ID        uuid.UUID      `gorm:"type:uuid; primary_key; default:gen_random_uuid()" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at; index" json:"deletedAt"`
}
