package repository

import (
	"context"
	"driver-service/internal/model"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DriverRepository interface {
	Create(ctx context.Context, driver *model.Driver) error
	FindByID(ctx context.Context, id string) (*model.Driver, error)
	FindAll(ctx context.Context, page, pageSize int) ([]*model.Driver, int64, error)
	FindByTaxiType(ctx context.Context, taksiType string) ([]*model.Driver, error)
	Update(ctx context.Context, id string, updates bson.M) error
	Delete(ctx context.Context, id string) error
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

func (r *driverRepository) FindByID(ctx context.Context, id string) (*model.Driver, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id format: %w", err)
	}

	var driver model.Driver
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&driver)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("driver not found")
		}
		return nil, err
	}

	return &driver, nil
}

func (r *driverRepository) Update(ctx context.Context, id string, updates bson.M) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid id format: %w", err)
	}

	updates["updatedAt"] = time.Now()

	result, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": objectID},
		bson.M{"$set": updates},
	)

	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("driver not found")
	}

	return nil
}

func (r *driverRepository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid id format: %w", err)
	}

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("driver not found")
	}

	return nil
}

func (r *driverRepository) FindAll(ctx context.Context, page, pageSize int) ([]*model.Driver, int64, error) {
	skip := (page - 1) * pageSize

	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(pageSize))
	findOptions.SetSort(bson.D{{Key: "createdAt", Value: -1}})

	cursor, err := r.collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var drivers []*model.Driver
	if err = cursor.All(ctx, &drivers); err != nil {
		return nil, 0, err
	}

	total, err := r.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}

	return drivers, total, nil
}

func (r *driverRepository) FindByTaxiType(ctx context.Context, taksiType string) ([]*model.Driver, error) {
	filter := bson.M{}
	if taksiType != "" {
		filter["taksiType"] = taksiType
	}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var drivers []*model.Driver
	if err = cursor.All(ctx, &drivers); err != nil {
		return nil, err
	}

	return drivers, nil
}
