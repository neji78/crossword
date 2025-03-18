// Puzzle model
type CrosswordPuzzle struct {
	ID          int       `db:"id"`
	UserID      int       `db:"user_id"`
	Width       int       `db:"width"`
	Height      int       `db:"height"`
	Complexity  string    `db:"complexity"`
	Words       string    `db:"words"` // Store words as JSON or simple CSV format
	CreatedAt   time.Time `db:"created_at"`
}

// Create the crossword_puzzles table
func initPuzzleTable() {
	schema := `CREATE TABLE IF NOT EXISTS crossword_puzzles (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		width INTEGER,
		height INTEGER,
		complexity TEXT,
		words TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);`
	db.MustExec(schema)
}
