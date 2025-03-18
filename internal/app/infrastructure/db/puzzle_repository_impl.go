package db

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/yourusername/crossword/internal/app/domain/puzzle"
)

type puzzleRepositoryImpl struct {
	db *sqlx.DB
}

func NewPuzzleRepository(db *sqlx.DB) puzzle.PuzzleRepository {
	return &puzzleRepositoryImpl{db: db}
}

func (r *puzzleRepositoryImpl) Create(p *puzzle.Puzzle) error {
	query := `INSERT INTO crossword_puzzles (user_id, width, height, complexity, words, created_at) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query, p.UserID, p.Width, p.Height, p.Complexity, p.Words, time.Now())
	return err
}

func (r *puzzleRepositoryImpl) FindByID(id int) (*puzzle.Puzzle, error) {
	var p puzzle.Puzzle
	query := `SELECT * FROM crossword_puzzles WHERE id = ?`
	err := r.db.Get(&p, query, id)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *puzzleRepositoryImpl) FindByUserID(userID int) ([]puzzle.Puzzle, error) {
	var puzzles []puzzle.Puzzle
	query := `SELECT * FROM crossword_puzzles WHERE user_id = ?`
	err := r.db.Select(&puzzles, query, userID)
	return puzzles, err
}