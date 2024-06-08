package models

type Room struct {
	ID                int     `json:"id"`
	Name              string  `json:"name"`
	Temperature       float64 `json:"temperature"`
	Humidity          float64 `json:"humidity"`
	IsFrontDoorLocked bool    `json:"is_front_door_locked"`
	IsCameraActive    bool    `json:"is_camera_active"`
	IsAlarmActive     bool    `json:"is_alarm_active"`
}
