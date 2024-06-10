package handler

type DeviceActivityRequest struct {
	ID     int  `json:"id"`
	Active bool `json:"active"`
}
