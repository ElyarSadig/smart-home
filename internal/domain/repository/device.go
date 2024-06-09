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

func (r *DeviceRepository) GetById(ctx context.Context, id int) (*models.Device, error) {
	query := `SELECT * FROM devices WHERE id = ?`
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error preparing: %v", err)
		return nil, config.InternalServerError
	}
	row := stmt.QueryRowContext(ctx, id)
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
