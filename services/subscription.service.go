package services

import (
	"github.com/anthomir/GoProject/initializers"
	"github.com/anthomir/GoProject/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubscriptionService struct {
	db *gorm.DB
}

func NewSubscriptionService(db *gorm.DB) *SubscriptionService {
	return &SubscriptionService{db: db}
}

func (s *SubscriptionService) IsUserSubscribedToCourse(userId, courseId uuid.UUID) (bool, error) {
	var subscription models.Subscription
	err := initializers.DB.Model(&subscription).Where("user_id = ? AND course_id = ?", userId, courseId).First(&subscription).Error

	if err == gorm.ErrRecordNotFound {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}
