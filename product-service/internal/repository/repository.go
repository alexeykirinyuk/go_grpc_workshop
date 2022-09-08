package repository

import (
	"context"
	"github.com/alexeykirinyuk/go_grpc_workshop/product_service/internal/model"
	"github.com/pkg/errors"
)

type ProductRepository struct {
}

func New() *ProductRepository {
	return &ProductRepository{}
}

func (r ProductRepository) CreateProduct(_ context.Context) (model.Product, error) {
	return model.Product{}, errors.New("not implemented")
}
