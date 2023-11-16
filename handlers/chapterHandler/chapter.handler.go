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

	errCreation := chapterService.Create(&chapter)
	if errCreation != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errCreation.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Created successfully"})
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