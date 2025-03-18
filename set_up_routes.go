package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	loadConfig()
	initDB()

	app := fiber.New()

	// Authentication routes
	app.Post("/register", registerUser)
	app.Post("/login", loginUser)

	// Protected route
	app.Get("/protected", authMiddleware, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Welcome to the protected area"})
	})

	app.Listen(":" + viper.GetString("server.port"))
}
