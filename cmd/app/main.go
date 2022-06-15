package main

import (
	"github.com/joho/godotenv"
	"github.com/krobus00/technical-test-rest-api/bootstrap"
	"go.uber.org/fx"
)

func main() {
	godotenv.Load()
	fx.New(bootstrap.AppModule).Run()
}
