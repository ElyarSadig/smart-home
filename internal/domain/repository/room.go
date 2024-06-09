package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/elyarsadig/smart-home-iot/config"
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
		return nil, config.InternalServerError
	}
	row := stmt.QueryRowContext(ctx, name)
	var room models.Room
	err = row.Scan(
		&room.ID,
		&room.Name,
		&room.Temperature,
		&room.Humidity,
		&room.IsFrontDoorLocked,
		&room.IsCameraActive,
		&room.IsAlarmActive,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, config.NotFoundError
		}
		log.Printf("Error scaning row: %v", err)
		return nil, config.InternalServerError
	}
	return &room, nil
}
