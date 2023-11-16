package initializers

import "github.com/anthomir/GoProject/models"

func SyncDatabase() {

	DB.AutoMigrate(
		&models.Course{},
		&models.Subscription{},
		&models.Review{},
		&models.Chapter{},
		&models.UserGroup{},
		&models.Scoreboard{}, 
		&models.Group{}, 
		&models.User{},
	)
}
