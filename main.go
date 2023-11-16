package main

import (
	"github.com/anthomir/GoProject/initializers"
	"github.com/anthomir/GoProject/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	router := gin.Default()
	routes.SetupRoutes(router)

	router.Run()
}
