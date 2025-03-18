package handlers

import (
	"github.com/gofiber/fiber/v2"
	"crossword/internal/app/services"
	"net/http"
)

type PuzzleHandler struct {
	PuzzleService services.PuzzleService
}

func NewPuzzleHandler(puzzleService services.PuzzleService) *PuzzleHandler {
	return &PuzzleHandler{PuzzleService: puzzleService}
}

// GeneratePuzzle handles the request to generate a new crossword puzzle
func (h *PuzzleHandler) GeneratePuzzle(c *fiber.Ctx) error {
	var request struct {
		Width      int    `json:"width"`
		Height     int    `json:"height"`
		Complexity string `json:"complexity"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	puzzle, err := h.PuzzleService.GeneratePuzzle(request.Width, request.Height, request.Complexity)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate puzzle"})
	}

	return c.JSON(fiber.Map{"puzzle_id": puzzle.ID, "message": "Puzzle generated successfully"})
}

// GetPuzzleByID handles the request to retrieve a specific puzzle by ID
func (h *PuzzleHandler) GetPuzzleByID(c *fiber.Ctx) error {
	id := c.Params("id")
	puzzle, err := h.PuzzleService.GetPuzzleByID(id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Puzzle not found"})
	}

	return c.JSON(puzzle)
}

// GetPuzzlesByUser handles the request to retrieve all puzzles for a specific user
func (h *PuzzleHandler) GetPuzzlesByUser(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int) // Assuming userID is set in the middleware
	puzzles, err := h.PuzzleService.GetPuzzlesByUser(userID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve puzzles"})
	}

	return c.JSON(puzzles)
}