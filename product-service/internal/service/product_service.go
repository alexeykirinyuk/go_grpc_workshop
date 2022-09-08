package service

import (
	"context"
	"github.com/alexeykirinyuk/go_grpc_workshop/product_service/internal/model"
	"github.com/pkg/errors"
)

type Service struct {
	r RepositoryInterface
}

type RepositoryInterface interface {
	CreateProduct(_ context.Context) (model.Product, error)
}

func New(repository RepositoryInterface) *Service {
	return &Service{
		r: repository,
	}
}

func (s Service) CreateProduct(ctx context.Context) (*model.Product, error) {
	return nil, errors.New("not implemented")
}
