package device

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type deviceRepository struct {
	collection *mongo.Collection
}

func (r *deviceRepository) Create(ctx context.Context, device *Device) (string, error) {
	result, err := r.collection.InsertOne(ctx, device)
	if err != nil {
		return "", err
	}
	oid := result.InsertedID.(primitive.ObjectID)
	return oid.Hex(), nil
}

func (r *deviceRepository) FindAll(ctx context.Context) ([]Device, error) {
	var d []Device

	cur, err := r.collection.Find(ctx, bson.D{{}})
	if err != nil {
		return d, err
	}

	if err = cur.All(ctx, &d); err != nil {
		return d, err
	}

	err = cur.Close(ctx)
	if err != nil {
		return d, err
	}

	if len(d) == 0 {
		d = make([]Device, 0)
	}

	return d, nil
}

func (r *deviceRepository) FindOne(ctx context.Context, id string) (d Device, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return d, err
	}
	query := bson.M{"_id": oid}
	result := r.collection.FindOne(ctx, query)
	if result.Err() != nil {
		return d, result.Err()
	}
	if err = result.Decode(&d); err != nil {
		return d, err
	}
	return d, nil
}

func (r *deviceRepository) Update(ctx context.Context, device *Device) error {
	oid, err := primitive.ObjectIDFromHex(device.Id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": oid}

	deviceBytes, _ := bson.Marshal(device)
	var updateDevice bson.M
	err = bson.Unmarshal(deviceBytes, &updateDevice)
	if err != nil {
		return err
	}
	delete(updateDevice, "_id")

	update := bson.M{
		"$set": updateDevice,
	}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("not found")
	}
	return nil
}

func (r *deviceRepository) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": oid}

	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("not found")
	}
	return nil
}

func NewDeviceRepository(collection *mongo.Collection) *deviceRepository {
	return &deviceRepository{
		collection: collection,
	}
}
