package infrastructure

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	Zap *zap.SugaredLogger
}

func NewLogger(env Env) Logger {

	var config = zap.Config{}
	if env.Environment != "local" {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	logger, _ := config.Build(zap.Hooks(func(entry zapcore.Entry) error {
		return nil
	}))

	sugar := logger.Sugar()

	return Logger{
		Zap: sugar,
	}

}
