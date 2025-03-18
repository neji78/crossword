package api

import (
	"github.com/gofiber/fiber/v2"
	"crossword/internal/app/api/handlers"
)

func SetupRoutes(app *fiber.App) {
	// User routes
	app.Post("/register", handlers.RegisterUser)
	app.Post("/login", handlers.LoginUser)

	// Protected routes
	app.Use(handlers.AuthMiddleware)
	app.Post("/generate", handlers.GeneratePuzzle)
	app.Get("/puzzles/:id", handlers.GetPuzzleByID)
	app.Get("/puzzles/user", handlers.GetPuzzlesByUser)
}