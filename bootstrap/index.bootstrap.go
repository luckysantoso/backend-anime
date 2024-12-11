package bootstrap

import (
	"gin-gorm/configs"
	"gin-gorm/configs/app_config"
	"gin-gorm/database"
	"gin-gorm/routes"
	"log"
	"time"

	"github.com/gin-contrib/cors"
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

	// Add CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://example.com"}, // Ganti dengan domain front-end Anda
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Init routes
	routes.InitRoutes(app)

	// Run app
	app.Run(app_config.PORT)
}
