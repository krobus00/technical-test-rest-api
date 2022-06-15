package route

import (
	"github.com/krobus00/technical-test-rest-api/api/controller"
	local_middleware "github.com/krobus00/technical-test-rest-api/api/middleware"
	"github.com/krobus00/technical-test-rest-api/infrastructure"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func NewRoutes(
	router *echo.Echo,
	handler controller.Handler,
	config infrastructure.Env,
) {
	e := router

	service := e.Group("/service")
	decodeJWTMiddleware := local_middleware.DecodeJWTTokenMiddleware(config.AccessTokenSecret)

	service.POST("/user/register", handler.UserController.HandleUserRegister)
	service.POST("/user/login", handler.UserController.HandleUserLogin)
	service.GET("/user/me", handler.UserController.HandleGetUserInfo, decodeJWTMiddleware)
}

var Module = fx.Options(
	fx.Invoke(NewRoutes),
)
