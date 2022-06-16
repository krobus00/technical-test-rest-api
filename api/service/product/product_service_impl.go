package product

import (
	"context"
	"time"

	"github.com/krobus00/technical-test-rest-api/constant"
	"github.com/krobus00/technical-test-rest-api/model"
	"github.com/krobus00/technical-test-rest-api/model/database"
	"github.com/microcosm-cc/bluemonday"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (svc *service) Store(ctx context.Context, payload *model.CreateProductRequest) (*model.ProductResponse, error) {
	p := bluemonday.StrictPolicy()
	userID, err := primitive.ObjectIDFromHex(ctx.Value("userID").(string))
	if err != nil {
		return nil, err
	}
	newProduct := &database.Product{
		UserID:      userID,
		Name:        payload.Name,
		Description: payload.Description,
		Price:       payload.Price,
		DateColumn: database.DateColumn{
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
			DeletedAt: nil,
		},
	}
	newProduct, err = svc.repository.ProductRepository.Store(ctx, svc.db, newProduct)
	if err != nil {
		return nil, err
	}
	return &model.ProductResponse{
		ID:          newProduct.ID.Hex(),
		UserID:      newProduct.ID.Hex(),
		Name:        p.Sanitize(newProduct.Name),
		Description: p.Sanitize(newProduct.Description),
		Price:       newProduct.Price,
		DateColumn: model.DateColumn{
			CreatedAt: newProduct.CreatedAt,
			UpdatedAt: newProduct.UpdatedAt,
			DeletedAt: newProduct.DeletedAt,
		},
	}, nil
}

func (svc *service) FindAll(ctx context.Context, payload *model.PaginationRequest) (*model.PaginationResponse, error) {
	p := bluemonday.StrictPolicy()
	resp := new(model.PaginationResponse)
	payload.BuildRequest()
	userID, err := primitive.ObjectIDFromHex(ctx.Value("userID").(string))
	if err != nil {
		return nil, err
	}
	role := ctx.Value("role").(string)

	products := make([]*model.ProductResponse, 0)

	var filter *model.ProductFilter

	if role == constant.ROLE_USER {
		filter = &model.ProductFilter{
			UserID: userID,
		}
	}

	results, err := svc.repository.ProductRepository.FindAll(ctx, svc.db, filter, payload)
	for _, v := range results {
		products = append(products, &model.ProductResponse{
			ID:          v.ID.Hex(),
			UserID:      v.UserID.Hex(),
			Name:        p.Sanitize(v.Name),
			Description: p.Sanitize(v.Description),
			Price:       v.Price,
			DateColumn: model.DateColumn{
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
				DeletedAt: v.DeletedAt,
			},
		})
	}
	count, err := svc.repository.ProductRepository.Count(ctx, svc.db, filter, payload)
	resp.BuildResponse(payload, products, count)
	return resp, nil
}

func (svc *service) FindProductByID(ctx context.Context, payload *model.GetProductDetailRequest) (*model.ProductResponse, error) {
	p := bluemonday.StrictPolicy()
	userID, err := primitive.ObjectIDFromHex(ctx.Value("userID").(string))
	if err != nil {
		return nil, err
	}
	role := ctx.Value("role").(string)

	product := new(model.ProductResponse)

	var filter *model.ProductFilter

	if role == constant.ROLE_USER {
		filter = &model.ProductFilter{
			UserID: userID,
		}
	}
	productID, err := primitive.ObjectIDFromHex(payload.ID)
	if err != nil {
		return nil, err
	}
	input := &database.Product{
		ID: productID,
	}
	result, err := svc.repository.ProductRepository.FindProductByID(ctx, svc.db, filter, input)
	if err != nil {
		return nil, err
	}

	product = &model.ProductResponse{
		ID:          result.ID.Hex(),
		UserID:      result.UserID.Hex(),
		Name:        p.Sanitize(result.Name),
		Description: p.Sanitize(result.Description),
		Price:       result.Price,
		DateColumn: model.DateColumn{
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
			DeletedAt: result.DeletedAt,
		},
	}
	return product, nil
}

func (svc *service) Update(ctx context.Context, payload *model.UpdateProductRequest) (*model.ProductResponse, error) {
	p := bluemonday.StrictPolicy()
	userID, err := primitive.ObjectIDFromHex(ctx.Value("userID").(string))
	if err != nil {
		return nil, err
	}
	role := ctx.Value("role").(string)

	product := new(model.ProductResponse)

	var filter *model.ProductFilter

	if role == constant.ROLE_USER {
		filter = &model.ProductFilter{
			UserID: userID,
		}
	}
	productID, err := primitive.ObjectIDFromHex(payload.ID)
	if err != nil {
		return nil, err
	}
	input := &database.Product{
		ID: productID,
	}
	result, err := svc.repository.ProductRepository.FindProductByID(ctx, svc.db, filter, input)
	if err != nil {
		return nil, err
	}

	result = &database.Product{
		ID:          result.ID,
		UserID:      result.UserID,
		Name:        payload.Name,
		Description: payload.Description,
		Price:       payload.Price,
		DateColumn: database.DateColumn{
			CreatedAt: result.CreatedAt,
			UpdatedAt: time.Now().Unix(),
			DeletedAt: result.DeletedAt,
		},
	}

	result, err = svc.repository.ProductRepository.Update(ctx, svc.db, filter, result)
	if err != nil {
		return nil, err
	}

	product = &model.ProductResponse{
		ID:          result.ID.Hex(),
		UserID:      result.UserID.Hex(),
		Name:        p.Sanitize(result.Name),
		Description: p.Sanitize(result.Description),
		Price:       result.Price,
		DateColumn: model.DateColumn{
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
			DeletedAt: result.DeletedAt,
		},
	}

	return product, nil
}

func (svc *service) Delete(ctx context.Context, payload *model.DeleteProductRequest) error {
	userID, err := primitive.ObjectIDFromHex(ctx.Value("userID").(string))
	if err != nil {
		return err
	}
	role := ctx.Value("role").(string)

	var filter *model.ProductFilter

	if role == constant.ROLE_USER {
		filter = &model.ProductFilter{
			UserID: userID,
		}
	}
	productID, err := primitive.ObjectIDFromHex(payload.ID)
	if err != nil {
		return err
	}
	input := &database.Product{
		ID: productID,
	}
	result, err := svc.repository.ProductRepository.FindProductByID(ctx, svc.db, filter, input)
	if err != nil {
		return err
	}

	tn := time.Now().Unix()

	result = &database.Product{
		ID:          result.ID,
		UserID:      result.UserID,
		Name:        result.Name,
		Description: result.Description,
		Price:       result.Price,
		DateColumn: database.DateColumn{
			CreatedAt: result.CreatedAt,
			UpdatedAt: tn,
			DeletedAt: &tn,
		},
	}

	err = svc.repository.ProductRepository.Delete(ctx, svc.db, filter, result)
	if err != nil {
		return err
	}

	return nil
}
