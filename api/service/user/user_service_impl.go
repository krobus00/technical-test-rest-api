package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/krobus00/technical-test-rest-api/constant"
	"github.com/krobus00/technical-test-rest-api/model"
	"github.com/krobus00/technical-test-rest-api/model/database"
	"github.com/krobus00/technical-test-rest-api/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (svc *service) RegisterUser(ctx context.Context, payload *model.RegisterUserRequest) (*model.RegisterUserResponse, error) {

	password, err := util.HashPassword(payload.Password)
	if err != nil {
		return nil, err
	}
	newUser := &database.User{
		Username: payload.Username,
		Password: password,
		Role:     constant.DEFAULT_ROLE,
		DateColumn: database.DateColumn{
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
			DeletedAt: nil,
		},
	}

	user, err := svc.repository.UserRepository.Store(ctx, svc.db, newUser)
	if err != nil {
		return nil, err
	}

	accessToken, err := util.CreateToken(user.ID.Hex(), svc.config.AccessTokenDuration, svc.config.AccessTokenSecret)
	refreshToken, err := util.CreateToken(user.ID.Hex(), svc.config.RefreshTokenDuration, svc.config.RefreshTokenSecret)

	newSession := &database.Session{
		Username:     user.Username,
		RefreshToken: refreshToken.Token,
		IsBlocked:    false,
		DateColumn: database.DateColumn{
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
			DeletedAt: nil,
		},
	}
	_, err = svc.repository.SessionRepository.Store(ctx, svc.db, newSession)
	if err != nil {
		return nil, err
	}

	return &model.RegisterUserResponse{
		AccessToken:           accessToken.Token,
		AccessTokenExpiredAt:  accessToken.Exp,
		RefreshToken:          refreshToken.Token,
		RefreshTokenExpiredAt: refreshToken.Exp,
	}, nil
}

func (svc *service) LoginUser(ctx context.Context, payload *model.LoginUserRequest) (*model.LoginUserResponse, error) {

	checkUser := &database.User{
		Username: payload.Username,
	}
	user, err := svc.repository.UserRepository.FindUserByUsername(ctx, svc.db, checkUser)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("username not found")
	}
	if !util.CheckPasswordHash(payload.Password, user.Password) {
		return nil, errors.New("Wrong password")
	}

	accessToken, err := util.CreateToken(user.ID.Hex(), svc.config.AccessTokenDuration, svc.config.AccessTokenSecret)
	refreshToken, err := util.CreateToken(user.ID.Hex(), svc.config.RefreshTokenDuration, svc.config.RefreshTokenSecret)

	newSession := &database.Session{
		Username:     user.Username,
		RefreshToken: refreshToken.Token,
		IsBlocked:    false,
		DateColumn: database.DateColumn{
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
			DeletedAt: nil,
		},
	}
	_, err = svc.repository.SessionRepository.Store(ctx, svc.db, newSession)
	if err != nil {
		return nil, err
	}

	return &model.LoginUserResponse{
		AccessToken:           accessToken.Token,
		AccessTokenExpiredAt:  accessToken.Exp,
		RefreshToken:          refreshToken.Token,
		RefreshTokenExpiredAt: refreshToken.Exp,
	}, nil
}

func (svc *service) GetUserInfo(ctx context.Context) (*model.UserResponse, error) {
	userID, err := primitive.ObjectIDFromHex(ctx.Value("userID").(string))
	if err != nil {
		return nil, err
	}
	checkUser := &database.User{
		ID: userID,
	}
	user, err := svc.repository.UserRepository.FindUserByID(ctx, svc.db, checkUser)
	if err != nil {
		return nil, err
	}
	fmt.Println()
	fmt.Println(user.DateColumn)
	fmt.Println()

	return &model.UserResponse{
		ID:       user.ID.Hex(),
		Username: user.Username,
		Role:     user.Role,
		DateColumn: model.DateColumn{
			CreatedAt: user.DateColumn.CreatedAt,
			UpdatedAt: user.DateColumn.UpdatedAt,
		},
	}, nil
}
