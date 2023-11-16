package chapterHandler

import (
	"net/http"
	"strconv"

	"github.com/anthomir/GoProject/models"
	"github.com/anthomir/GoProject/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Find Chapter By Id
func FindById(c *gin.Context) {
	idParam := c.Param("id")
	chapterService := services.NewChapterService(&gorm.DB{})

	response, err := chapterService.FindById(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

// Find Chapters By User
func FindAllByCourseId(c *gin.Context) {
	idParam := c.Param("id")
	userID, err := uuid.Parse(idParam)
	if err!= nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
	}
	skipParam, err1 := strconv.Atoi(c.DefaultQuery("skip", "0"))
	takeParam, err2 := strconv.Atoi(c.DefaultQuery("take", "10"))

	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid skip or take parameters"})
		return
	}

	chapterService := services.NewChapterService(&gorm.DB{})

	response, err := chapterService.FindAllByCourseId(userID, skipParam, takeParam)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

// Create Chapter
func Create(c *gin.Context) {
	chapterService := services.NewChapterService(&gorm.DB{})

	var chapter models.Chapter
	if err := c.ShouldBindJSON(&chapter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := chapter.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	result, errCreation := chapterService.Create(&chapter)
	if errCreation != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errCreation.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result ,"msg": "Created successfully"})
}

// Delete All Chapters
func DeleteAll(c *gin.Context) {
	chapterService := services.NewChapterService(&gorm.DB{})
	err := chapterService.DeleteAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Deleted successfully"})
}

func AddUserChapterHandler(c *gin.Context) {
	chapterService := services.NewChapterService(&gorm.DB{})

	userAuth, _ := c.Get("user")
	userParam, ok := userAuth.(models.User)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	chapterId, err := uuid.Parse(c.Param("chapterId"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
        return
    }
    completed, _ := strconv.ParseBool(c.Param("completed"))

    userChapter := models.UserChapter{
        UserID:     userParam.ID,
        ChapterID:  chapterId,
        Completed:  completed,
    }

    errServ := chapterService.AddUserChapter(&userChapter)
    if errServ != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add UserChapter record"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "UserChapter record added successfully"})
}
