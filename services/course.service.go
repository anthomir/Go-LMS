package services

import (
	"errors"

	"github.com/anthomir/GoProject/initializers"
	"github.com/anthomir/GoProject/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CourseService struct {
	db *gorm.DB
}

func NewCourseService(db *gorm.DB) *CourseService {
	return &CourseService{db: db}
}

func (s *CourseService) FindById(id string) (*models.Course, error) {
	var course models.Course

	result := initializers.DB.Where("id = ?", id).First(&course)
	if result.Error != nil {
		return nil, result.Error
	}
	return &course, nil
}

func (s *CourseService) Create(courseModel *models.Course) (*models.Course, error) {
    result := initializers.DB.Create(courseModel)

    if result.Error != nil {
        return nil, result.Error
    }

    // Retrieve the created object with associated users populated
    createdCourse := &models.Course{}
    if err := initializers.DB.Preload("Subscribers").First(createdCourse, courseModel.ID).Error; err != nil {
        return nil, err
    }

    return createdCourse, nil
}

func (s *CourseService) FindAll(category, minPrice, maxPrice, title string) (*[]models.CourseWithReviews, error) {
	var courseWithReviews []models.CourseWithReviews

	query := initializers.DB.Model(&models.Course{}).
		Select("course.*, COUNT(Review.id) as Review_count").
		Joins("LEFT JOIN Review ON course.id = Review.course_id").
		Group("course.id")

	if category != "" {
		query = query.Where("category = ?", category)
	}

	if minPrice != "" {
		query = query.Where("price >= ?", minPrice)
	}

	if maxPrice != "" {
		query = query.Where("price <= ?", maxPrice)
	}

	if title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}

	result := query.Find(&courseWithReviews)

	if result.Error != nil {
		return nil, result.Error
	}

	return &courseWithReviews, nil
}

func (s *CourseService) DeleteAll() error {
	result := initializers.DB.Unscoped().Where("1 = 1").Delete(&models.Course{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *CourseService) FindByUserId(userId uuid.UUID) ([]models.CourseWithReviews, error) {
	var courseWithReviews []models.CourseWithReviews

	if err := initializers.DB.
		Model(&models.Course{}).
		Select("course.*, COUNT(reviews.id) as reviews_count").
		Joins("LEFT JOIN reviews ON course.id = reviews.course_id").
		Where("course.created_by = ?", userId).
		Group("course.id").
		Find(&courseWithReviews).
		Error; err != nil {
		// Handle the error, e.g., user not found
		return nil, err
	}

	return courseWithReviews, nil
}

func (s *CourseService) SubscribeUserToCourse(user *models.User, course *models.Course) error {
	if err := initializers.DB.Model(user).Association("CourseSubscriptions").Append(course); err != nil {
		return errors.New("failed to subscribe user to the course")
	}
	return nil
}
