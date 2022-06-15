package session

import (
	"context"

	"github.com/krobus00/technical-test-rest-api/infrastructure"
	"github.com/krobus00/technical-test-rest-api/model/database"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	tag = `[SessionRepository]`

	tracingStore = "StoreSession"
)

type (
	SessionRepository interface {
		GetCollectionName() string
		Store(ctx context.Context, db *mongo.Database, input *database.Session) (*database.Session, error)
		DeleteSessionByRefreshToken(ctx context.Context, db *mongo.Database, input *database.Session) (int64, error)
	}
	repository struct {
		logger infrastructure.Logger
	}
)

func New(infrastructure infrastructure.Infrastructure) SessionRepository {
	return &repository{
		logger: infrastructure.Logger,
	}
}

func (r *repository) GetCollectionName() string {
	return "sessions"
}
