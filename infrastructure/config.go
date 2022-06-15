package infrastructure

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Env struct {
	AppName     string
	AppPort     string
	Environment string

	MongoDSN      string
	MongoDatabase string

	AccessTokenSecret    string
	AccessTokenDuration  time.Duration
	RefreshTokenSecret   string
	RefreshTokenDuration time.Duration
}

func NewEnv() Env {
	env := Env{}
	env.LoadEnv()
	return env
}

func (env *Env) LoadEnv() error {
	_ = godotenv.Load()

	env.AppName = os.Getenv("APP_NAME")
	env.AppPort = os.Getenv("APP_PORT")
	env.Environment = os.Getenv("ENVIRONMENT")

	env.MongoDSN = os.Getenv("MONGO_DSN")
	env.MongoDatabase = os.Getenv("MONGO_DATABASE")

	env.AccessTokenSecret = os.Getenv("ACCESS_TOKEN_SECRET")
	accessTokenDurationInSecond := os.Getenv("ACCESS_TOKEN_DURATION_IN_SECOND")
	accessTokenDuration, err := strconv.Atoi(accessTokenDurationInSecond)
	if err != nil {
		return err
	}
	env.AccessTokenDuration = time.Duration(accessTokenDuration * int(time.Second))

	env.RefreshTokenSecret = os.Getenv("REFRESH_TOKEN_SECRET")
	refreshTokenDurationInSecond := os.Getenv("REFRESH_TOKEN_DURATION_IN_SECOND")
	refreshTokenDuration, err := strconv.Atoi(refreshTokenDurationInSecond)
	if err != nil {
		return err
	}
	env.RefreshTokenDuration = time.Duration(refreshTokenDuration * int(time.Second))

	return nil
}
