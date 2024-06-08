package models

import "time"

type Device struct {
	ID                     int       `json:"id"`
	Name                   string    `json:"name"`
	Status                 bool      `json:"status"`
	ActiveFrom             time.Time `json:"active_from"`
	ActiveUntil            time.Time `json:"active_until"`
	RoomID                 int       `json:"room_id"`
	EnergyConsumingPerHour float64   `json:"energy_consuming_per_hour"`
}
