package handler

type DeviceActivityRequest struct {
	ID     int  `json:"id"`
	Active bool `json:"active"`
}

type RoomStatusRequest struct {
	ID           int  `json:"id"`
	DoorStatus   bool `json:"door_status"`
	CameraStatus bool `json:"camera_status"`
	AlarmStatus  bool `json:"alarm_status"`
}

type EnergySavingResponse struct {
	TotalEnergySaving float64 `json:"total_energy_saving"`
}
