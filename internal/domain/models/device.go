package models

import (
	"time"
)

type Device struct {
	ID                     int        `json:"id"`
	Name                   string     `json:"name"`
	IsActive               bool       `json:"is_active"`
	UpdatedAT              time.Time `json:"updated_at"`
	RoomID                 int        `json:"room_id"`
	EnergyConsumingPerHour float64    `json:"energy_consuming_per_hour"`
}
