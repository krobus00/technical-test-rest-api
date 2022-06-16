package user

import (
	"context"

	"github.com/krobus00/technical-test-rest-api/api/repository"
	"github.com/krobus00/technical-test-rest-api/infrastructure"
	"github.com/krobus00/technical-test-rest-api/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	tag = `[UserService]`

	tracingRegisterUser = "RegisterUser"
	tracingLoginUser    = "LoginUser"
	tracingGetUserInfo  = "GetUserInfo"
	tracingRefreshToken = "RefreshToken"
)

type (
	UserService interface {
		RegisterUser(ctx context.Context, payload *model.RegisterUserRequest) (*model.TokenResponse, error)
		LoginUser(ctx context.Context, payload *model.LoginUserRequest) (*model.TokenResponse, error)
		GetUserInfo(ctx context.Context) (*model.UserResponse, error)
		RefreshToken(ctx context.Context) (*model.TokenResponse, error)
	}
	service struct {
		logger     infrastructure.Logger
		db         *mongo.Database
		repository repository.Repository
		config     infrastructure.Env
	}
)

func New(
	infrastructure infrastructure.Infrastructure,
	repository repository.Repository,
) UserService {
	return &service{
		logger:     infrastructure.Logger,
		db:         infrastructure.Database.NoSQL,
		config:     infrastructure.Env,
		repository: repository,
	}
}
