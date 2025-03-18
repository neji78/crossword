package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	_ "github.com/mattn/go-sqlite3"
)

var db *sqlx.DB

func InitDB() {
	var err error
	db, err = sqlx.Connect(viper.GetString("database.type"), viper.GetString("database.path"))
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	schema := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE,
		email TEXT UNIQUE,
		password TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	CREATE TABLE IF NOT EXISTS crossword_puzzles (
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

func GetDB() *sqlx.DB {
	return db
}