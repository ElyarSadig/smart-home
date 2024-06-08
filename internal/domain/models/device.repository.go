package models

import (
	"context"
)

type DeviceRepository interface {
	GetByName(ctx context.Context, name string) (*Device, error)
}
