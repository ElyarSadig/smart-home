package models

import (
	"context"
)

type DeviceRepository interface {
	GetById(ctx context.Context, id int) (*Device, error)
	UpdateDeviceActivity(ctx context.Context, id int, activity bool) error
}
