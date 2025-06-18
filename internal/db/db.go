package db

import (
	"database/sql"
	"log"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Init() {
	var err error
	DB, err := sql.Open("sqlite", "./forum.db")

	if err != nil {
		log.Fatal(err)
	}
	initSchema()
}

func initSchema() {
	schema := `
	CREATE TABLE IF NOT EXISTS threads (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS comments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		thread_id INTEGER NOT NULL,
		content TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(thread_id) REFERENCES threads(id) ON DELETE CASCADE
	);`

	_, err := DB.Exec(schema)
	if err != nil {
		log.Fatalf("Błąd tworzenia schematu bazy danych: %v", err)
	}
}
