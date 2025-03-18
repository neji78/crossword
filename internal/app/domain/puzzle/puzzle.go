package puzzle

import "time"

// Puzzle represents the crossword puzzle entity.
type Puzzle struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Width       int       `json:"width"`
	Height      int       `json:"height"`
	Complexity  string    `json:"complexity"`
	Words       string    `json:"words"` // Store words as JSON or simple CSV format
	CreatedAt   time.Time `json:"created_at"`
}