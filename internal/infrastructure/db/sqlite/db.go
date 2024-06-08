package sqlite

import (
	"database/sql"
	"log"

	"github.com/elyarsadig/smart-home-iot/config"
	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

func InitDB(cfg *config.DatabaseConfig) *sql.DB {
	db, err := sql.Open("sqlite3", cfg.FilePath)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	err = migrate(db)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to SQLite database!")
	return db
}
