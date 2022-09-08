package product_service

import (
	"context"
	pb "github.com/alexeykirinyuk/go_grpc_workshop/category-service/pkg/category-service"
	"github.com/alexeykirinyuk/go_grpc_workshop/product_service/internal/pkg/internal_errors"
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
		return nil, internal_errors.WrongCategory
	}

	if err := p.repo.SaveProduct(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}
