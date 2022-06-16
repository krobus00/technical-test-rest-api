package user

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/krobus00/technical-test-rest-api/constant"
	"github.com/krobus00/technical-test-rest-api/model"
	"github.com/krobus00/technical-test-rest-api/model/database"
	"github.com/krobus00/technical-test-rest-api/util"
	"github.com/microcosm-cc/bluemonday"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (svc *service) RegisterUser(ctx context.Context, payload *model.RegisterUserRequest) (*model.TokenResponse, error) {

	password, err := util.HashPassword(payload.Password)
	if err != nil {
		svc.logger.Zap.Error(fmt.Sprintf("%s %s with error: %v", tag, tracingRegisterUser, err))
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

	user, err := svc.repository.UserRepository.FindUserByUsername(ctx, svc.db, newUser)
	if err != nil {
		svc.logger.Zap.Error(fmt.Sprintf("%s %s with error: %v", tag, tracingRegisterUser, err))
		return nil, err
	}
	if user != nil {
		return nil, model.NewHttpCustomError(http.StatusBadRequest, errors.New("Username already taken"))
	}

	user, err = svc.repository.UserRepository.Store(ctx, svc.db, newUser)
	if err != nil {
		svc.logger.Zap.Error(fmt.Sprintf("%s %s with error: %v", tag, tracingRegisterUser, err))
		return nil, err
	}

	accessToken, err := util.CreateToken(user.ID.Hex(), newUser.Role, svc.config.AccessTokenDuration, svc.config.AccessTokenSecret)
	if err != nil {
		svc.logger.Zap.Error(fmt.Sprintf("%s %s with error: %v", tag, tracingRegisterUser, err))
		return nil, err
	}
	refreshToken, err := util.CreateToken(user.ID.Hex(), newUser.Role, svc.config.RefreshTokenDuration, svc.config.RefreshTokenSecret)
	if err != nil {
		svc.logger.Zap.Error(fmt.Sprintf("%s %s with error: %v", tag, tracingRegisterUser, err))
		return nil, err
	}

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
		svc.logger.Zap.Error(fmt.Sprintf("%s %s with error: %v", tag, tracingRegisterUser, err))
		return nil, err
	}

	return &model.TokenResponse{
		AccessToken:           accessToken.Token,
		AccessTokenExpiredAt:  accessToken.Exp,
		RefreshToken:          refreshToken.Token,
		RefreshTokenExpiredAt: refreshToken.Exp,
	}, nil
}

func (svc *service) LoginUser(ctx context.Context, payload *model.LoginUserRequest) (*model.TokenResponse, error) {

	checkUser := &database.User{
		Username: payload.Username,
	}
	user, err := svc.repository.UserRepository.FindUserByUsername(ctx, svc.db, checkUser)
	if err != nil {
		svc.logger.Zap.Error(fmt.Sprintf("%s %s with error: %v", tag, tracingLoginUser, err))
		return nil, err
	}
	if user == nil {
		return nil, model.NewHttpCustomError(http.StatusBadRequest, errors.New("Username not found"))
	}
	if !util.CheckPasswordHash(payload.Password, user.Password) {
		return nil, model.NewHttpCustomError(http.StatusBadRequest, errors.New("Wrong password"))
	}

	accessToken, err := util.CreateToken(user.ID.Hex(), user.Role, svc.config.AccessTokenDuration, svc.config.AccessTokenSecret)
	if err != nil {
		svc.logger.Zap.Error(fmt.Sprintf("%s %s with error: %v", tag, tracingLoginUser, err))
		return nil, err
	}
	refreshToken, err := util.CreateToken(user.ID.Hex(), user.Role, svc.config.RefreshTokenDuration, svc.config.RefreshTokenSecret)
	if err != nil {
		svc.logger.Zap.Error(fmt.Sprintf("%s %s with error: %v", tag, tracingLoginUser, err))
		return nil, err
	}

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
		svc.logger.Zap.Error(fmt.Sprintf("%s %s with error: %v", tag, tracingLoginUser, err))
		return nil, err
	}

	return &model.TokenResponse{
		AccessToken:           accessToken.Token,
		AccessTokenExpiredAt:  accessToken.Exp,
		RefreshToken:          refreshToken.Token,
		RefreshTokenExpiredAt: refreshToken.Exp,
	}, nil
}

func (svc *service) GetUserInfo(ctx context.Context) (*model.UserResponse, error) {
	p := bluemonday.StrictPolicy()
	userID, err := primitive.ObjectIDFromHex(ctx.Value("userID").(string))
	if err != nil {
		svc.logger.Zap.Error(fmt.Sprintf("%s %s with error: %v", tag, tracingGetUserInfo, err))
		return nil, err
	}
	checkUser := &database.User{
		ID: userID,
	}
	user, err := svc.repository.UserRepository.FindUserByID(ctx, svc.db, checkUser)
	if err != nil {
		svc.logger.Zap.Error(fmt.Sprintf("%s %s with error: %v", tag, tracingGetUserInfo, err))
		return nil, err
	}

	return &model.UserResponse{
		ID:       user.ID.Hex(),
		Username: p.Sanitize(user.Username),
		Role:     user.Role,
		DateColumn: model.DateColumn{
			CreatedAt: user.DateColumn.CreatedAt,
			UpdatedAt: user.DateColumn.UpdatedAt,
		},
	}, nil
}

func (svc *service) RefreshToken(ctx context.Context) (*model.TokenResponse, error) {
	userID, err := primitive.ObjectIDFromHex(ctx.Value("userID").(string))
	if err != nil {
		svc.logger.Zap.Error(fmt.Sprintf("%s %s with error: %v", tag, tracingRefreshToken, err))
		return nil, err
	}
	checkUser := &database.User{
		ID: userID,
	}
	user, err := svc.repository.UserRepository.FindUserByID(ctx, svc.db, checkUser)
	if err != nil {
		svc.logger.Zap.Error(fmt.Sprintf("%s %s with error: %v", tag, tracingRefreshToken, err))
		return nil, err
	}

	session := &database.Session{
		RefreshToken: ctx.Value("token").(string),
	}

	removed, err := svc.repository.SessionRepository.DeleteSessionByRefreshToken(ctx, svc.db, session)
	if err != nil {
		svc.logger.Zap.Error(fmt.Sprintf("%s %s with error: %v", tag, tracingRefreshToken, err))
		return nil, err
	}
	if removed != 1 {
		return nil, model.NewHttpCustomError(http.StatusUnprocessableEntity, errors.New("Invalid refresh token"))
	}

	accessToken, err := util.CreateToken(user.ID.Hex(), user.Role, svc.config.AccessTokenDuration, svc.config.AccessTokenSecret)
	if err != nil {
		svc.logger.Zap.Error(fmt.Sprintf("%s %s with error: %v", tag, tracingRefreshToken, err))
		return nil, err
	}
	refreshToken, err := util.CreateToken(user.ID.Hex(), user.Role, svc.config.RefreshTokenDuration, svc.config.RefreshTokenSecret)
	if err != nil {
		svc.logger.Zap.Error(fmt.Sprintf("%s %s with error: %v", tag, tracingRefreshToken, err))
		return nil, err
	}

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
		svc.logger.Zap.Error(fmt.Sprintf("%s %s with error: %v", tag, tracingRefreshToken, err))
		return nil, err
	}
	return &model.TokenResponse{
		AccessToken:           accessToken.Token,
		AccessTokenExpiredAt:  accessToken.Exp,
		RefreshToken:          refreshToken.Token,
		RefreshTokenExpiredAt: refreshToken.Exp,
	}, nil
}
