package models

import "time"

// User represents the user entity in the system.
type User struct {
	ID        int       `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"-" db:"password"` // Password should not be exposed
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}