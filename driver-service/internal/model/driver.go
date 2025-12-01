package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Driver struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `bson:"firstName"`
	LastName  string             `bson:"lastName"`
	Plate     string             `bson:"plate"`
	TaxiType  string             `bson:"taxiType"`
	CarBrand  string             `bson:"carBrand"`
	CarModel  string             `bson:"carModel"`
	Location  struct {
		Lat float64 `bson:"lat"`
		Lon float64 `bson:"lon"`
	} `bson:"location"`

	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}
