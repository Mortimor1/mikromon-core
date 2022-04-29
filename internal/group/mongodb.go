package group

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type groupRepository struct {
	collection *mongo.Collection
}

func (r *groupRepository) Create(ctx context.Context, group *Group) (string, error) {
	result, err := r.collection.InsertOne(ctx, group)
	if err != nil {
		return "", err
	}
	oid := result.InsertedID.(primitive.ObjectID)
	return oid.Hex(), nil
}

func (r *groupRepository) FindAll(ctx context.Context) ([]Group, error) {
	var g []Group

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

	if len(g) == 0 {
		g = make([]Group, 0)
	}

	return g, nil
}

func (r *groupRepository) FindOne(ctx context.Context, id string) (g Group, err error) {
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

func (r *groupRepository) Update(ctx context.Context, group *Group) error {
	oid, err := primitive.ObjectIDFromHex(group.Id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": oid}

	groupBytes, _ := bson.Marshal(group)
	var updateGroup bson.M
	err = bson.Unmarshal(groupBytes, &updateGroup)
	if err != nil {
		return err
	}
	delete(updateGroup, "_id")

	update := bson.M{
		"$set": updateGroup,
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

func (r *groupRepository) Delete(ctx context.Context, id string) error {
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

func NewGroupRepository(collection *mongo.Collection) *groupRepository {
	return &groupRepository{
		collection: collection,
	}
}
