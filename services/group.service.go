package services

import (
	dto "github.com/anthomir/GoProject/dtos"
	"github.com/anthomir/GoProject/initializers"
	"github.com/anthomir/GoProject/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GroupService struct {
	DB *gorm.DB
}

func NewGroupService(db *gorm.DB) *GroupService {
	return &GroupService{DB: db}
}

func (s *GroupService) FindById(id string) (*models.Group, error) {
	var group models.Group

	result := initializers.DB.Where("id = ?", id).First(&group)
	if result.Error != nil {
		return nil, result.Error
	}
	return &group, nil
}

func (s *GroupService) CreateGroupWithUsers(groupDto dto.GroupInternalDto) (*models.Group, error) {
	var users []models.User

	if err := initializers.DB.Where("id IN (?)", groupDto.Users).Find(&users).Error; err != nil {
		return nil, err
	}

	newGroup := models.Group{
        CourseID: groupDto.CourseID,
		Title:       groupDto.Title,
		Description: groupDto.Description,
		CreatedBy:   groupDto.CreatedBy,
		Users:       users,
	}

	if err := initializers.DB.Create(&newGroup).Error; err != nil {
		return nil, err
	}

	if err := initializers.DB.Model(&newGroup).Preload("Users").Find(&newGroup).Error; err != nil {
		return nil, err
	}

	return &newGroup, nil
}

func (s *GroupService) DeleteAll() error {
	result := initializers.DB.Unscoped().Where("1 = 1").Delete(&models.Group{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *GroupService) DeleteById(id uuid.UUID) error {
	var group models.Group

	if err := initializers.DB.Where("id = ?", id).First(&group).Error; err != nil {
		return err
	}

	if err := initializers.DB.Delete(&group).Error; err != nil {
		return err
	}

	return nil
}

func (s *GroupService) GetGroupsByUserID(userID uuid.UUID) ([]models.Group, error) {
    var user models.User
    if err := initializers.DB.Where("id = ?", userID).Preload("Groups").First(&user).Error; err != nil {
        return nil, err
    }

    return user.Groups, nil
}

func (groupService *GroupService) AddUsersToGroup(groupID uuid.UUID, userIds []uuid.UUID) (*models.Group, error) {
    // Fetch the group by ID
    var group models.Group
    if err := initializers.DB.Preload("Users").Where("id = ?", groupID).First(&group).Error; err != nil {
        return nil, err
    }

    // Fetch the users by their IDs
    var users []models.User
    if err := initializers.DB.Find(&users, userIds).Error; err != nil {
        return nil, err
    }

    // Check for existing users and add only those not already in the group
    for _, newUser := range users {
        exists := false
        for _, existingUser := range group.Users {
            if existingUser.ID == newUser.ID {
                exists = true
                break
            }
        }
        if !exists {
            group.Users = append(group.Users, newUser)
        }
    }

    // Save the updated group back to the database
    if err := initializers.DB.Save(&group).Error; err != nil {
        return nil, err
    }

    return &group, nil
}

func (groupService *GroupService) RemoveUsersFromGroup(groupID uuid.UUID, userIDs []uuid.UUID) (*models.Group, error) {
    // Fetch the group by ID
    var group models.Group
    if err := initializers.DB.Preload("Users").Where("id = ?", groupID).First(&group).Error; err != nil {
        return nil, err
    }

    // Create a map to efficiently check the existence of user IDs
    userIDsMap := make(map[uuid.UUID]bool)
    for _, id := range userIDs {
        userIDsMap[id] = true
    }

    // Filter out the users with the specified IDs
    var updatedUsers []models.User
    for _, user := range group.Users {
        if !userIDsMap[user.ID] {
            updatedUsers = append(updatedUsers, user)
        }
    }

    // Update the group's Users field with the filtered users
    group.Users = updatedUsers

    // Save the updated group back to the database
    if err := initializers.DB.Save(&group).Error; err != nil {
        return nil, err
    }

    return &group, nil
}

func (s *GroupService) GetGroupDetails(groupID string) (*dto.GroupDetails, error) {
	chapterService := NewChapterService(&gorm.DB{})

    var group models.Group
    result := initializers.DB.Model(&models.Group{}).Preload("Users").Where("id = ?", groupID).First(&group)
    if result.Error != nil {
        return nil, result.Error
    }

    var groupDetails dto.GroupDetails
    groupDetails.Group = group
    groupDetails.Users = make([]dto.UserWithScore, len(group.Users))

    for i, user := range group.Users {
        completedChaptersCount, err := chapterService.GetCompletedChaptersCount(user.ID, group.CourseID)
        if err != nil {
            return nil, err
        }

        // Calculate total chapters count for the specific course
        totalChaptersCount, err := chapterService.GetTotalChaptersCountForCourse(group.CourseID)
        if err != nil {
            return nil, err
        }

        if totalChaptersCount == 0 {
            groupDetails.Users[i] = dto.UserWithScore{User: user, Score: 0}
        } else {
            score := float64(completedChaptersCount) / float64(totalChaptersCount)
            groupDetails.Users[i] = dto.UserWithScore{
                User:  user,
                Score: score,
            }
        }
    }

    return &groupDetails, nil
}
