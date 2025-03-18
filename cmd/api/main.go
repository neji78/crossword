package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"

	"crossword/internal/app/api/routes"
	"crossword/internal/app/infrastructure/db"
)

func main() {
	// Load configuration
	if err := loadConfig(); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Initialize database
	if err := db.InitDB(); err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// Create a new Fiber app
	app := fiber.New()

	// Setup routes
	routes.SetupRoutes(app)

	// Start the server
	if err := app.Listen(":" + viper.GetString("server.port")); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func loadConfig() error {
	viper.SetConfigFile("config.yml")
	viper.AutomaticEnv()

	return viper.ReadInConfig()
}