package user

import (
	"context"
	"fmt"

	"github.com/krobus00/technical-test-rest-api/model/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *repository) Store(ctx context.Context, db *mongo.Database, input *database.User) (*database.User, error) {
	var id primitive.ObjectID

	result, err := db.Collection(r.GetCollectionName()).InsertOne(ctx, input)
	if err != nil {
		r.logger.Zap.Error(fmt.Sprintf("%s %s with error: %v", tag, tracingStore, err))
		return nil, err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		id = oid
	}

	return &database.User{
		ID:       id,
		Username: input.Username,
		Password: "",
		Role:     input.Role,
	}, nil
}

func (r *repository) FindUserByUsername(ctx context.Context, db *mongo.Database, input *database.User) (*database.User, error) {
	filter := bson.M{"username": input.Username}

	result := new(database.User)
	err := db.Collection(r.GetCollectionName()).FindOne(ctx, filter).Decode(result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		r.logger.Zap.Error(fmt.Sprintf("%s %s with error: %v", tag, tracingFindUserByUsername, err))
		return nil, err
	}

	return result, nil
}

func (r *repository) FindUserByID(ctx context.Context, db *mongo.Database, input *database.User) (*database.User, error) {
	filter := bson.M{"_id": input.ID}

	result := new(database.User)
	err := db.Collection(r.GetCollectionName()).FindOne(ctx, filter).Decode(result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		r.logger.Zap.Error(fmt.Sprintf("%s %s with error: %v", tag, tracingFindUserByID, err))
		return nil, err
	}

	return result, nil
}
