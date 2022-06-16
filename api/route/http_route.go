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
	decodeAccessTokenJWTMiddleware := local_middleware.DecodeJWTTokenMiddleware(config.AccessTokenSecret)
	decodeRefreshTokenMiddleware := local_middleware.DecodeJWTTokenMiddleware(config.RefreshTokenSecret)

	service.POST("/user/register", handler.UserController.HandleUserRegister)
	service.POST("/user/login", handler.UserController.HandleUserLogin)
	service.GET("/user/me", handler.UserController.HandleGetUserInfo, decodeAccessTokenJWTMiddleware)
	service.GET("/user/refresh-token", handler.UserController.HandleRefreshToken, decodeRefreshTokenMiddleware)

	service.POST("/products", handler.ProductController.HandleCreateProduct, decodeAccessTokenJWTMiddleware)
	service.GET("/products", handler.ProductController.HandleGetAllProduct, decodeAccessTokenJWTMiddleware)
	service.GET("/products/:id", handler.ProductController.HandleGetProductDetail, decodeAccessTokenJWTMiddleware)
	service.PATCH("/products/:id", handler.ProductController.HandleUpdateProduct, decodeAccessTokenJWTMiddleware)
	service.DELETE("/products/:id", handler.ProductController.HandleDeleteProduct, decodeAccessTokenJWTMiddleware)
}

var Module = fx.Options(
	fx.Invoke(NewRoutes),
)
