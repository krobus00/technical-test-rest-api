package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductResponse struct {
	ID          string  `json:"id"`
	UserID      string  `json:"userId"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	DateColumn
}

type ProductFilter struct {
	UserID primitive.ObjectID
}

type CreateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type GetProductDetailRequest struct {
	ID string `json:"id" param:"id"`
}

type UpdateProductRequest struct {
	ID          string  `json:"id" param:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type DeleteProductRequest struct {
	ID string `json:"id" param:"id"`
}
