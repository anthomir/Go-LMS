package services

import (
	"errors"
	"fmt"

	"github.com/anthomir/GoProject/initializers"
	"github.com/anthomir/GoProject/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) GetByEmail(email string) (*models.User, error) {
	var user models.User
	result := initializers.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (s *UserService) Create(newUser *models.User) (models.User, error) {

	_, err := s.GetByEmail(newUser.Email)
	if err == nil {
		return models.User{}, errors.New("user already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
	if err != nil {
		return models.User{}, errors.New("unable to hash password")
	}

	newUser.Password = string(hash)

	result := initializers.DB.Create(&newUser)
	if result.Error != nil {
		return models.User{}, errors.New("Unable to create User: " + result.Error.Error())
	}

	newUser.Password = ""
	return *newUser, nil
}

func (s *UserService) Authenticate(email, password string) (*models.User, error) {

	// Find the user by email
	user, err := s.GetByEmail(email)
	if user == nil {
		fmt.Print("************* ERR: " + err.Error())

		return nil, nil
	}

	// Verify the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err // Password doesn't match
	}

	return user, nil
}

func (s *UserService) GetAll() *[]models.User {
	var users []models.User

	result := initializers.DB.Find(&users)
	if result.Error != nil {
		return nil
	}

	return &users
}

func (s *UserService) FindAll(query string, skip int, take int) (*[]models.User, error) {
	var users []models.User

	result := initializers.DB.Where(query).Offset(skip). // Skip records (page-1)*perPage
										Limit(take). // Take a limited number of records
										Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return &users, nil
}

func (s *UserService) FindById(id uuid.UUID, preloadProperties []string) (*models.User, error) {
	var user models.User

	query := initializers.DB.Model(&user)
	for _, prop := range preloadProperties {
		query = query.Preload(prop)
	}

	result := query.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
