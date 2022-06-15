package session

import (
	"context"

	"github.com/krobus00/technical-test-rest-api/model/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *repository) Store(ctx context.Context, db *mongo.Database, input *database.Session) (*database.Session, error) {
	var id primitive.ObjectID

	result, err := db.Collection(r.GetCollectionName()).InsertOne(ctx, input)
	if err != nil {
		r.logger.Zap.Error(err.Error())
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
