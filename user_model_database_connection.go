package main

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// DB instance
var db *sqlx.DB

type User struct {
	ID        int       `db:"id"`
	Username  string    `db:"username"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
}

func initDB() {
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
	);`
	db.MustExec(schema)
}
