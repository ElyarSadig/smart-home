package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/elyarsadig/smart-home-iot/config"
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
		return nil, config.InternalServerError
	}
	row := stmt.QueryRowContext(ctx, name)
	var device models.Device
	err = row.Scan(
        &device.ID,
        &device.Name,
        &device.IsActive,
        &device.ActiveFrom,
        &device.RoomID,
        &device.EnergyConsumingPerHour,
    )
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, config.NotFoundError
		}
		log.Printf("Error scaning row: %v", err)
		return nil, config.InternalServerError
	}
	return &device, nil
}
