package dto

import "github.com/anthomir/GoProject/models"

type GroupDetails struct {
	Group models.Group          `json:"group"`
	Users []UserWithScore 		`json:"users"`
}

// UserWithScore represents a user along with their calculated score.
type UserWithScore struct {
	User  models.User    `json:"user"`
	Score float64 		 `json:"score"`
}