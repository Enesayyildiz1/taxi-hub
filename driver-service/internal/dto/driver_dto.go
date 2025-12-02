package dto

type CreateDriverRequest struct {
	FirstName string  `json:"firstName" validate:"required"`
	LastName  string  `json:"lastName" validate:"required"`
	Plate     string  `json:"plate" validate:"required"`
	TaxiType  string  `json:"taxiType" validate:"required"`
	CarBrand  string  `json:"carBrand" validate:"required"`
	CarModel  string  `json:"carModel" validate:"required"`
	Lat       float64 `json:"lat"`
	Lon       float64 `json:"lon"`
}

type DriverResponse struct {
	ID string `json:"id"`
}
