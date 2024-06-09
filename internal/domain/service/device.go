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

func (s *DeviceService) GetById(ctx context.Context, id int) (*models.Device, error) {
	device, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return device, nil
}
