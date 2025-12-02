package repository

import (
	"context"
	"driver-service/internal/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DriverRepository interface {
	Create(ctx context.Context, driver *model.Driver) error
}

type driverRepository struct {
	collection *mongo.Collection
}

func NewDriverRepository(collection *mongo.Collection) DriverRepository {
	return &driverRepository{
		collection: collection,
	}
}

func (r *driverRepository) Create(ctx context.Context, driver *model.Driver) error {
	driver.ID = primitive.NewObjectID()
	driver.CreatedAt = time.Now()
	driver.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, driver)
	return err
}
