package controller

import (
	"github.com/krobus00/technical-test-rest-api/api/controller/product"
	"github.com/krobus00/technical-test-rest-api/api/controller/user"
	"go.uber.org/fx"
)

type Handler struct {
	fx.In

	UserController    user.Controller
	ProductController product.Controller
}

func NewHandler() *Handler {
	return &Handler{}
}

var Module = fx.Options(
	fx.Populate(NewHandler()),
)
