package bootstrap

import (
	"context"
	"fmt"
	"net/http"

	"github.com/common-nighthawk/go-figure"
	"github.com/krobus00/technical-test-rest-api/api/controller"
	"github.com/krobus00/technical-test-rest-api/api/repository"
	"github.com/krobus00/technical-test-rest-api/api/route"
	"github.com/krobus00/technical-test-rest-api/api/service"
	"github.com/krobus00/technical-test-rest-api/infrastructure"
	"github.com/labstack/echo/v4"

	"go.uber.org/fx"
)

var AppModule = fx.Options(
	infrastructure.Module,
	repository.Module,
	service.Module,
	controller.Module,
	route.Module,
	fx.Invoke(appBootstrap),
)

func appBootstrap(
	lifecycle fx.Lifecycle,
	handler *echo.Echo,
	env infrastructure.Env,
	logger infrastructure.Logger,
	database infrastructure.Database,
) {

	appStop := func(ctx context.Context) error {
		logger.Zap.Info("Stopping Application")
		conn := database.NoSQL
		conn.Client().Disconnect(ctx)
		return nil
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Zap.Info("Starting Application")
			figure.NewColorFigure(env.AppName, "", "purple", true).Print()
			go func() {
				err := database.NoSQL.Client().Ping(ctx, nil)
				if err != nil {
					logger.Zap.Panic(err)
				} else {
					logger.Zap.Info("Database connected")
				}
				PORT := "5000"
				if env.AppPort != "" {
					PORT = env.AppPort
				}
				logger.Zap.Info(fmt.Sprintf("APP RUNNING ON http://0.0.0.0:%s", PORT))
				if err := handler.Start(fmt.Sprintf(":%s", PORT)); err != nil && err != http.ErrServerClosed {
					handler.Logger.Fatal("shutting down the server")
				}
			}()
			return nil
		},
		OnStop: appStop,
	})
}
