package product

import (
	"context"

	"github.com/krobus00/technical-test-rest-api/infrastructure"
	"github.com/krobus00/technical-test-rest-api/model"
	"github.com/krobus00/technical-test-rest-api/model/database"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	tag = `[ProductRepository]`

	tracingStore           = "StoreProduct"
	tracingCount           = "CountProduct"
	tracingFindAll         = "FindAllProduct"
	tracingFindProductByID = "FindProductByID"
	tracingUpdate          = "UpdateProduct"
	tracingDelete          = "DeleteProduct"
)

type (
	ProductRepository interface {
		GetCollectionName() string
		Store(ctx context.Context, db *mongo.Database, input *database.Product) (*database.Product, error)
		Count(ctx context.Context, db *mongo.Database, filter *model.ProductFilter, paginationFilter *model.PaginationRequest) (int64, error)
		FindAll(ctx context.Context, db *mongo.Database, filter *model.ProductFilter, paginationFilter *model.PaginationRequest) ([]*database.Product, error)
		FindProductByID(ctx context.Context, db *mongo.Database, filter *model.ProductFilter, input *database.Product) (*database.Product, error)
		Update(ctx context.Context, db *mongo.Database, filter *model.ProductFilter, input *database.Product) (*database.Product, error)
		Delete(ctx context.Context, db *mongo.Database, filter *model.ProductFilter, input *database.Product) error
	}
	repository struct {
		logger infrastructure.Logger
	}
)

func New(infrastructure infrastructure.Infrastructure) ProductRepository {
	return &repository{
		logger: infrastructure.Logger,
	}
}

func (r *repository) GetCollectionName() string {
	return "products"
}
