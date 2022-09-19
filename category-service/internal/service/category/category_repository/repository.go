package category_repository

import (
	"context"
	"github.com/alexeykirinyuk/go_grpc_workshop/category-service/internal/service/category"
)

type CategoryRepository struct {
}

func New() *CategoryRepository {
	return &CategoryRepository{}
}

var categories category.Categories

func (r CategoryRepository) GetCategories(_ context.Context) (category.Categories, error) {
	return categories, nil
}
