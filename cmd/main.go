package main

import (
	"log"

	"github.com/elyarsadig/smart-home-iot/internal/infrastructure/db/sqlite"
)

func main() {
	db, err := sqlite.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = sqlite.CreateTables(db)
	if err != nil {
		log.Fatal(err)
	}
}
