package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/yourusername/crossword/internal/app/domain/user"
)

type userRepositoryImpl struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) user.UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) Create(user *user.User) error {
	_, err := r.db.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)", user.Username, user.Email, user.Password)
	return err
}

func (r *userRepositoryImpl) FindByID(id int) (*user.User, error) {
	var u user.User
	err := r.db.Get(&u, "SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepositoryImpl) FindByUsername(username string) (*user.User, error) {
	var u user.User
	err := r.db.Get(&u, "SELECT * FROM users WHERE username = ?", username)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepositoryImpl) Update(user *user.User) error {
	_, err := r.db.Exec("UPDATE users SET email = ?, password = ? WHERE id = ?", user.Email, user.Password, user.ID)
	return err
}

func (r *userRepositoryImpl) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}