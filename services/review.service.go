package services

import (
	"github.com/anthomir/GoProject/initializers"
	"github.com/anthomir/GoProject/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReviewService struct {
	db *gorm.DB
}

func NewReviewService(db *gorm.DB) *ReviewService {
	return &ReviewService{db: db}
}

func (s *ReviewService) AddReview(course *models.Course, createdBy *models.User, rating float64, description string) error {
    review := models.Review{
        Course:      *course,
        Rating:      rating,
        Description: description,
        CreatedBy:   createdBy.ID,
    }

    result := initializers.DB.Create(&review)
    if result.Error != nil {
        return result.Error
    }

    return nil
}

func (s *ReviewService) RemoveReview(course *models.Course, createdBy *models.User) error {
    result := initializers.DB.Where("created_by_id = ? AND course_id = ?", createdBy.ID, course.ID).Delete(&models.Review{})
    if result.Error != nil {
        return result.Error
    }
    return nil
}
func (s *ReviewService) CheckReview(courseId, userId uuid.UUID) (*models.Review, error) {
	var reviewItem models.Review
	result := initializers.DB.Where("created_by_id = ? AND course_id = ?", userId, courseId).First(&reviewItem)
	if result.Error != nil {
		return nil, result.Error
	}

	return &reviewItem, nil
}
