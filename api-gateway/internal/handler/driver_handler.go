package handler

import (
	"api-gateway/internal/proxy"

	"github.com/gin-gonic/gin"
)

type DriverHandler struct {
	proxy *proxy.DriverProxy
}

func NewDriverHandler(driverProxy *proxy.DriverProxy) *DriverHandler {
	return &DriverHandler{
		proxy: driverProxy,
	}
}

// CreateDriver godoc
// @Summary      Create a new driver
// @Description  Register a new taxi driver in the system
// @Tags         Drivers
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        driver body CreateDriverRequest true "Driver data"
// @Success      201 {object} DriverResponse
// @Failure      400 {object} ErrorResponse
// @Failure      401 {object} ErrorResponse
// @Failure      500 {object} ErrorResponse
// @Router       /api/v1/drivers [post]
func (h *DriverHandler) CreateDriver(c *gin.Context) {
	h.proxy.Forward(c)
}

// GetDrivers godoc
// @Summary      List all drivers
// @Description  Get a paginated list of all drivers
// @Tags         Drivers
// @Accept       json
// @Produce      json
// @Param        page query int false "Page number" default(1)
// @Param        pageSize query int false "Page size" default(20)
// @Success      200 {object} DriverListResponse
// @Failure      500 {object} ErrorResponse
// @Router       /api/v1/drivers [get]
func (h *DriverHandler) GetDrivers(c *gin.Context) {
	h.proxy.Forward(c)
}

// GetDriver godoc
// @Summary      Get driver by ID
// @Description  Get detailed information about a specific driver
// @Tags         Drivers
// @Accept       json
// @Produce      json
// @Param        id path string true "Driver ID"
// @Success      200 {object} SingleDriverResponse
// @Failure      404 {object} ErrorResponse
// @Router       /api/v1/drivers/{id} [get]
func (h *DriverHandler) GetDriver(c *gin.Context) {
	h.proxy.Forward(c)
}

// GetNearbyDrivers godoc
// @Summary      Get nearby drivers
// @Description  Find drivers within specified radius using GPS coordinates
// @Tags         Drivers
// @Accept       json
// @Produce      json
// @Param        lat query number true "Latitude" example(41.0082)
// @Param        lon query number true "Longitude" example(28.9784)
// @Param        taksiType query string false "Taxi type" Enums(sari, turuncu, siyah)
// @Param        radius query number false "Radius in km" default(6)
// @Success      200 {object} NearbyDriversResponse
// @Failure      400 {object} ErrorResponse
// @Failure      500 {object} ErrorResponse
// @Router       /api/v1/drivers/nearby [get]
func (h *DriverHandler) GetNearbyDrivers(c *gin.Context) {
	h.proxy.Forward(c)
}

// UpdateDriver godoc
// @Summary      Update driver
// @Description  Update driver information
// @Tags         Drivers
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Driver ID"
// @Param        driver body UpdateDriverRequest true "Updated driver data"
// @Success      200 {object} SuccessResponse
// @Failure      400 {object} ErrorResponse
// @Failure      401 {object} ErrorResponse
// @Failure      404 {object} ErrorResponse
// @Router       /api/v1/drivers/{id} [put]
func (h *DriverHandler) UpdateDriver(c *gin.Context) {
	h.proxy.Forward(c)
}

// DeleteDriver godoc
// @Summary      Delete driver
// @Description  Remove a driver from the system
// @Tags         Drivers
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Driver ID"
// @Success      200 {object} SuccessResponse
// @Failure      401 {object} ErrorResponse
// @Failure      404 {object} ErrorResponse
// @Router       /api/v1/drivers/{id} [delete]
func (h *DriverHandler) DeleteDriver(c *gin.Context) {
	h.proxy.Forward(c)
}

// Models
type CreateDriverRequest struct {
	FirstName string  `json:"firstName" example:"Ahmet" binding:"required"`
	LastName  string  `json:"lastName" example:"Yılmaz" binding:"required"`
	Plate     string  `json:"plate" example:"34ABC123" binding:"required"`
	TaxiType  string  `json:"taksiType" example:"sari" enums:"sari,turuncu,siyah" binding:"required"`
	CarBrand  string  `json:"carBrand" example:"Toyota" binding:"required"`
	CarModel  string  `json:"carModel" example:"Corolla" binding:"required"`
	Lat       float64 `json:"lat" example:"41.0082" binding:"required"`
	Lon       float64 `json:"lon" example:"28.9784" binding:"required"`
}

type UpdateDriverRequest struct {
	FirstName *string  `json:"firstName,omitempty" example:"Ahmet Ali"`
	LastName  *string  `json:"lastName,omitempty" example:"Yılmaz"`
	Lat       *float64 `json:"lat,omitempty" example:"41.0100"`
	Lon       *float64 `json:"lon,omitempty" example:"28.9800"`
}

type DriverInfo struct {
	ID        string  `json:"id" example:"675abc123def456"`
	FirstName string  `json:"firstName" example:"Ahmet"`
	LastName  string  `json:"lastName" example:"Yılmaz"`
	Plate     string  `json:"plate" example:"34ABC123"`
	TaxiType  string  `json:"taksiType" example:"sari"`
	CarBrand  string  `json:"carBrand" example:"Toyota"`
	CarModel  string  `json:"carModel" example:"Corolla"`
	Lat       float64 `json:"lat" example:"41.0082"`
	Lon       float64 `json:"lon" example:"28.9784"`
	CreatedAt string  `json:"createdAt" example:"2024-12-06T15:30:00Z"`
}

type DriverResponse struct {
	Message string     `json:"message" example:"Driver created successfully"`
	Data    DriverInfo `json:"data"`
}

type SingleDriverResponse struct {
	Data DriverInfo `json:"data"`
}

type DriverListResponse struct {
	Data       []DriverInfo `json:"data"`
	Total      int64        `json:"total" example:"42"`
	Page       int          `json:"page" example:"1"`
	PageSize   int          `json:"pageSize" example:"20"`
	TotalPages int          `json:"totalPages" example:"3"`
}

type NearbyDriverInfo struct {
	ID         string  `json:"id" example:"675abc123"`
	FirstName  string  `json:"firstName" example:"Mehmet"`
	LastName   string  `json:"lastName" example:"Kaya"`
	Plate      string  `json:"plate" example:"34XYZ789"`
	DistanceKm float64 `json:"distanceKm" example:"1.23"`
}

type NearbyDriversResponse struct {
	Count int                `json:"count" example:"5"`
	Data  []NearbyDriverInfo `json:"data"`
}

type SuccessResponse struct {
	Message string `json:"message" example:"Operation successful"`
}
