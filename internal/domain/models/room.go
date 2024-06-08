package models

type Room struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}
