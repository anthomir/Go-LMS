package userHandler

import (
	"net/http"
	"strconv"

	dto "github.com/anthomir/GoProject/dtos"
	"github.com/anthomir/GoProject/jwt"
	"github.com/anthomir/GoProject/models"
	"github.com/anthomir/GoProject/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserSignUp(c *gin.Context) {
	userService := services.NewUserService(&gorm.DB{})
	var userModel models.User

	if err := c.ShouldBindJSON(&userModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if err := userModel.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := userService.Create(&userModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func UserLogin(c *gin.Context) {
	var loginDTO dto.UserLoginDTO

	if err := c.ShouldBindJSON(&loginDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	if err := loginDTO.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Authenticate the user using the UserService
	userService := services.NewUserService(&gorm.DB{})
	user, err := userService.Authenticate(loginDTO.Email, loginDTO.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Authentication failed"})
		return
	}

	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate a JWT for the user
	tokenString, err := jwt.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT"})
		return
	}

	user.Password = ""

	// Adding Token to Cookie
	// c.SetSameSite(http.SameSiteLaxMode)
	// c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{"data": map[string]interface{}{"user": user, "token": tokenString}})

}
func GetUserProfile(c *gin.Context) {
	data, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func FindAllTeachers(c *gin.Context) {
	userService := services.NewUserService(&gorm.DB{})
	skipParam, err1 := strconv.Atoi(c.DefaultQuery("skip", "0"))
	takeParam, err2 := strconv.Atoi(c.DefaultQuery("take", "10"))
	queryParam:= c.DefaultQuery("query", "")

	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid skip or take parameters"})
		return
	}

	response, err := userService.FindAll(queryParam, skipParam, takeParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Authentication failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

