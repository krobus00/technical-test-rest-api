package user

import (
	"context"
	"fmt"
	"net/http"

	ut "github.com/go-playground/universal-translator"
	"github.com/krobus00/technical-test-rest-api/api/service/user"
	"github.com/krobus00/technical-test-rest-api/infrastructure"
	"github.com/krobus00/technical-test-rest-api/model"
	"github.com/krobus00/technical-test-rest-api/util"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

const (
	tag = "[UserController]"

	tracingHandleUserRegister = "HandleUserRegister"
	tracingHandleUserLogin    = "HandleUserLogin"
	tracingHandleGetUserInfo  = "HandleGetUserInfo"
	tracingHandleRefreshToken = "HandleRefreshToken"
)

type Controller struct {
	fx.In

	Logger      infrastructure.Logger
	UserService user.UserService
	Translator  *ut.UniversalTranslator
}

func (c *Controller) HandleUserRegister(eCtx echo.Context) error {
	ctx := eCtx.Request().Context()

	payload := new(model.RegisterUserRequest)
	if err := eCtx.Bind(payload); err != nil {
		c.Logger.Zap.Error(fmt.Sprintf("%s %s with error: %v", tag, tracingHandleUserRegister, err))
		return err
	}

	if err := eCtx.Validate(payload); err != nil {
		trans := util.TranslatorFromRequestHeader(eCtx, c.Translator)
		return echo.NewHTTPError(http.StatusBadRequest, util.BuildValidationErrors(err, trans))
	}

	resp, err := c.UserService.RegisterUser(ctx, payload)
	if err != nil {
		c.Logger.Zap.Error(fmt.Sprintf("%s %s with error: %v", tag, tracingHandleUserRegister, err))
		return err
	}

	return eCtx.JSON(http.StatusOK, resp)
}

func (c *Controller) HandleUserLogin(eCtx echo.Context) error {
	ctx := eCtx.Request().Context()

	payload := new(model.LoginUserRequest)
	if err := eCtx.Bind(payload); err != nil {
		c.Logger.Zap.Error(fmt.Sprintf("%s %s with error: %v", tag, tracingHandleUserLogin, err))
		return err
	}

	if err := eCtx.Validate(payload); err != nil {
		trans := util.TranslatorFromRequestHeader(eCtx, c.Translator)
		return echo.NewHTTPError(http.StatusBadRequest, util.BuildValidationErrors(err, trans))
	}

	resp, err := c.UserService.LoginUser(ctx, payload)
	if err != nil {
		c.Logger.Zap.Error(fmt.Sprintf("%s %s with error: %v", tag, tracingHandleUserLogin, err))
		return err
	}

	return eCtx.JSON(http.StatusOK, resp)
}

func (c *Controller) HandleGetUserInfo(eCtx echo.Context) error {
	ctx := eCtx.Request().Context()

	ctx = context.WithValue(ctx, "userID", eCtx.Get("userID").(string))
	resp, err := c.UserService.GetUserInfo(ctx)
	if err != nil {
		c.Logger.Zap.Error(fmt.Sprintf("%s %s with error: %v", tag, tracingHandleGetUserInfo, err))
		return err
	}

	return eCtx.JSON(http.StatusOK, resp)
}

func (c *Controller) HandleRefreshToken(eCtx echo.Context) error {
	ctx := eCtx.Request().Context()

	ctx = context.WithValue(ctx, "userID", eCtx.Get("userID").(string))
	ctx = context.WithValue(ctx, "token", eCtx.Get("token").(string))
	resp, err := c.UserService.RefreshToken(ctx)
	if err != nil {
		c.Logger.Zap.Error(fmt.Sprintf("%s %s with error: %v", tag, tracingHandleRefreshToken, err))
		return err
	}

	return eCtx.JSON(http.StatusOK, resp)
}
