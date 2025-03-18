package puzzle

type PuzzleRepository interface {
	Save(puzzle Puzzle) error
	FindByID(id int) (Puzzle, error)
	FindAll() ([]Puzzle, error)
	Delete(id int) error
}