package user

import (
	"context"
	"net/http"

	"github.com/krobus00/technical-test-rest-api/api/service/user"
	"github.com/krobus00/technical-test-rest-api/infrastructure"
	"github.com/krobus00/technical-test-rest-api/model"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

const (
	tag = "[UserController]"
)

type Controller struct {
	fx.In

	Logger      infrastructure.Logger
	UserService user.UserService
}

func (c *Controller) HandleUserRegister(eCtx echo.Context) error {
	ctx := eCtx.Request().Context()

	payload := new(model.RegisterUserRequest)
	if err := eCtx.Bind(payload); err != nil {
		return err
	}

	// if err := eCtx.Validate(payload); err != nil {
	// 	trans := kro_util.TranslatorFromRequestHeader(eCtx, c.Translator)
	// 	return echo.NewHTTPError(http.StatusBadRequest, kro_util.BuildValidationErrors(err, trans))
	// }
	resp, err := c.UserService.RegisterUser(ctx, payload)
	if err != nil {
		return err
	}

	return eCtx.JSON(http.StatusOK, resp)
}
func (c *Controller) HandleUserLogin(eCtx echo.Context) error {
	ctx := eCtx.Request().Context()

	payload := new(model.LoginUserRequest)
	if err := eCtx.Bind(payload); err != nil {
		return err
	}

	// if err := eCtx.Validate(payload); err != nil {
	// 	trans := kro_util.TranslatorFromRequestHeader(eCtx, c.Translator)
	// 	return echo.NewHTTPError(http.StatusBadRequest, kro_util.BuildValidationErrors(err, trans))
	// }
	resp, err := c.UserService.LoginUser(ctx, payload)
	if err != nil {
		return err
	}

	return eCtx.JSON(http.StatusOK, resp)
}

func (c *Controller) HandleGetUserInfo(eCtx echo.Context) error {
	ctx := eCtx.Request().Context()

	// if err := eCtx.Validate(payload); err != nil {
	// 	trans := kro_util.TranslatorFromRequestHeader(eCtx, c.Translator)
	// 	return echo.NewHTTPError(http.StatusBadRequest, kro_util.BuildValidationErrors(err, trans))
	// }
	ctx = context.WithValue(ctx, "userID", eCtx.Get("userID").(string))
	resp, err := c.UserService.GetUserInfo(ctx)
	if err != nil {
		return err
	}

	return eCtx.JSON(http.StatusOK, resp)
}
