package product

import (
	"context"

	"github.com/krobus00/technical-test-rest-api/api/repository"
	"github.com/krobus00/technical-test-rest-api/infrastructure"
	"github.com/krobus00/technical-test-rest-api/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	tag = `[ProductService]`

	tracingStore           = "StoreProduct"
	tracingFindAll         = "FindAllProducts"
	tracingFindProductByID = "FindProductByID"
	tracingUpdate          = "UpdateProduct"
	tracingDelete          = "DeleteProduct"
)

type (
	ProductService interface {
		Store(ctx context.Context, payload *model.CreateProductRequest) (*model.ProductResponse, error)
		FindAll(ctx context.Context, payload *model.PaginationRequest) (*model.PaginationResponse, error)
		FindProductByID(ctx context.Context, payload *model.GetProductDetailRequest) (*model.ProductResponse, error)
		Update(ctx context.Context, payload *model.UpdateProductRequest) (*model.ProductResponse, error)
		Delete(ctx context.Context, payload *model.DeleteProductRequest) error
	}
	service struct {
		logger     infrastructure.Logger
		db         *mongo.Database
		repository repository.Repository
		config     infrastructure.Env
	}
)

func New(
	infrastructure infrastructure.Infrastructure,
	repository repository.Repository,
) ProductService {
	return &service{
		logger:     infrastructure.Logger,
		db:         infrastructure.Database.NoSQL,
		config:     infrastructure.Env,
		repository: repository,
	}
}
