package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/elyarsadig/smart-home-iot/internal/domain/models"
)

type DeviceRepository struct {
	db *sql.DB
}

func NewDevice(db *sql.DB) models.DeviceRepository {
	return &DeviceRepository{
		db: db,
	}
}

func (r *DeviceRepository) GetByName(ctx context.Context, name string) (*models.Device, error) {
	query := `SELECT * FROM devices WHERE name = ?`
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error preparing: %v", err)
		return nil, err
	}
	row := stmt.QueryRowContext(ctx, name)
	var device models.Device
	err = row.Scan(&device)
	if err != nil {
		log.Printf("Error scaning row: %v", err)
		return nil, err
	}
	return &device, nil
}
