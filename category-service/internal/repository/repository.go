package repository

import (
	"context"
	"github.com/alexeykirinyuk/go_grpc_workshop/category-service/internal/model"
)

type CategoryRepository struct {
}

var categories = model.Categories{
	{
		ID:   1,
		Name: "Toys",
	},
	{
		ID:   2,
		Name: "Laptops",
	},
	{
		ID:   3,
		Name: "Auto",
	},
}

func New() *CategoryRepository {
	return &CategoryRepository{}
}

func (r CategoryRepository) GetCategories(_ context.Context) (model.Categories, error) {
	return categories, nil
}
