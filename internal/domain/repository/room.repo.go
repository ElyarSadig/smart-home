package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/elyarsadig/smart-home-iot/internal/domain/models"
)

type RoomRepository struct {
	db *sql.DB
}

func NewRoom(db *sql.DB) models.RoomRepository {
	return &RoomRepository{
		db: db,
	}
}

func (r *RoomRepository) GetByName(ctx context.Context, name string) (*models.Room, error) {
	query := `SELECT * FROM rooms WHERE name = ?`
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error preparing: %v", err)
		return nil, err
	}
	row := stmt.QueryRowContext(ctx, name)
	var room models.Room
	err = row.Scan(&room)
	if err != nil {
		log.Printf("Error scaning row: %v", err)
		return nil, err
	}
	return &room, nil
}
