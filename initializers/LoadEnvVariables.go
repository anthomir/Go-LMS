package initializers

import (
	"fmt"
	"log"

	"github.com/anthomir/GoProject/colors"
	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(colors.Red + "Error loading Environment Variables: " + err.Error() + "" + colors.Reset)
	}
	fmt.Println(colors.Green + "Env Variables Loaded" + colors.Reset)
}
