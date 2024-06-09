package sqlite

import (
	"database/sql"

	"github.com/elyarsadig/smart-home-iot/internal/domain/models"
)

func seed(db *sql.DB) error {
	livingRoom := models.Room{
		Name:              "Living Room",
		Temperature:       20.4,
		Humidity:          34.21,
		IsFrontDoorLocked: false,
		IsCameraActive:    false,
		IsAlarmActive:     false,
	}
	result, err := db.Exec("INSERT INTO rooms (name, temperature, humidity, is_front_door_locked, is_camera_active, is_alarm_active) VALUES (?, ?, ?, ?, ?, ?)", livingRoom.Name, livingRoom.Temperature, livingRoom.Humidity, livingRoom.IsFrontDoorLocked, livingRoom.IsCameraActive, livingRoom.IsAlarmActive)
	if err != nil {
		return err
	}
	roomID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	err = insertDevices(db, int(roomID))
	if err != nil {
		return err
	}
	return nil
}

func insertDevices(db *sql.DB, roomID int) error {
	devices := []models.Device{
		{Name: "Smart TV", IsActive: false, RoomID: roomID, EnergyConsumingPerHour: 150.55},
		{Name: "WiFi Router", IsActive: false, RoomID: roomID, EnergyConsumingPerHour: 120.33},
		{Name: "Air Conditioner", IsActive: false, RoomID: roomID, EnergyConsumingPerHour: 367.32},
		{Name: "Lights", IsActive: false, RoomID: roomID, EnergyConsumingPerHour: 90.43},
		{Name: "Air Humidifier", IsActive: false, RoomID: roomID, EnergyConsumingPerHour: 320.64},
		{Name: "Speaker", IsActive: false, RoomID: roomID, EnergyConsumingPerHour: 180.78},
	}

	for _, device := range devices {
		_, err := db.Exec("INSERT INTO devices (name, is_active, room_id, energy_consuming_per_hour) VALUES(?, ?, ?, ?)",
			device.Name, device.IsActive, device.RoomID, device.EnergyConsumingPerHour)
		if err != nil {
			return err
		}
	}
	return nil
}
