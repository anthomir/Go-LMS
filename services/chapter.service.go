package services

import (
	"fmt"

	"github.com/anthomir/GoProject/initializers"
	"github.com/anthomir/GoProject/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ChapterService struct {
	db *gorm.DB
}

func NewChapterService(db *gorm.DB) *ChapterService {
	return &ChapterService{db: db}
}

func (s *ChapterService) FindById(id string) (*models.Chapter, error) {
	var chapter models.Chapter
	result := initializers.DB.Where("id = ?", id).First(&chapter)
	if result.Error != nil {
		return nil, result.Error
	}
	return &chapter, nil
}

func (s *ChapterService) Create(chapterToCreate *models.Chapter) error {
	result := initializers.DB.Create(chapterToCreate)
	fmt.Print("result: ", result)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *ChapterService) FindAllByCourseId(courseId uuid.UUID, skip int, take int) (*[]models.Chapter, error) {
	var chapters []models.Chapter
	initializers.DB.Where("course_id = ?", courseId).Find(&chapters).Offset(skip).Take(take)
	return &chapters, nil
}

func (s *ChapterService) DeleteAll() error {
	result := initializers.DB.Unscoped().Where("1 = 1").Delete(&models.Chapter{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
