package models

import "time"

// Puzzle represents a crossword puzzle entity.
type Puzzle struct {
	ID          int       `json:"id" db:"id"`
	UserID      int       `json:"user_id" db:"user_id"`
	Width       int       `json:"width" db:"width"`
	Height      int       `json:"height" db:"height"`
	Complexity  string    `json:"complexity" db:"complexity"`
	Words       string    `json:"words" db:"words"` // Store words as JSON or simple CSV format
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}