package services

import (
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
 
func (s *ChapterService) Create(chapterToCreate *models.Chapter) (*models.Chapter, error) {
	newChapter := models.Chapter{
		CourseID:    chapterToCreate.CourseID,
		Title:       chapterToCreate.Title,
		Description: chapterToCreate.Description,
		VideoUrl: chapterToCreate.VideoUrl, 
	}

	var course models.Course
	if err := initializers.DB.First(&course, "id = ?", chapterToCreate.CourseID).Error; err != nil {
		return nil, err
	}

	newChapter.Course = course

	if err :=initializers.DB.Create(&newChapter).Error; err != nil {
		return nil, err
	}

	return &newChapter, nil
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

func (s *ChapterService) AddUserChapter(userChapter *models.UserChapter) error {
    result := initializers.DB.Create(userChapter)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func (s *ChapterService) GetCompletedChaptersCount(userID , courseID uuid.UUID) (int64, error) {
    var user models.User
    var count int64

    // Assuming you have a database instance named "db" initialized somewhere

    // Query the user and preload the CompletedChapters
    if err := initializers.DB.Preload("CompletedChapters").Where("id = ?", userID).First(&user).Error; err != nil {
        return 0, err
    }

    // Iterate through the completed chapters and count those in the specified course
    for _, userChapter := range user.CompletedChapters {
        if userChapter.Chapter.CourseID == courseID {
            count++
        }
    }

    return count, nil
}

// GetTotalChaptersCountForCourse retrieves the total count of chapters for a user in a specific course.
func (s *ChapterService) GetTotalChaptersCountForCourse(courseID uuid.UUID) (int64, error) {
    var count int64
	
    if err := initializers.DB.Model(&models.Chapter{}).Where("course_id = ?", courseID).Count(&count).Error; err != nil {
        return 0, err
    }

    return count, nil
}