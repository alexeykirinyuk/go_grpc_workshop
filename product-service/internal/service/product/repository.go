package product_service

import (
	"context"
	"github.com/jmoiron/sqlx"
)

var nextID int64 = 1

type Repository struct {
	db *sqlx.DB
}

func newRepo(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) SaveProduct(ctx context.Context, product *Product) error {
	product.ID = nextID

	nextID++

	return nil
}
