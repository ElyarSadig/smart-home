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

func (r *RoomRepository) GetById(ctx context.Context, id int) (*models.Room, error) {
	query := `SELECT * FROM rooms WHERE id = ?`
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error preparing: %v", err)
		return nil, config.InternalServerError
	}
	row := stmt.QueryRowContext(ctx, id)
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

func (r *RoomRepository) UpdateRoomStatus(ctx context.Context, id int, frontDoorStatus, cameraStatus, alarmStatus bool) error {
	room, err := r.GetById(ctx, id)
	if err != nil {
		return err
	}
	room.IsFrontDoorLocked = frontDoorStatus
	room.IsCameraActive = cameraStatus
	room.IsAlarmActive = alarmStatus
	query := `UPDATE rooms SET is_front_door_locked = ?, is_camera_active = ?, is_alarm_active = ? WHERE id = ?`
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error preparing: %v", err)
		return config.InternalServerError
	}
	_, err = stmt.ExecContext(ctx, room.IsFrontDoorLocked, room.IsCameraActive, room.IsAlarmActive, room.ID)
	if err != nil {
		log.Printf("Error executing: %v", err)
		return config.InternalServerError
	}
	return nil
}

func (r *RoomRepository) GetRoomInActiveDevices(ctx context.Context, id int) ([]models.Device, error) {
	room, err := r.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	query := `SELECT * FROM devices WHERE room_id = ? AND is_active = false`
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error preparing: %v", err)
		return nil, config.InternalServerError
	}
	rows, err := stmt.QueryContext(ctx, room.ID)
	if err != nil {
		log.Printf("Error executing: %v", err)
		return nil, config.InternalServerError
	}
	devices := make([]models.Device, 0, 6)
	for rows.Next() {
		err = rows.Err()
		if err != nil {
			log.Printf("Error iterating rows: %v", err)
			return nil, config.InternalServerError
		}
		var device models.Device
		err = rows.Scan(
			&device.ID,
			&device.Name,
			&device.IsActive,
			&device.UpdatedAT,
			&device.RoomID,
			&device.EnergyConsumingPerHour,
		)
		if err != nil {
			log.Printf("Error scanning: %v", err)
			return nil, config.InternalServerError
		}
		devices = append(devices, device)
	}
	err = rows.Close()
	if err != nil {
		log.Printf("Error closing rows: %v", err)
	}
	return devices, nil
}
