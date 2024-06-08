package models

import "context"

type RoomRepository interface {
	GetByName(ctx context.Context, name string) (*Room, error)
}
