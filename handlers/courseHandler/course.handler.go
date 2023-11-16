package courseHandler

import (
	"net/http"
	"os"
	"path/filepath"

	dto "github.com/anthomir/GoProject/dtos"
	"github.com/anthomir/GoProject/models"
	"github.com/anthomir/GoProject/services"
	"github.com/anthomir/GoProject/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func FindById(c *gin.Context) {
	idParam := c.Param("id")
	courseService := services.NewCourseService(&gorm.DB{})

	response, err := courseService.FindById(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func FindAll(c *gin.Context) {
	courseService := services.NewCourseService(&gorm.DB{})
	response, err := courseService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if len(*response) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No data found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func FindByUserId(c *gin.Context) {
	idParam := c.Param("id")
	userID, parseErr := uuid.Parse(idParam)
	if parseErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": parseErr,
		})
		return
	}
	courseService := services.NewCourseService(&gorm.DB{})

	response, err := courseService.FindByUserId(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func UploadHandler(c *gin.Context) {
	formKey := "file"

	file, err := c.FormFile(formKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File not provided"})
		return
	}

	if !utils.IsValidImage(c, formKey) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file format. Only PNG or JPEG allowed."})
		return
	}

	uploadPath := "./public"
	if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
		if os.IsPermission(err) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
		}
		return
	}

	newFileNameWithExt := uuid.New().String() + filepath.Ext(file.Filename)
	filePath := filepath.Join(uploadPath, newFileNameWithExt)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save the uploaded file"})
		return
	}

	// Include file path in the success response
	c.JSON(http.StatusOK, gin.H{
		"message":      "File uploaded successfully",
		"filePath":     os.Getenv("PROEJCT_URL")+ filePath,
	})
}

func Create(c *gin.Context) {
    courseService := services.NewCourseService(&gorm.DB{})
    user := c.MustGet("user").(*models.User)

    // Parse JSON request body into CourseCreateDto
    var courseDto dto.CourseCreateDto
    if err := c.ShouldBindJSON(&courseDto); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := courseDto.Validate(); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    courseModel := models.Course{
        Title:       courseDto.Title,
        Description: courseDto.Description,
        Price:       courseDto.Price,
        CreatedBy:   user.ID,
    }

    createdCourse, errCreation := courseService.Create(&courseModel)
    if errCreation != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": errCreation.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": createdCourse, "msg": "Created successfully"})
}

func DeleteAll(c *gin.Context) {
	courseService := services.NewCourseService(&gorm.DB{})
	err := courseService.DeleteAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Deleted successfully"})
}

func ReviewById(c *gin.Context) {
    idParam := c.Param("id")
    user := c.MustGet("user").(*models.User)

    courseService := services.NewCourseService(&gorm.DB{})
    reviewService := services.NewReviewService(&gorm.DB{})

    var reviewDto dto.ReviewCreateDto
    if err := c.ShouldBindJSON(&reviewDto); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Find the course by ID
    course, err := courseService.FindById(idParam)
    if err != nil || course == nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    checkReviewResponse, _ := reviewService.CheckReview(course.ID, user.ID)
	if err != nil{
		c.JSON(http.StatusConflict, gin.H{
            "error": "Internal server error",
        })
        return
	}
    if checkReviewResponse != nil {
        c.JSON(http.StatusConflict, gin.H{
            "error": "Already Reviewed",
        })
        return
    }

    // Add a new review for the user and course
    errReview := reviewService.AddReview(course, user, reviewDto.Rating, reviewDto.Description )
    if errReview != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": errReview.Error(),
        })
        return
    }
    c.JSON(http.StatusOK, gin.H{"msg": "Review created successfully"})
}



// ***************** Files ******************* // 

func SubscribeToCourse(c *gin.Context) {
	idParam := c.Param("id")

	userAuth, _ := c.Get("user")
	userParam, ok := userAuth.(models.User)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Check if the course exists
	courseService := services.NewCourseService(&gorm.DB{})
	course, err := courseService.FindById(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userService := services.NewUserService(&gorm.DB{})
	user, err := userService.FindById(userParam.ID, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	subscriptionService := services.NewSubscriptionService(&gorm.DB{})
	isSubscribed, err := subscriptionService.IsUserSubscribedToCourse(user.ID, course.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if isSubscribed {
		c.JSON(http.StatusConflict, gin.H{"error": "User is already subscribed to this course"})
		return
	}

	// Subscribe the user to the course
	if err := courseService.SubscribeUserToCourse(user, course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Subscribed to course sucessfully"})
}

func PrivateVideoUploadHandler(c *gin.Context) {
	formKey := "file"

	file, err := c.FormFile(formKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File not provided"})
		return
	}

	uploadPath := "./private"
	if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
		if os.IsPermission(err) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
		}
		return
	}

	newFileNameWithExt := uuid.New().String() + filepath.Ext(file.Filename)
	filePath := filepath.Join(uploadPath, newFileNameWithExt)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save the uploaded file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Private video uploaded successfully",
		"newFileName":  newFileNameWithExt,
	})
}

func PrivateVideoStreamHandler(c *gin.Context) {
	videoFileName := c.Query("filename")

	videoPath := "./private"
	videoFilePath := filepath.Join(videoPath, videoFileName)


	_, err := os.Stat(videoFilePath)
	if os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
		return
	}


	c.File(videoFilePath)
}