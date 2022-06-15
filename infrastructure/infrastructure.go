package infrastructure

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type Infrastructure struct {
	fx.In

	Logger   Logger
	Router   *echo.Echo
	Env      Env
	Database Database
}

func NewInfrastructure() *Infrastructure {
	return &Infrastructure{}
}

var Module = fx.Options(
	fx.Provide(NewEnv),
	fx.Provide(NewLogger),
	fx.Provide(NewRouter),
	fx.Provide(NewDatabase),

	fx.Populate(NewInfrastructure()),
)
