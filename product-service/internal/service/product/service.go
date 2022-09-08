package product_service

import (
	"context"
	"errors"
	pb "github.com/alexeykirinyuk/go_grpc_workshop/category-service/pkg/category-service"
)

//go:generate mockgen -destination=service_mocks_test.go -self_package=github.com/alexeykirinyuk/go_grpc_workshop/product_service/internal/service/product -package=product_service . IRepository,ICategoryClient

type IRepository interface {
	SaveProduct(ctx context.Context, product *Product) error
}

type ICategoryClient interface {
	IsCategoryExists(ctx context.Context, id int64) (exists bool, err error)
}

type Service struct {
	repo   IRepository
	client ICategoryClient
}

func NewService(grpcClient pb.CategoryServiceClient) *Service {
	return &Service{
		repo:   newRepo(),
		client: newCategoryClient(grpcClient),
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

	exists, err := p.client.IsCategoryExists(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, errors.New("category does not exists")
	}

	if err := p.repo.SaveProduct(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}
