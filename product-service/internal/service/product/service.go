package product_service

import (
	"context"
)

//go:generate mockgen -destination=repository_mock_test.go -self_package=github.com/alexeykirinyuk/go_grpc_workshop/product_service/internal/service/product -package=product_service . IRepository

type IRepository interface {
	SaveProduct(ctx context.Context, product *Product) error
}

type Service struct {
	repo IRepository
}

func NewService() *Service {
	return &Service{
		repo: newRepo(),
	}
}

func (p *Service) CreateProduct(
	ctx context.Context,
	name string,
	categoryID int64,
) (*Product, error) {
	product := &Product{
		ID:         0,
		Name:       name,
		CategoryId: categoryID,
	}

	if err := p.repo.SaveProduct(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}
