package models

import "github.com/google/uuid"

type UserGroup struct {
    ID      uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key" json:"id"`
    UserID  uuid.UUID `gorm:"column:user_id" json:"userId"`
    GroupID uuid.UUID `gorm:"column:group_id" json:"groupId"`
}

func (UserGroup) TableName() string {
    return "user_group"
}
