package db

import (
	"context"
	"fmt"
	"github.com/Mortimor1/mikromon-core/internal/group"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	collection mongo.Collection
}

func (r *repository) Create(ctx context.Context, group *group.Group) (string, error) {
	result, err := r.collection.InsertOne(ctx, group)
	if err != nil {
		return "", err
	}
	oid := result.InsertedID.(primitive.ObjectID)
	return oid.Hex(), nil
}

func (r *repository) FindAll(ctx context.Context) ([]group.Group, error) {
	var g []group.Group

	cur, err := r.collection.Find(ctx, bson.D{{}})
	if err != nil {
		return g, err
	}

	if err = cur.All(ctx, &g); err != nil {
		return g, err
	}

	err = cur.Close(ctx)
	if err != nil {
		return g, err
	}

	return g, nil
}

func (r *repository) FindOne(ctx context.Context, id string) (g group.Group, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return g, err
	}
	query := bson.M{"_id": oid}
	result := r.collection.FindOne(ctx, query)
	if result.Err() != nil {
		return g, result.Err()
	}
	if err = result.Decode(&g); err != nil {
		return g, err
	}
	return g, nil
}

func (r *repository) Update(ctx context.Context, group *group.Group) error {
	oid, err := primitive.ObjectIDFromHex(group.Id)
	if err != nil {
		return err
	}

	bsonGroup, err := bson.Marshal(group)
	if err != nil {
		return err
	}

	result, err := r.collection.UpdateByID(ctx, oid, bsonGroup)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("not found")
	}
	return nil
}

func (r *repository) Delete(ctx context.Context, id string) error {
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

func NewRepository(collection mongo.Collection) *repository {
	return &repository{
		collection: collection,
	}
}
