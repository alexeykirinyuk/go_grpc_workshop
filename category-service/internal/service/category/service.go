package category

import (
	"context"
	"github.com/alexeykirinyuk/go_grpc_workshop/category-service/internal/pkg/internal_errors"
	"github.com/pkg/errors"
)

type Service struct {
	r Repository
}

type Repository interface {
	GetCategories(ctx context.Context) (Categories, error)
}

func New(repository Repository) *Service {
	return &Service{
		r: repository,
	}
}

var ErrNoCategory = errors.Wrap(internal_errors.ErrNotFound, "category not found")

func (s Service) GetCategoryByID(ctx context.Context, id uint64) (*Category, error) {
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
