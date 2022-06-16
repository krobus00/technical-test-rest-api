package session

import (
	"context"

	"github.com/krobus00/technical-test-rest-api/model/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *repository) Store(ctx context.Context, db *mongo.Database, input *database.Session) (*database.Session, error) {
	var id primitive.ObjectID

	result, err := db.Collection(r.GetCollectionName()).InsertOne(ctx, input)
	if err != nil {
		r.logger.Zap.Error(err.Error())
		return nil, err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		id = oid
	}

	return &database.Session{
		ID:           id,
		Username:     input.Username,
		RefreshToken: input.RefreshToken,
		IsBlocked:    input.IsBlocked,
	}, nil
}

func (r *repository) DeleteSessionByRefreshToken(ctx context.Context, db *mongo.Database, input *database.Session) (int64, error) {
	filter := bson.M{"refresh_token": input.RefreshToken}

	result, err := db.Collection(r.GetCollectionName()).DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}
