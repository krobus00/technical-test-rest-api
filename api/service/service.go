package service

import (
	"github.com/krobus00/technical-test-rest-api/api/service/product"
	"github.com/krobus00/technical-test-rest-api/api/service/user"
	"go.uber.org/fx"
)

type Service struct {
	fx.In

	UserService    user.UserService
	ProductService product.ProductService
}

var Module = fx.Options(
	fx.Provide(user.New),
	fx.Provide(product.New),
)
