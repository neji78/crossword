#!/bin/bash

# This script is for migrating the database schema for the crossword application.

set -e

# Define the database file
DB_FILE="crossword.db"

# Create the database and tables if they do not exist
sqlite3 $DB_FILE <<EOF
CREATE TABLE IF NOT EXISTS users (
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
);

CREATE TABLE IF NOT EXISTS crossword_words (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    word TEXT,
    clue TEXT,
    difficulty TEXT,
    theme TEXT,
    language TEXT,
    added_by INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS puzzle_words (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    puzzle_id INTEGER,
    word_id INTEGER,
    position_x INTEGER,
    position_y INTEGER,
    direction TEXT,
    FOREIGN KEY(puzzle_id) REFERENCES crossword_puzzles(id),
    FOREIGN KEY(word_id) REFERENCES crossword_words(id)
);
EOF

echo "Database migration completed successfully."