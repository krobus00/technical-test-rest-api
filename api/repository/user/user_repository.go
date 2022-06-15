package user

import (
	"context"

	"github.com/krobus00/technical-test-rest-api/infrastructure"
	"github.com/krobus00/technical-test-rest-api/model/database"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	tag = `[UserRepository]`

	tracingStore = "StoreUser"
)

type (
	UserRepository interface {
		GetCollectionName() string
		Store(ctx context.Context, db *mongo.Database, input *database.User) (*database.User, error)
		FindUserByUsername(ctx context.Context, db *mongo.Database, input *database.User) (*database.User, error)
		FindUserByID(ctx context.Context, db *mongo.Database, input *database.User) (*database.User, error)
	}
	repository struct {
		logger infrastructure.Logger
	}
)

func New(infrastructure infrastructure.Infrastructure) UserRepository {
	return &repository{
		logger: infrastructure.Logger,
	}
}

func (r *repository) GetCollectionName() string {
	return "users"
}
