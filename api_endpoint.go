package main

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// POST /generate - Create a new crossword puzzle
func generatePuzzleHandler(c *fiber.Ctx) error {
	var request struct {
		Width      int    `json:"width"`
		Height     int    `json:"height"`
		Complexity string `json:"complexity"`
	}

	// Parse the request body
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Get user ID from JWT or session
	userID := 1 // For now, we'll hardcode it. You should fetch from the JWT

	// Generate puzzle
	puzzle, err := generatePuzzle(request.Width, request.Height, request.Complexity, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate puzzle"})
	}

	return c.JSON(fiber.Map{"puzzle_id": puzzle.ID, "message": "Puzzle generated successfully"})
}

// GET /puzzles/{id} - Retrieve a specific puzzle
func getPuzzleByIDHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	puzzleID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid puzzle ID"})
	}

	// Fetch puzzle from DB
	var puzzle CrosswordPuzzle
	err = db.Get(&puzzle, "SELECT * FROM crossword_puzzles WHERE id = ?", puzzleID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Puzzle not found"})
	}

	return c.JSON(fiber.Map{"puzzle": puzzle})
}

// GET /puzzles/user - Get all puzzles of a user
func getPuzzlesByUserHandler(c *fiber.Ctx) error {
	// Get user ID from JWT or session
	userID := 1 // For now, we'll hardcode it. You should fetch from the JWT

	// Fetch all puzzles for the user
	var puzzles []CrosswordPuzzle
	err := db.Select(&puzzles, "SELECT * FROM crossword_puzzles WHERE user_id = ?", userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve puzzles"})
	}

	return c.JSON(fiber.Map{"puzzles": puzzles})
}

func setupRoutes(app *fiber.App) {
	// Routes for crossword puzzles
	app.Post("/generate", generatePuzzleHandler)
	app.Get("/puzzles/:id", getPuzzleByIDHandler)
	app.Get("/puzzles/user", getPuzzlesByUserHandler)
}
