package service

import (
	"context"

	"github.com/elyarsadig/smart-home-iot/internal/domain/models"
)

type RoomService struct {
	repository models.RoomRepository
}

func NewRoom(repo models.RoomRepository) *RoomService {
	return &RoomService{
		repository: repo,
	}
}

func (s *RoomService) GetByName(ctx context.Context, name string) (*models.Room, error) {
	room, err := s.repository.GetByName(ctx, name)
	if err != nil {
		return nil, err
	}
	return room, nil
}
