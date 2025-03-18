import (
	"encoding/json"
	"math/rand"
	"strings"
	"time"
)

// Word structure
type Word struct {
	Word     string `json:"word"`
	Position struct {
		X int `json:"x"`
		Y int `json:"y"`
	} `json:"position"`
}

// Generate crossword puzzle
func generatePuzzle(width, height int, complexity string, userID int) (CrosswordPuzzle, error) {
	// Fetch words based on complexity
	words := fetchWordsByComplexity(complexity)

	// Initialize grid
	grid := createEmptyGrid(width, height)

	// Try to place words randomly
	var placedWords []Word
	for _, word := range words {
		if placeWordInGrid(grid, word) {
			placedWords = append(placedWords, Word{Word: word, Position: struct{ X, Y int }{X: rand.Intn(width), Y: rand.Intn(height)}})
		}
	}

	// Store the puzzle in DB
	wordData, _ := json.Marshal(placedWords)
	puzzle := CrosswordPuzzle{
		UserID:     userID,
		Width:      width,
		Height:     height,
		Complexity: complexity,
		Words:      string(wordData),
	}

	// Save the puzzle in DB
	_, err := db.Exec("INSERT INTO crossword_puzzles (user_id, width, height, complexity, words) VALUES (?, ?, ?, ?, ?)",
		puzzle.UserID, puzzle.Width, puzzle.Height, puzzle.Complexity, puzzle.Words)

	return puzzle, err
}

// Helper function to fetch words based on complexity
func fetchWordsByComplexity(complexity string) []string {
	// Fetch words from DB or from a static list based on complexity
	// Example: Return simple words for 'easy', longer words for 'hard'
	words := []string{}
	if complexity == "easy" {
		words = append(words, "cat", "dog", "fish")
	} else if complexity == "medium" {
		words = append(words, "elephant", "giraffe", "monkey")
	} else {
		words = append(words, "hippopotamus", "extraordinary", "pneumonia")
	}
	return words
}

// Helper function to create an empty grid
func createEmptyGrid(width, height int) [][]string {
	grid := make([][]string, height)
	for i := range grid {
		grid[i] = make([]string, width)
	}
	return grid
}

// Helper function to place a word in the grid
func placeWordInGrid(grid [][]string, word string) bool {
	// Randomly place word (horizontal, vertical, diagonal)
	// Here we're just placing them horizontally for simplicity.
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i])-len(word); j++ {
			if canPlaceWord(grid, i, j, word) {
				for k := 0; k < len(word); k++ {
					grid[i][j+k] = string(word[k])
				}
				return true
			}
		}
	}
	return false
}

// Check if a word can be placed in a specific position
func canPlaceWord(grid [][]string, x, y int, word string) bool {
	for i := 0; i < len(word); i++ {
		if grid[x][y+i] != "" {
			return false
		}
	}
	return true
}
