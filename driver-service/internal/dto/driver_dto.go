package dto

type CreateDriverRequest struct {
	FirstName string  `json:"firstName" validate:"required"`
	LastName  string  `json:"lastName" validate:"required"`
	Plate     string  `json:"plate" validate:"required"`
	TaxiType  string  `json:"taksiType" validate:"required"`
	CarBrand  string  `json:"carBrand" validate:"required"`
	CarModel  string  `json:"carModel" validate:"required"`
	Lat       float64 `json:"lat"`
	Lon       float64 `json:"lon"`
}

type UpdateDriverRequest struct {
	FirstName *string  `json:"firstName,omitempty" binding:"omitempty,min=2,max=50"`
	LastName  *string  `json:"lastName,omitempty" binding:"omitempty,min=2,max=50"`
	Plate     *string  `json:"plate,omitempty" binding:"omitempty,len=8"`
	TaxiType  *string  `json:"taksiType,omitempty" binding:"omitempty,oneof=sari turuncu siyah"`
	CarBrand  *string  `json:"carBrand,omitempty" binding:"omitempty,min=2,max=50"`
	CarModel  *string  `json:"carModel,omitempty" binding:"omitempty,min=2,max=50"`
	Lat       *float64 `json:"lat,omitempty" binding:"omitempty,latitude"`
	Lon       *float64 `json:"lon,omitempty" binding:"omitempty,longitude"`
}

type DriverListResponse struct {
	Data       []*DriverResponse `json:"data"`
	Total      int64             `json:"total"`
	Page       int               `json:"page"`
	PageSize   int               `json:"pageSize"`
	TotalPages int               `json:"totalPages"`
}

type DriverResponse struct {
	ID        string  `json:"id"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Plate     string  `json:"plate"`
	TaxiType  string  `json:"taksiType"`
	CarBrand  string  `json:"carBrand"`
	CarModel  string  `json:"carModel"`
	Lat       float64 `json:"lat"`
	Lon       float64 `json:"lon"`
	CreatedAt string  `json:"createdAt"`
}

type NearbyDriverRequest struct {
	Lat      float64 `form:"lat" binding:"required,latitude"`
	Lon      float64 `form:"lon" binding:"required,longitude"`
	TaxiType string  `form:"taksiType" binding:"omitempty,oneof=sari turuncu siyah"`
	RadiusKm float64 `form:"radius" binding:"omitempty,min=0.1,max=50"`
}

type NearbyDriverResponse struct {
	ID         string  `json:"id"`
	FirstName  string  `json:"firstName"`
	LastName   string  `json:"lastName"`
	Plate      string  `json:"plate"`
	TaxiType   string  `json:"taksiType"`
	CarBrand   string  `json:"carBrand"`
	CarModel   string  `json:"carModel"`
	DistanceKm float64 `json:"distanceKm"`
	Lat        float64 `json:"lat"`
	Lon        float64 `json:"lon"`
}
