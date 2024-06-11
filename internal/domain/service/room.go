package service

import (
	"context"
	"time"

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

func (s *RoomService) GetById(ctx context.Context, id int) (*models.Room, error) {
	room, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (s *RoomService) UpdateRoomStatus(ctx context.Context, id int, frontDoorStatus, cameraStatus, alarmStatus bool) error {
	err := s.repository.UpdateRoomStatus(ctx, id, frontDoorStatus, cameraStatus, alarmStatus)
	if err != nil {
		return err
	}
	return nil
}

func (s *RoomService) GetRoomEnergySaving(ctx context.Context, id int) (float64, error) {
	devices, err := s.repository.GetRoomInActiveDevices(ctx, id)
	if err != nil {
		return 0, err
	}
	now := time.Now()
	var totalEnergySaving float64
	for _, device := range devices {
		duration := now.Sub(*device.UpdatedAT)
		totalEnergySaving += duration.Minutes() * device.EnergyConsumingPerHour / 60
	}
	return totalEnergySaving, nil
}
