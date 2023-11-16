package models

import "github.com/google/uuid"

type UserGroup struct {
	UserId  uuid.UUID `gorm:"column:user_id;primaryKey" json:"userId"`
	GroupId uuid.UUID `gorm:"column:group_id;primaryKey" json:"groupId"`
}

func (UserGroup) TableName() string {
	return "user_group"
}
