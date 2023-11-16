package initializers

import (
	"fmt"
	"log"

	"github.com/anthomir/GoProject/models"
)

func SyncDatabase() {

	err := DB.AutoMigrate(
		&models.Chapter{},
		&models.Course{},
		&models.Group{}, 
		&models.Review{},
		&models.Subscription{},
		&models.User{},
		&models.UserChapter{}, 
		&models.UserGroup{},

	)
	if err != nil {
		log.Fatal(err)
		fmt.Println("*********************************")
		fmt.Println("*********************************")
		fmt.Println("*********************************")
		fmt.Println("*********************************")
		fmt.Println("*********************************")
		fmt.Println("*********************************")
		fmt.Println("*********************************")
		fmt.Println(err)
		fmt.Println("*********************************")
		fmt.Println("*********************************")
		fmt.Println("*********************************")
		fmt.Println("*********************************")
		fmt.Println("*********************************")
		fmt.Println("*********************************")
		fmt.Println("*********************************")
	 }
}
