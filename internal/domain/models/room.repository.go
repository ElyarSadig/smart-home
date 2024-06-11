package models

import "context"

type RoomRepository interface {
	GetById(ctx context.Context, id int) (*Room, error)
	UpdateRoomStatus(ctx context.Context, id int, frontDoorStatus, cameraStatus, alarmStatus bool) error
	GetRoomInActiveDevices(ctx context.Context, id int) ([]Device, error)
}
