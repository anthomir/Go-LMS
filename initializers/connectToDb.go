package initializers

import (
	"fmt"
	"os"

	"github.com/anthomir/GoProject/colors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(colors.Red + "Failed connecting to the Database: " + err.Error() + colors.Reset)
		panic("")
	}

	fmt.Println(colors.Green + "Database Connected Successfully" + colors.Reset)
	DB = db
}
			