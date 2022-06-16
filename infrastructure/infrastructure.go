package infrastructure

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type Infrastructure struct {
	fx.In

	Logger     Logger
	Translator *ut.UniversalTranslator
	Router     *echo.Echo
	Env        Env
	Database   Database
}

func NewInfrastructure() *Infrastructure {
	return &Infrastructure{}
}

var Module = fx.Options(
	fx.Provide(NewEnv),
	fx.Provide(NewLogger),
	fx.Provide(NewTranslator),
	fx.Provide(NewRouter),
	fx.Provide(NewDatabase),
	fx.Provide(NewValidator),

	fx.Populate(NewInfrastructure()),
)
