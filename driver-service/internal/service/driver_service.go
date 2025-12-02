package service

import (
	"context"
	"driver-service/internal/dto"
	"driver-service/internal/model"
	"driver-service/internal/repository"
	"fmt"
	"strings"
)

type DriverService interface {
	CreateDriver(ctx context.Context, req *dto.CreateDriverRequest) (*dto.DriverResponse, error)
}

type driverService struct {
	repo repository.DriverRepository
}

func NewDriverService(repo repository.DriverRepository) DriverService {
	return &driverService{
		repo: repo,
	}
}

func (s *driverService) CreateDriver(ctx context.Context, req *dto.CreateDriverRequest) (*dto.DriverResponse, error) {
	if err := s.validateDriver(req); err != nil {
		return nil, err
	}

	driver := &model.Driver{
		FirstName: strings.TrimSpace(req.FirstName),
		LastName:  strings.TrimSpace(req.LastName),
		Plate:     strings.ToUpper(strings.TrimSpace(req.Plate)),
		TaxiType:  req.TaxiType,
		CarBrand:  strings.TrimSpace(req.CarBrand),
		CarModel:  strings.TrimSpace(req.CarModel),
		Location: struct {
			Lat float64 `bson:"lat"`
			Lon float64 `bson:"lon"`
		}{
			Lat: req.Lat,
			Lon: req.Lon,
		},
	}

	if err := s.repo.Create(ctx, driver); err != nil {
		return nil, fmt.Errorf("failed to create driver: %w", err)
	}

	response := &dto.DriverResponse{
		ID: driver.ID.Hex(),
	}

	return response, nil
}

func (s *driverService) validateDriver(req *dto.CreateDriverRequest) error {
	return nil
}
