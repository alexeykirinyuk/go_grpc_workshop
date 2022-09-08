package service

import (
	"context"
	"github.com/alexeykirinyuk/go_grpc_workshop/category_service/internal/model"
	"github.com/alexeykirinyuk/go_grpc_workshop/category_service/internal/pkg/internal_errors"
	"github.com/pkg/errors"
)

type Service struct {
	r RepositoryInterface
}

type RepositoryInterface interface {
	GetCategories(ctx context.Context) (model.Categories, error)
}

func New(repository RepositoryInterface) *Service {
	return &Service{
		r: repository,
	}
}

var ErrNoCategory = errors.Wrap(internal_errors.ErrNotFound, "category not found")

func (s Service) GetCategoryByID(ctx context.Context, id uint64) (*model.Category, error) {
	cats, err := s.r.GetCategories(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "repository.GetCategories")
	}

	cat := cats.FilterByID(id)
	if cat == nil {
		return nil, ErrNoCategory
	}

	return cat, nil
}
