package services

import (
	"encoding/json"
	"errors"
	"math/rand"
	"time"

	"crossword/internal/app/domain/puzzle"
	"crossword/pkg/models"
)

type PuzzleService struct {
	repo puzzle.PuzzleRepository
}

func NewPuzzleService(repo puzzle.PuzzleRepository) *PuzzleService {
	return &PuzzleService{repo: repo}
}

func (s *PuzzleService) GeneratePuzzle(width, height int, complexity string, userID int) (models.Puzzle, error) {
	words := s.fetchWordsByComplexity(complexity)
	grid := s.createEmptyGrid(width, height)

	var placedWords []models.Word
	for _, word := range words {
		if s.placeWordInGrid(grid, word) {
			placedWords = append(placedWords, models.Word{Word: word, Position: models.Position{X: rand.Intn(width), Y: rand.Intn(height)}})
		}
	}

	puzzleData, err := json.Marshal(placedWords)
	if err != nil {
		return models.Puzzle{}, err
	}

	puzzle := models.Puzzle{
		UserID:     userID,
		Width:      width,
		Height:     height,
		Complexity: complexity,
		Words:      string(puzzleData),
		CreatedAt:  time.Now(),
	}

	if err := s.repo.Save(puzzle); err != nil {
		return models.Puzzle{}, err
	}

	return puzzle, nil
}

func (s *PuzzleService) fetchWordsByComplexity(complexity string) []string {
	switch complexity {
	case "easy":
		return []string{"cat", "dog", "fish"}
	case "medium":
		return []string{"elephant", "giraffe", "monkey"}
	default:
		return []string{"hippopotamus", "extraordinary", "pneumonia"}
	}
}

func (s *PuzzleService) createEmptyGrid(width, height int) [][]string {
	grid := make([][]string, height)
	for i := range grid {
		grid[i] = make([]string, width)
	}
	return grid
}

func (s *PuzzleService) placeWordInGrid(grid [][]string, word string) bool {
	for i := 0; i < len(grid); i++ {
		for j := 0; j <= len(grid[i])-len(word); j++ {
			if s.canPlaceWord(grid, i, j, word) {
				for k := 0; k < len(word); k++ {
					grid[i][j+k] = string(word[k])
				}
				return true
			}
		}
	}
	return false
}

func (s *PuzzleService) canPlaceWord(grid [][]string, x, y int, word string) bool {
	for i := 0; i < len(word); i++ {
		if grid[x][y+i] != "" {
			return false
		}
	}
	return true
}