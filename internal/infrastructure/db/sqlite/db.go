package sqlite

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./internal/infrastructure/db/sqlite/smart_home.db")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Connected to SQLite database!")
	return db, nil
}
