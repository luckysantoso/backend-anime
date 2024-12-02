package bootstrap

import (
	"gin-gorm/configs"
	"gin-gorm/configs/app_config"
	"gin-gorm/database"
	"gin-gorm/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func BootstrapApp() {
	// Load .env file
	err := godotenv.Load()
  	if err != nil {
    	log.Println("Error loading .env file")
  	}
	// Init configs
	configs.InitConfigs()

	// Connect to database
	database.ConnectDatabase()

	// Init gin Engine
	app := gin.Default()
	// Init routes
	routes.InitRoutes(app)

	// Run app
	app.Run(app_config.PORT)
}