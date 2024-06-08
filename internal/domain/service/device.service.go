package service

import (
	"context"

	"github.com/elyarsadig/smart-home-iot/internal/domain/models"
)

type DeviceService struct {
	repository models.DeviceRepository
}

func NewDevice(repo models.DeviceRepository) *DeviceService {
	return &DeviceService{
		repository: repo,
	}
}

func (s *DeviceService) GetByName(ctx context.Context, name string) (*models.Device, error) {
	device, err := s.repository.GetByName(ctx, name)
	if err != nil {
		return nil, err
	}
	return device, nil
}
