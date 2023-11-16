package groupHandler

import (
	"net/http"

	dto "github.com/anthomir/GoProject/dtos"
	"github.com/anthomir/GoProject/models"
	"github.com/anthomir/GoProject/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Find Group By Id
func FindById(c *gin.Context) {
	idParam := c.Param("id")
	groupService := services.NewGroupService(&gorm.DB{})

	response, err := groupService.FindById(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

// Find Groups By User
func FindByUser(c *gin.Context) {
	userAuth, _ := c.Get("user")
	user, ok := userAuth.(*models.User)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected user type in context"})
        return
    }

	groupService := services.NewGroupService(&gorm.DB{})
	response, err := groupService.GetGroupsByUserID(user.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func AddUsersToGroup(c *gin.Context) {
    groupService := services.NewGroupService(&gorm.DB{})
    groupID, err := uuid.Parse(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
        return
    }

    var requestBody struct {
        UserIDs []uuid.UUID `json:"userIds"`
    }
    if err := c.ShouldBindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

    updatedGroup, err := groupService.AddUsersToGroup(groupID, requestBody.UserIDs)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, updatedGroup)
}

func RemoveUserFromGroupHandler(c *gin.Context) {
    groupService := services.NewGroupService(&gorm.DB{})

    groupID, err := uuid.Parse(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
        return
    }

    var requestBody struct {
        UserIDs []uuid.UUID `json:"userIds"`
    }
    if err := c.ShouldBindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

    updatedGroup, err := groupService.RemoveUsersFromGroup(groupID, requestBody.UserIDs)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, updatedGroup)
}


func Create(c *gin.Context) {
	var dtoData dto.GroupCreateDto
	user := c.MustGet("user").(*models.User)
	groupService := services.NewGroupService(&gorm.DB{})

	if err := c.ShouldBindJSON(&dtoData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	groupInternalDto := dto.GroupInternalDto{
		CourseID:    dtoData.CourseID,
		Title:       dtoData.Title,
		Description: dtoData.Description,
		Users: 		 dtoData.Users,
		CreatedBy:   user.ID,
	}
	
	createdGroup, err := groupService.CreateGroupWithUsers(groupInternalDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create group"})
		return
	}

	c.JSON(http.StatusCreated, createdGroup)
}

// Delete All Groups
func DeleteAll(c *gin.Context) {
	groupService := services.NewGroupService(&gorm.DB{})
	err := groupService.DeleteAll()
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Deleted successfully"})
}

// Delete All Groups
func DeleteById(c *gin.Context) {
	groupService := services.NewGroupService(&gorm.DB{})
	id := c.Param("id")
	groupID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	// Call the service to delete the group by ID
	if err := groupService.DeleteById(groupID); err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete group"})
		return
	}

	c.JSON(200, gin.H{"message": "Group deleted successfully"})
}

func GetGroupDetailsHandler(c *gin.Context) {
	groupService := services.NewGroupService(&gorm.DB{})
    groupID := c.Param("id")

    groupDetails, err := groupService.GetGroupDetails(groupID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve group details"})
        return
    }

    c.JSON(http.StatusOK, groupDetails)
}