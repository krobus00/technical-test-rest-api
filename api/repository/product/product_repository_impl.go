package product

import (
	"context"

	"github.com/krobus00/technical-test-rest-api/model"
	"github.com/krobus00/technical-test-rest-api/model/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repository) Store(ctx context.Context, db *mongo.Database, input *database.Product) (*database.Product, error) {
	var id primitive.ObjectID

	result, err := db.Collection(r.GetCollectionName()).InsertOne(ctx, input)
	if err != nil {
		r.logger.Zap.Error(err.Error())
		return nil, err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		id = oid
	}

	return &database.Product{
		ID:          id,
		UserID:      input.UserID,
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		DateColumn:  input.DateColumn,
	}, nil
}

func (r *repository) Count(ctx context.Context, db *mongo.Database, filter *model.ProductFilter, paginationFilter *model.PaginationRequest) (int64, error) {

	filterQuery := bson.M{"deleted_at": nil}
	if filter != nil {
		filterQuery["user_id"] = filter.UserID
	}

	if len(paginationFilter.Search) > 0 {
		filterQuery["name"] = bson.M{"$regex": paginationFilter.Search, "$options": "im"}
		filterQuery["description"] = bson.M{"$regex": paginationFilter.Search, "$options": "im"}
	}

	count, err := db.Collection(r.GetCollectionName()).CountDocuments(ctx, filterQuery)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 0, nil
		}
		return 0, err
	}

	return count, nil
}

func (r *repository) FindAll(ctx context.Context, db *mongo.Database, filter *model.ProductFilter, paginationFilter *model.PaginationRequest) ([]*database.Product, error) {
	results := make([]*database.Product, 0)

	filterQuery := bson.M{"deleted_at": nil}
	if filter != nil {
		filterQuery["user_id"] = filter.UserID
	}

	options := new(options.FindOptions)

	if len(paginationFilter.Search) > 0 {
		filterQuery["name"] = bson.M{"$regex": paginationFilter.Search, "$options": "im"}
		filterQuery["description"] = bson.M{"$regex": paginationFilter.Search, "$options": "im"}
	}

	if paginationFilter.Limit != 0 {
		options.SetSkip((paginationFilter.Page - 1) * paginationFilter.Limit)
		options.SetLimit(paginationFilter.Limit)
	}

	cur, err := db.Collection(r.GetCollectionName()).Find(ctx, filterQuery, options)
	defer cur.Close(ctx)

	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var el *database.Product
		if err := cur.Decode(&el); err != nil {
			r.logger.Zap.Error(err.Error())
		}

		results = append(results, el)
	}
	return results, nil
}

func (r *repository) FindProductByID(ctx context.Context, db *mongo.Database, filter *model.ProductFilter, input *database.Product) (*database.Product, error) {
	filterQuery := bson.M{"deleted_at": nil, "_id": input.ID}

	if filter != nil {
		filterQuery["user_id"] = filter.UserID
	}

	result := new(database.Product)
	err := db.Collection(r.GetCollectionName()).FindOne(ctx, filterQuery).Decode(result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}

func (r *repository) Update(ctx context.Context, db *mongo.Database, filter *model.ProductFilter, input *database.Product) (*database.Product, error) {
	filterQuery := bson.M{"deleted_at": nil, "_id": input.ID}

	if filter != nil {
		filterQuery["user_id"] = filter.UserID
	}

	fields := bson.M{"$set": input}

	_, err := db.Collection(r.GetCollectionName()).UpdateOne(ctx, filterQuery, fields)

	if err != nil {
		return nil, err
	}

	return input, nil
}

func (r *repository) Delete(ctx context.Context, db *mongo.Database, filter *model.ProductFilter, input *database.Product) error {
	filterQuery := bson.M{"deleted_at": nil, "_id": input.ID}

	if filter != nil {
		filterQuery["user_id"] = filter.UserID
	}

	fields := bson.M{"$set": input}

	_, err := db.Collection(r.GetCollectionName()).UpdateOne(ctx, filterQuery, fields)

	if err != nil {
		return err
	}

	return nil
}
