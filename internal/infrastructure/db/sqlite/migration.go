package sqlite

import (
	"database/sql"
	"fmt"
	"log"
)

var tableCreationQueries = map[string]string{
	"devices": DEVICE_TABLE_QUERY,
	"rooms":   ROOM_TABLE_QUERY,
}

func CreateTables(db *sql.DB) error {
	for tableName, createQuery := range tableCreationQueries {
		if _, err := db.Exec(createQuery); err != nil {
			return fmt.Errorf("failed to create table %s: %w", tableName, err)
		}
		log.Printf("Table %s created successfully!", tableName)
	}
	return nil
}
