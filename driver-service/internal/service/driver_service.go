package service

import (
	"context"
	"driver-service/internal/dto"
	"driver-service/internal/model"
	"driver-service/internal/repository"
	"fmt"
	"math"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

type DriverService interface {
	CreateDriver(ctx context.Context, req *dto.CreateDriverRequest) (*dto.DriverResponse, error)
	GetDriver(ctx context.Context, id string) (*dto.DriverResponse, error)
	GetDrivers(ctx context.Context, page, pageSize int) (*dto.DriverListResponse, error)
	UpdateDriver(ctx context.Context, id string, req *dto.UpdateDriverRequest) error
	DeleteDriver(ctx context.Context, id string) error
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

func (s *driverService) GetDriver(ctx context.Context, id string) (*dto.DriverResponse, error) {
	driver, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return s.modelToResponse(driver), nil
}

func (s *driverService) GetDrivers(ctx context.Context, page, pageSize int) (*dto.DriverListResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	drivers, total, err := s.repo.FindAll(ctx, page, pageSize)
	if err != nil {
		return nil, err
	}

	responseDrivers := make([]*dto.DriverResponse, len(drivers))
	for i, driver := range drivers {
		responseDrivers[i] = s.modelToResponse(driver)
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))

	return &dto.DriverListResponse{
		Data:       responseDrivers,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}, nil
}

func (s *driverService) UpdateDriver(ctx context.Context, id string, req *dto.UpdateDriverRequest) error {
	_, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	updates := bson.M{}

	if req.FirstName != nil {
		updates["firstName"] = strings.TrimSpace(*req.FirstName)
	}
	if req.LastName != nil {
		updates["lastName"] = strings.TrimSpace(*req.LastName)
	}
	if req.Plate != nil {
		plate := strings.ToUpper(strings.TrimSpace(*req.Plate))
		if len(plate) != 8 {
			return fmt.Errorf("plate must be exactly 8 characters")
		}
		updates["plate"] = plate
	}
	if req.TaxiType != nil {
		updates["taksiType"] = *req.TaxiType
	}
	if req.CarBrand != nil {
		updates["carBrand"] = strings.TrimSpace(*req.CarBrand)
	}
	if req.CarModel != nil {
		updates["carModel"] = strings.TrimSpace(*req.CarModel)
	}
	if req.Lat != nil {
		updates["location.lat"] = *req.Lat
	}
	if req.Lon != nil {
		updates["location.lon"] = *req.Lon
	}

	if len(updates) == 0 {
		return fmt.Errorf("no fields to update")
	}

	return s.repo.Update(ctx, id, updates)
}

func (s *driverService) DeleteDriver(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *driverService) modelToResponse(driver *model.Driver) *dto.DriverResponse {
	return &dto.DriverResponse{
		ID:        driver.ID.Hex(),
		FirstName: driver.FirstName,
		LastName:  driver.LastName,
		Plate:     driver.Plate,
		TaxiType:  driver.TaxiType,
		CarBrand:  driver.CarBrand,
		CarModel:  driver.CarModel,
		Lat:       driver.Location.Lat,
		Lon:       driver.Location.Lon,
		CreatedAt: driver.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}

func (s *driverService) validateDriver(req *dto.CreateDriverRequest) error {
	return nil
}
