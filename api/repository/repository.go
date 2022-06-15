package repository

import (
	"github.com/krobus00/technical-test-rest-api/api/repository/session"
	"github.com/krobus00/technical-test-rest-api/api/repository/user"
	"go.uber.org/fx"
)

type Repository struct {
	fx.In

	UserRepository    user.UserRepository
	SessionRepository session.SessionRepository
}

var Module = fx.Options(
	fx.Provide(user.New),
	fx.Provide(session.New),
)
