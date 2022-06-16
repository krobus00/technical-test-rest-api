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
	Name        string  `json:"name" validate:"required,min=5"`
	Description string  `json:"description" validate:"required,min=5"`
	Price       float64 `json:"price" validate:"required,min=1"`
}

type GetProductDetailRequest struct {
	ID string `json:"id" param:"id" validate:"required"`
}

type UpdateProductRequest struct {
	ID          string  `json:"id" param:"id" validate:"required"`
	Name        string  `json:"name" validate:"required,min=5"`
	Description string  `json:"description" validate:"required,min=5"`
	Price       float64 `json:"price" validate:"required,min=1"`
}

type DeleteProductRequest struct {
	ID string `json:"id" param:"id" validate:"required"`
}
